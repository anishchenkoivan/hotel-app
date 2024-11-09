package app

import (
	"context"
	"errors"
	"net/http"
	"time"

	"github.com/anishchenkoivan/hotel-app/booking-service/internal/app/handlers"
	"github.com/gorilla/mux"
	"golang.org/x/sync/errgroup"
)

type BookingServiceApp struct {
	httpServer *http.Server
}

func NewBookingServiceApp() *BookingServiceApp {
	handler := handlers.NewlHandler()

	router := mux.NewRouter().PathPrefix("/booking-service/api").Subrouter()

	router.HandleFunc("/hotel", handler.AddReservation).Methods("POST")
	router.HandleFunc("/hotel", handler.SearchByPhone).Methods("GET")

	httpServer := http.Server{
		Addr:    ":8080",
		Handler: router,
	}

	return &BookingServiceApp{&httpServer}
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
	shutdownCtx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()
	err := app.httpServer.Shutdown(shutdownCtx)
	return err
}
