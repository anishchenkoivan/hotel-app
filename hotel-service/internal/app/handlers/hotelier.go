package handlers

import (
	service "github.com/anishchenkoivan/hotel-app/hotel-service/internal/service/hotelier"
	"net/http"
)

type HotelierHandler struct {
	service service.HotelierService
}

func NewHotelierHandler() HotelierHandler {
	return HotelierHandler{service: service.NewHotelierService()}
}

func (handler *HotelierHandler) CreateHotelier(w http.ResponseWriter, r *http.Request) {

}

func (handler *HotelierHandler) UpdateHotelier(w http.ResponseWriter, r *http.Request) {

}

func (handler *HotelierHandler) DeleteHotelier(w http.ResponseWriter, r *http.Request) {

}
