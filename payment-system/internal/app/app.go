package app

import (
	"context"
	"errors"
	"fmt"
	"github.com/anishchenkoivan/hotel-app/payment-system/api/api_v1pb"
	"github.com/anishchenkoivan/hotel-app/payment-system/config"
	"github.com/anishchenkoivan/hotel-app/payment-system/internal/app/handlers"
	"github.com/anishchenkoivan/hotel-app/payment-system/internal/model"
	"github.com/gorilla/mux"
	"golang.org/x/sync/errgroup"
	"google.golang.org/grpc"
	"log"
	"net"
	"net/http"
	"os/signal"
	"sync"
	"syscall"
)

type PaymentSystemApp struct {
	server              *http.Server
	grpcServer          *grpc.Server
	config              config.Config
	bookingEntityByHash map[string]model.BookingEntity
	mu                  *sync.Mutex
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
		make(map[string]model.BookingEntity),
		&sync.Mutex{},
	}

	paymentHandler := handlers.NewHTTPPaymentHandler(config, paymentSystemApp.mu, paymentSystemApp.bookingEntityByHash)
	router.HandleFunc("/pay/{token}", paymentHandler.PaymentHandle)

	addPaymentGrpcHandler := handlers.NewAddPaymentGrpcHandler(config, paymentSystemApp.mu, paymentSystemApp.bookingEntityByHash)
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

	if err := app.server.Shutdown(shutdownCtx); err != nil {
		return err
	}
	log.Println("Payment System http server shutted down, waiting for grpc graceful shutdown")
	app.grpcServer.GracefulStop()
	log.Println("Payment System grpc server shutted down")
	log.Println("Payment System stopped")
	return nil
}
