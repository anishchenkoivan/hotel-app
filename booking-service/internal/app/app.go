package app

import (
	"github.com/anishchenkoivan/hotel-app/booking-service/internal/app/handlers"
	"github.com/gorilla/mux"
)

type BookingServiceApp struct {
	handler handlers.Handler
}

func NewBookingServiceApp() *BookingServiceApp {
	app := BookingServiceApp{handler: handlers.NewlHandler()}
	router := mux.NewRouter().PathPrefix("/booking-service/api").Subrouter()

	router.HandleFunc("/hotel", app.handler.CreateReservation).Methods("POST")
	router.HandleFunc("/hotel", app.handler.FindReservation).Methods("GET")

	return &app
}

func (app *BookingServiceApp) Start() error {
	return nil
}

func (app *BookingServiceApp) Shutdown() error {
	return nil
}
