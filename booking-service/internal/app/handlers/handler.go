package handlers

import (
	"github.com/anishchenkoivan/hotel-app/booking-service/internal/service"
	"net/http"
)

type Handler struct {
	service service.Service
}

func NewlHandler() Handler {
	return Handler{service: service.NewService()}
}

func (handler *Handler) CreateReservation(w http.ResponseWriter, r *http.Request) {

}

func (handler *Handler) FindReservation(w http.ResponseWriter, r *http.Request) {

}
