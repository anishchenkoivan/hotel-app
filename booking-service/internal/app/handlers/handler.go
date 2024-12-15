package handlers

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"

	"github.com/google/uuid"
	"github.com/gorilla/mux"

	"github.com/anishchenkoivan/hotel-app/booking-service/internal/service"
)

type Handler struct {
	service *service.Service
}

func NewlHandler(service *service.Service) Handler {
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

// GetRoomReservations
// @Summary Search reservation by room id
// @Produce json
// @Param room_id path string true "Room id"
// @Success 200 {object} ReservationsArrayDto
// @Failure 400 {object} string
// @Failure 500 {object} string
// @Router /get-room-reservations/{room_id} [get]
func (handler *Handler) GetRoomReservations(w http.ResponseWriter, r *http.Request) {
	roomIdString := mux.Vars(r)["room_id"]
	roomIdBytes := []byte(roomIdString)
	roomId, err := uuid.FromBytes(roomIdBytes)

	if err != nil {
		http.Error(w, "Failed to decode room id", http.StatusBadRequest)
		return
	}

	reservs, err := handler.service.GetRoomReservations(roomId)

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
// @Param Reservation body CreateReservationDto true "Reservation parametres"
// @Success 200 {object} NewReservationDto
// @Failure 400 {object} string
// @Failure 500 {object} string
// @Router /add-reservation [post]
func (handler *Handler) AddReservation(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	encoder := json.NewEncoder(w)

	err := r.ParseForm()

	if err != nil {
		http.Error(w, "Failed to parse request", http.StatusBadRequest)
		return
	}

	query := CreateReservationDto{}
	err = decoder.Decode(&query)

	if err != nil {
		http.Error(w, "Failed to parse request", http.StatusBadRequest)
		return
	}

	data, err := ReservationFromDto(query)

	if err != nil {
		fmt.Println(err.Error())
		http.Error(w, "Failed to parse request", http.StatusBadRequest)
		return
	}

	id, paymentURL, err := handler.service.AddReservation(data)

	if err != nil {
		if errors.Is(err, service.RepositoryError) {
			log.Printf("BookingService: Failed to insert: %v", err)
			http.Error(w, "Failed to insert", http.StatusInternalServerError)
		} else if errors.Is(err, service.ReservationAlreadyExists) {
			http.Error(w, "Room is reserved on selected time range", http.StatusBadRequest)
		} else if errors.Is(err, service.InvalidReservation) {
			http.Error(w, "Invalid reservation", http.StatusBadRequest)
		} else if errors.Is(err, service.HotelServiceError) {
			log.Printf("BookingService: Failed to get room price: %v", err)
			http.Error(w, "Failed to get room price", http.StatusInternalServerError)
		} else if errors.Is(err, service.PayemntSystemError) {
			log.Printf("BookingService: can't create payemnt webhook: %v", err)
			http.Error(w, "Can't creaate payemnt webohook", http.StatusInternalServerError)
		} else {
			log.Printf("BookingService: Unknown server error: %v", err)
			http.Error(w, "Unknown server error", http.StatusInternalServerError)
		}
		return
	}

	resp := NewReservationDto{Id: id, PaymentUrl: paymentURL}
	err = encoder.Encode(resp)

	if err != nil {
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
		return
	}
}
