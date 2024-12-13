package app

import (
	"context"
	"errors"
	"net/http"
	"time"

//   "github.com/anishchenkoivan/hotel-app/api/code/bookingservice_api"
	"github.com/anishchenkoivan/hotel-app/booking-service/config"
	"github.com/anishchenkoivan/hotel-app/booking-service/internal/app/handlers"
	"github.com/anishchenkoivan/hotel-app/booking-service/internal/service"
	"github.com/gorilla/mux"
	"golang.org/x/sync/errgroup"
)

type BookingServiceApp struct {
	httpServer *http.Server
	config     config.Config
}

func NewBookingServiceApp(repo service.Repository, cfg config.Config) *BookingServiceApp {
	handler := handlers.NewlHandler(repo)

	router := mux.NewRouter().PathPrefix("/booking-service/api").Subrouter()

	router.HandleFunc("/add-reservation", handler.AddReservation).Methods("POST")
	router.HandleFunc("/get-by-id/{id}", handler.GetById).Methods("GET")
	router.HandleFunc("/search-by-phone/{phone}", handler.SearchByPhone).Methods("GET")
	router.HandleFunc("/get-room-reservations/{room_id}", handler.SearchByPhone).Methods("GET")

	httpServer := http.Server{
		Addr:    cfg.Server.Host + ":" + cfg.Server.Port,
		Handler: router,
	}

	return &BookingServiceApp{&httpServer, cfg}
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
