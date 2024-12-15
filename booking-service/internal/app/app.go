package app

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net"
	"net/http"
	"time"

	"github.com/anishchenkoivan/hotel-app/api/code/bookingservice_api"
	"github.com/anishchenkoivan/hotel-app/booking-service/config"
	"github.com/anishchenkoivan/hotel-app/booking-service/internal/app/handlers"
	"github.com/anishchenkoivan/hotel-app/booking-service/internal/clients"
	"github.com/anishchenkoivan/hotel-app/booking-service/internal/repository"
	"github.com/anishchenkoivan/hotel-app/booking-service/internal/service"
	"github.com/gorilla/mux"
	"golang.org/x/sync/errgroup"
	"google.golang.org/grpc"
)

type BookingServiceApp struct {
	httpServer *http.Server
	grpcServer *grpc.Server
	config     config.Config
}

func NewBookingServiceApp(conf config.Config) (*BookingServiceApp, error) {
	repo, err := repository.NewPostgresRepository(conf.Db)

	if err != nil {
		return nil, fmt.Errorf("Can't create repository: %v", err)
	}

	err = repo.Migrate(conf.Db.Migrations)

	if err != nil {
		return nil, fmt.Errorf("Can't migrate db: %v", err)
	}

	hsClient, err := clients.NewHotelService(conf.HotelService)

	if err != nil {
		return nil, fmt.Errorf("Can't connect to hotel-service: %v", err)
	}

	notificationServiceClient, err := clients.NewNotificationService(conf.Kafka)
	if err != nil {
		return nil, fmt.Errorf("Can't connect to notification-service: %v", err)
	}

	ps_client, err := clients.NewPayementSystem(conf.PaymentSystem)

	if err != nil {
		return nil, fmt.Errorf("Can't connect to payemnt system: %v", err)
	}

	service := service.NewService(repo, hsClient, *notificationServiceClient, ps_client)
	handler := handlers.NewlHandler(&service)
	router := mux.NewRouter().PathPrefix("/booking-service/api").Subrouter()

	router.HandleFunc("/add-reservation", handler.AddReservation).Methods("POST")
	router.HandleFunc("/get-by-id/{id}", handler.GetById).Methods("GET")
	router.HandleFunc("/search-by-phone/{phone}", handler.SearchByPhone).Methods("GET")
	router.HandleFunc("/get-room-reservations/{room_id}", handler.SearchByPhone).Methods("GET")

	httpServer := http.Server{
		Addr:    conf.Server.Host + ":" + conf.Server.Port,
		Handler: router,
	}

	grpcServer := grpc.NewServer()
	grpcHandler := handlers.NewHotelServiceGrpcHandler(&service)
	bookingservice_api.RegisterBookingServiceServer(grpcServer, grpcHandler)

	return &BookingServiceApp{&httpServer, grpcServer, conf}, nil
}

func (app *BookingServiceApp) Start(ctx context.Context) error {
	group, ctx := errgroup.WithContext(ctx)

	group.Go(func() error {
		log.Printf("BookingService: Starting http server on %s:%s", app.config.Server.Host, app.config.Server.Port)
		if err := app.httpServer.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			return err
		}
		return nil
	})

	group.Go(func() error {
		log.Printf("BookingService: Starting grpc server on %s:%s", app.config.BookingService.Host, app.config.BookingService.Port)
		listener, err := net.Listen("tcp", fmt.Sprintf("%s:%s", app.config.BookingService.Host, app.config.BookingService.Port))
		if err != nil {
			return err
		}
		return app.grpcServer.Serve(listener)
	})

	group.Go(func() error {
		<-ctx.Done()
		err := app.Shutdown()
		return err
	})

	err := group.Wait()
	return err
}

func (app *BookingServiceApp) Shutdown() error {
	shutdownCtx, cancel := context.WithTimeout(context.Background(), app.config.App.ShutdownTimeout*time.Second)
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

func (app *BookingServiceApp) stopHttpServer(ctx context.Context) error {
	if err := app.httpServer.Shutdown(ctx); err != nil {
		return err
	}

	log.Println("BookingService: HTTP server has stopped")
	return nil
}

func (app *BookingServiceApp) stopGrpcServer(ctx context.Context) {
	done := make(chan struct{})

	go func() {
		app.grpcServer.GracefulStop()
		close(done)
	}()

	select {
	case <-done:
		log.Println("BookingService: gRPC server has stopped")
	case <-ctx.Done():
		app.grpcServer.Stop()
		log.Println("BookingService: gRPC server has reached stop timeout and has been stopped forcefully")
	}
}
