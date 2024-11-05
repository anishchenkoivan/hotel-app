package handlers

import (
	service "github.com/anishchenkoivan/hotel-app/hotel-service/internal/service/hotel"
	"net/http"
)

type HotelHandler struct {
	service service.HotelService
}

func NewHotelHandler() HotelHandler {
	return HotelHandler{service: service.NewHotelService()}
}

func (handler *HotelHandler) CreateHotel(w http.ResponseWriter, r *http.Request) {

}

func (handler *HotelHandler) UpdateHotel(w http.ResponseWriter, r *http.Request) {

}

func (handler *HotelHandler) FindHotelById(w http.ResponseWriter, r *http.Request) {

}

func (handler *HotelHandler) FindAllHotels(w http.ResponseWriter, r *http.Request) {

}

func (handler *HotelHandler) DeleteHotelById(w http.ResponseWriter, r *http.Request) {

}
