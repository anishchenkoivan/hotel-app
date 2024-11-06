package app

import (
	"context"
	"errors"
	"github.com/anishchenkoivan/hotel-app/hotel-service/internal/app/handlers"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"os/signal"
	"syscall"
	"time"

	"golang.org/x/sync/errgroup"
)

type HotelServiceApp struct {
	hotelHandler    handlers.HotelHandler
	hotelierHandler handlers.HotelierHandler
	roomHandler     handlers.RoomHandler
	server          http.Server
}

func NewHotelServiceApp() *HotelServiceApp {
	router := mux.NewRouter().PathPrefix("/hotel-service/api").Subrouter()

	hotelApp := HotelServiceApp{
		handlers.NewHotelHandler(),
		handlers.NewHotelierHandler(),
		handlers.NewRoomHandler(),
		http.Server{
			Addr:    ":8080",
			Handler: router,
		},
	}

	router.HandleFunc("/hotel", hotelApp.hotelHandler.CreateHotel).Methods("POST")
	router.HandleFunc("/hotel", hotelApp.hotelHandler.FindAllHotels).Methods("GET")
	router.HandleFunc("/hotel/{id}", hotelApp.hotelHandler.UpdateHotel).Methods("PUT")
	router.HandleFunc("/hotel/{id}", hotelApp.hotelHandler.FindHotelById).Methods("GET")
	router.HandleFunc("/hotel/{id}", hotelApp.hotelHandler.DeleteHotelById).Methods("DELETE")

	router.HandleFunc("/hotelier", hotelApp.hotelierHandler.CreateHotelier).Methods("POST")
	router.HandleFunc("/hotelier/{id}", hotelApp.hotelierHandler.UpdateHotelier).Methods("PUT")
	router.HandleFunc("/hotelier/{id}", hotelApp.hotelierHandler.DeleteHotelierById).Methods("DELETE")

	router.HandleFunc("/room", hotelApp.roomHandler.FindAllRooms).Methods("POST")
	router.HandleFunc("/room", hotelApp.roomHandler.FindAllRooms).Methods("GET")
	router.HandleFunc("/room/{id}", hotelApp.roomHandler.UpdateRoom).Methods("PUT")
	router.HandleFunc("/room/{id}", hotelApp.roomHandler.FindRoomById).Methods("GET")
	router.HandleFunc("/room/{id}", hotelApp.roomHandler.DeleteRoomById).Methods("DELETE")

	return &hotelApp
}

func (app *HotelServiceApp) Start() error {
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
		<-ctx.Done()
		return app.Stop()
	})

	log.Println("Hotel Service started on port 8080")

	if err := group.Wait(); err != nil {
		return err
	}

	return nil
}

func (app *HotelServiceApp) Stop() error {
	shutdownCtx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := app.server.Shutdown(shutdownCtx); err != nil {
		return err
	}

	log.Println("Hotel Service stopped")
	return nil
}
