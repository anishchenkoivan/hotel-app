package app

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/anishchenkoivan/hotel-app/booking-service/config"
	"github.com/anishchenkoivan/hotel-app/booking-service/internal/app/handlers"
	"github.com/anishchenkoivan/hotel-app/booking-service/internal/clients"
	"github.com/anishchenkoivan/hotel-app/booking-service/internal/repository"
	"github.com/anishchenkoivan/hotel-app/booking-service/internal/service"
	"github.com/gorilla/mux"
	"golang.org/x/sync/errgroup"
)

type BookingServiceApp struct {
	httpServer *http.Server
	config     config.Config
}

func NewBookingServiceApp(conf config.Config) (*BookingServiceApp, error) {
	repo, err := repository.NewPostgresRepository(conf.Db)

	if err != nil {
		return &BookingServiceApp{}, fmt.Errorf("Can't create repository: %v", err)
	}

	hs_client, err := clients.NewHotelService(conf.HotelService)

	if err != nil {
		return &BookingServiceApp{}, fmt.Errorf("Can't connect to hotel-service: %v", err)
	}

	service := service.NewService(repo, hs_client)
	handler := handlers.NewlHandler(service)
	router := mux.NewRouter().PathPrefix("/booking-service/api").Subrouter()

	router.HandleFunc("/add-reservation", handler.AddReservation).Methods("POST")
	router.HandleFunc("/get-by-id/{id}", handler.GetById).Methods("GET")
	router.HandleFunc("/search-by-phone/{phone}", handler.SearchByPhone).Methods("GET")
	router.HandleFunc("/get-room-reservations/{room_id}", handler.SearchByPhone).Methods("GET")

	httpServer := http.Server{
		Addr:    conf.Server.Host + ":" + conf.Server.Port,
		Handler: router,
	}

	return &BookingServiceApp{&httpServer, conf}, nil
}

func (app *BookingServiceApp) Start(ctx context.Context) error {
	group, ctx := errgroup.WithContext(ctx)

	group.Go(func() error {
		if err := app.httpServer.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			return err
		}
		return nil
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
	err := app.httpServer.Shutdown(shutdownCtx)
	return err
}
