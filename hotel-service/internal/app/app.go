package app

import (
	"github.com/anishchenkoivan/hotel-app/hotel-service/internal/app/handlers"
	"github.com/gorilla/mux"
)

type HotelServiceApp struct {
	hotelHandler    handlers.HotelHandler
	hotelierHandler handlers.HotelierHandler
	roomHandler     handlers.RoomHandler
}

func NewHotelServiceApp() *HotelServiceApp {
	router := mux.NewRouter().PathPrefix("/hotel-service/api").Subrouter()

	hotelApp := HotelServiceApp{
		handlers.NewHotelHandler(),
		handlers.NewHotelierHandler(),
		handlers.NewRoomHandler(),
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
	return nil
}

func (app *HotelServiceApp) Shutdown() error {
	return nil
}
