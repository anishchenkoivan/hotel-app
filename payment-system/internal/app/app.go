package app

import (
	"context"
	"errors"
	"fmt"
	"github.com/anishchenkoivan/hotel-app/payment-system/api/api_v1pb"
	"github.com/anishchenkoivan/hotel-app/payment-system/config"
	"github.com/anishchenkoivan/hotel-app/payment-system/internal/app/handlers"
	"github.com/anishchenkoivan/hotel-app/payment-system/internal/clients"
	"github.com/anishchenkoivan/hotel-app/payment-system/internal/service"
	"github.com/gorilla/mux"
	"golang.org/x/sync/errgroup"
	"google.golang.org/grpc"
	"log"
	"net"
	"net/http"
	"os/signal"
	"syscall"
)

type PaymentSystemApp struct {
	server     *http.Server
	grpcServer *grpc.Server
	config     config.Config
}

func NewPaymentSystemApp(config config.Config) *PaymentSystemApp {
	router := mux.NewRouter().PathPrefix("/payment-system/api").Subrouter()
	paymentSystemApp := PaymentSystemApp{
		&http.Server{
			Addr:    config.ServerHost + ":" + config.ServerPort,
			Handler: router,
		},
		grpc.NewServer(),
		config,
	}

	bookingServiceClient := clients.NewBookingService(paymentSystemApp.config)
	paymentSystemService := service.NewPaymentSystemService(paymentSystemApp.config, bookingServiceClient)
	paymentHandler := handlers.NewHTTPPaymentHandler(paymentSystemService)
	router.HandleFunc("/pay/{token}", paymentHandler.PaymentHandle)

	addPaymentGrpcHandler := handlers.NewAddPaymentGrpcHandler(paymentSystemService)
	api_v1pb.RegisterPaymentSystemServer(paymentSystemApp.grpcServer, addPaymentGrpcHandler)
	return &paymentSystemApp
}

func (app *PaymentSystemApp) Start() error {
	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	group, ctx := errgroup.WithContext(ctx)
	group.Go(func() error {
		if err := app.server.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			return err
		}
		return nil
	})

	group.Go(func() error {
		listener, err := net.Listen("tcp", fmt.Sprintf("%s:%s", app.config.ServerGrpcHost, app.config.ServerGrpcPort))
		if err != nil {
			return err
		}
		if err = app.grpcServer.Serve(listener); err != nil {
			return err
		}
		return nil
	})

	group.Go(func() error {
		<-ctx.Done()
		return app.Stop()
	})

	log.Println("Payment System started on port" + app.config.ServerPort)

	if err := group.Wait(); err != nil {
		return err
	}

	return nil
}

func (app *PaymentSystemApp) Stop() error {
	shutdownCtx, cancel := context.WithTimeout(context.Background(), app.config.AppShutdownTimeout)
	defer cancel()

	done := make(chan error)

	go func() {
		done <- app.stopHttpServer(shutdownCtx)
	}()

	go func() {
		app.stopGrpcServer(shutdownCtx)
	}()

	if err := <-done; err != nil {
		return err
	}

	return nil
}

func (app *PaymentSystemApp) stopHttpServer(ctx context.Context) error {
	if err := app.server.Shutdown(ctx); err != nil {
		return err
	}

	log.Println("Payment System HTTP server has stopped")
	return nil
}

func (app *PaymentSystemApp) stopGrpcServer(ctx context.Context) {
	done := make(chan struct{})

	go func() {
		app.grpcServer.GracefulStop()
		close(done)
	}()

	select {
	case <-done:
		log.Println("Payment System gRPC server has stopped gracefully")
	case <-ctx.Done():
		app.grpcServer.Stop()
		log.Println("Payment System gRPC server has reached stop timeout and has been stopped forcefully")
	}
}
