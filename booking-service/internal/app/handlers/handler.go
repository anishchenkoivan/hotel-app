package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/anishchenkoivan/hotel-app/booking-service/internal/service"
)

type Handler struct {
	service service.Service
}

func NewlHandler(repo service.Repository) Handler {
  service := service.NewService(repo)
	return Handler{service: service}
}

func (handler *Handler) GetById(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	encoder := json.NewEncoder(w)

	r.ParseForm()
	query := GetByIdQuery{}
	err := decoder.Decode(&query)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

  reserv, err := handler.service.GetById(query.Id)

  if err != nil {
    w.WriteHeader(http.StatusInternalServerError)
    return
  }

  resp := GetByIdResponse{reserv}
  err = encoder.Encode(resp)

  if err != nil {
    w.WriteHeader(http.StatusInternalServerError)
    return
  }

	w.WriteHeader(http.StatusOK)
}

func (handler *Handler) SearchByPhone(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	encoder := json.NewEncoder(w)

	r.ParseForm()
	query := SearchByPhoneQuery{}
	err := decoder.Decode(&query)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

  reservations, err := handler.service.SearchByPhone(query.phone)

  if err != nil {
    w.WriteHeader(http.StatusInternalServerError)
    return
  }

  resp := SearchByPhoneResponse{reservations}
  err = encoder.Encode(resp)

  if err != nil {
    w.WriteHeader(http.StatusInternalServerError)
    return
  }

	w.WriteHeader(http.StatusOK)
}

func (handler *Handler) AddReservation(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	encoder := json.NewEncoder(w)

	r.ParseForm()
	query := AddReservationQuery{}
	err := decoder.Decode(&query)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	uuid, err := handler.service.AddReservation(query.ReservationData)

  if err != nil {
    w.WriteHeader(http.StatusInternalServerError)
    return
  }

	resp := AddReservationResponse{uuid}
	err = encoder.Encode(resp)

  if err != nil {
    w.WriteHeader(http.StatusInternalServerError)
    return
  }

	w.WriteHeader(http.StatusOK)
}
