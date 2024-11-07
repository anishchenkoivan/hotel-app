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
	server http.Server
}

func NewHotelServiceApp() *HotelServiceApp {
	router := mux.NewRouter().PathPrefix("/hotel-service/api").Subrouter()

	hotelApp := HotelServiceApp{
		http.Server{
			Addr:    ":8080",
			Handler: router,
		},
	}

	hotelHandler := handlers.NewHotelHandler()
	hotelierHandler := handlers.NewHotelierHandler()
	roomHandler := handlers.NewRoomHandler()

	router.HandleFunc("/hotel", hotelHandler.CreateHotel).Methods("POST")
	router.HandleFunc("/hotel", hotelHandler.FindAllHotels).Methods("GET")
	router.HandleFunc("/hotel/{id}", hotelHandler.UpdateHotel).Methods("PUT")
	router.HandleFunc("/hotel/{id}", hotelHandler.FindHotelById).Methods("GET")
	router.HandleFunc("/hotel/{id}", hotelHandler.DeleteHotel).Methods("DELETE")

	router.HandleFunc("/hotelier", hotelierHandler.CreateHotelier).Methods("POST")
	router.HandleFunc("/hotelier/{id}", hotelierHandler.UpdateHotelier).Methods("PUT")
	router.HandleFunc("/hotelier/{id}", hotelierHandler.DeleteHotelier).Methods("DELETE")

	router.HandleFunc("/room", roomHandler.FindAllRooms).Methods("POST")
	router.HandleFunc("/room", roomHandler.FindAllRooms).Methods("GET")
	router.HandleFunc("/room/{id}", roomHandler.UpdateRoom).Methods("PUT")
	router.HandleFunc("/room/{id}", roomHandler.FindRoomById).Methods("GET")
	router.HandleFunc("/room/{id}", roomHandler.DeleteRoom).Methods("DELETE")

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
