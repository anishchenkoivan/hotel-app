package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/google/uuid"
	"github.com/gorilla/mux"

	"github.com/anishchenkoivan/hotel-app/booking-service/internal/service"
)

type Handler struct {
	service service.Service
}

func NewlHandler(repo service.Repository) Handler {
	service := service.NewService(repo)
	return Handler{service: service}
}

// GetById
// @Summary Get reservation by ID
// @Produce json
// @Param id path string true "Reservation ID"
// @Success 200 {object} ReservationDto
// @Failure 400 {object} string
// @Failure 500 {object} string
// @Router /get-by-id/{id} [get]
func (handler *Handler) GetById(w http.ResponseWriter, r *http.Request) {
	id, err := uuid.Parse(mux.Vars(r)["id"])

	if err != nil {
		http.Error(w, "Failed to parse request", http.StatusBadRequest)
		return
	}

	reserv, err := handler.service.GetById(id)

	if err != nil {
		http.Error(w, "Failed to find by id", http.StatusInternalServerError)
		return
	}

	resp := ReservationDtoFromModel(reserv)
	encoder := json.NewEncoder(w)
	err = encoder.Encode(resp)

	if err != nil {
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
		return
	}
}

// SearchByPhone
// @Summary Search reservation by phone
// @Produce json
// @Param phone path string true "Client phone"
// @Success 200 {object} ReservationsArrayDto
// @Failure 400 {object} string
// @Failure 500 {object} string
// @Router /search-by-phone/{phone} [get]
func (handler *Handler) SearchByPhone(w http.ResponseWriter, r *http.Request) {
	phone := mux.Vars(r)["phone"]
	reservs, err := handler.service.SearchByPhone(phone)

	if err != nil {
		http.Error(w, "Failed to search", http.StatusInternalServerError)
		return
	}

	resp := ReservationsArrayDtoFromModelsArray(reservs)
	encoder := json.NewEncoder(w)
	err = encoder.Encode(resp)

	if err != nil {
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
		return
	}
}

// AddReservation
// @Summary Add reservation
// @Accept json
// @Produce json
// @Param phone query ReservationDto true "Reservation parametres"
// @Success 200 {object} uuid.UUID
// @Failure 400 {object} string
// @Failure 500 {object} string
// @Router /add-reservation [post]
func (handler *Handler) AddReservation(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	encoder := json.NewEncoder(w)

	r.ParseForm()
	query := ReservationDto{}
	err := decoder.Decode(&query)

	if err != nil {
		http.Error(w, "Failed to parse request", http.StatusBadRequest)
		return
	}

	data, err := ReservationDataFromDto(query)

	if err != nil {
    fmt.Println(err.Error())
		http.Error(w, "Failed to parse request", http.StatusBadRequest)
		return
	}

	id, serr := handler.service.AddReservation(data)

	if serr != nil {
		if serr.ErrType == service.RepositoryError {
			http.Error(w, "Failed to insert", http.StatusInternalServerError)
		} else if serr.ErrType == service.ReservationAlreadyExists {
			http.Error(w, "Room is reserved on selected time range", http.StatusBadRequest)
		} else {
			http.Error(w, "Unknown server error", http.StatusInternalServerError)
		}
		return
	}

	resp := AddReservationResponse{id}
	err = encoder.Encode(resp)

	if err != nil {
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
		return
	}
}
