package handlers

import (
	"encoding/json"
	"errors"
	"github.com/anishchenkoivan/hotel-app/hotel-service/internal/app/handlers/dto"
	"github.com/anishchenkoivan/hotel-app/hotel-service/internal/apperrors"
	service "github.com/anishchenkoivan/hotel-app/hotel-service/internal/service/hotel"
	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"net/http"
)

type HotelHandler struct {
	service *service.HotelService
}

func NewHotelHandler(service *service.HotelService) HotelHandler {
	return HotelHandler{service: service}
}

// CreateHotel
// @Summary Create a new hotel
// @Accept json
// @Produce json
// @Param hotel body dto.HotelModifyDto true "Hotel data"
// @Success 201 {object} uuid.UUID
// @Failure 400 {object} string
// @Failure 500 {object} string
// @Router /hotel [post]
func (handler *HotelHandler) CreateHotel(w http.ResponseWriter, r *http.Request) {
	var hotelDto dto.HotelModifyDto
	err := json.NewDecoder(r.Body).Decode(&hotelDto)
	if err != nil {
		handler.handleError(apperrors.NewParsingError("CreateHotel: "+err.Error()), w)
	}

	hotelId, err := handler.service.CreateHotel(hotelDto)
	if err != nil {
		handler.handleError(err, w)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	err = json.NewEncoder(w).Encode(hotelId)
	if err != nil {
		handler.handleError(apperrors.NewParsingError("CreateHotel: "+err.Error()), w)
	}
}

// DeleteHotel
// @Summary Delete a hotel
// @Accept json
// @Produce json
// @Param id path string true "Hotel ID"
// @Success 204 "No Content"
// @Failure 400 {object} string
// @Failure 500 {object} string
// @Router /hotel/{id} [delete]
func (handler *HotelHandler) DeleteHotel(w http.ResponseWriter, r *http.Request) {
	hotelId, err := uuid.Parse(mux.Vars(r)["id"])
	if err != nil {
		handler.handleError(apperrors.NewParsingError("DeleteHotel: "+err.Error()), w)
	}

	err = handler.service.DeleteHotel(hotelId)
	if err != nil {
		handler.handleError(err, w)
	}
	w.WriteHeader(http.StatusNoContent)
}

// UpdateHotel
// @Summary Update a hotel
// @Accept json
// @Produce json
// @Param id path string true "Hotel ID"
// @Param hotel body dto.HotelModifyDto true "Hotel data"
// @Success 200 "No Content"
// @Failure 400 {object} string
// @Failure 500 {object} string
// @Router /hotel/{id} [put]
func (handler *HotelHandler) UpdateHotel(w http.ResponseWriter, r *http.Request) {
	hotelId, err := uuid.Parse(mux.Vars(r)["id"])
	if err != nil {
		handler.handleError(apperrors.NewParsingError("DeleteHotel: "+err.Error()), w)
	}
	var hotelModifyDto dto.HotelModifyDto
	err = json.NewDecoder(r.Body).Decode(&hotelModifyDto)
	if err != nil {
		handler.handleError(apperrors.NewParsingError("DeleteHotel: "+err.Error()), w)
	}

	err = handler.service.UpdateHotel(hotelId, hotelModifyDto)
	if err != nil {
		handler.handleError(err, w)
	}

	w.WriteHeader(http.StatusOK)
}

// FindHotelById
// @Summary Get a hotel by ID
// @Accept json
// @Produce json
// @Param id path string true "Hotel ID"
// @Success 200 {object} dto.HotelDto
// @Failure 400 {object} string
// @Failure 500 {object} string
// @Router /hotel/{id} [get]
func (handler *HotelHandler) FindHotelById(w http.ResponseWriter, r *http.Request) {
	hotelId, err := uuid.Parse(mux.Vars(r)["id"])
	if err != nil {
		handler.handleError(err, w)
		return
	}

	hotel, err := handler.service.GetHotelById(hotelId)
	if err != nil {
		handler.handleError(err, w)
	}

	err = json.NewEncoder(w).Encode(hotel)
	if err != nil {
		handler.handleError(apperrors.NewParsingError("FindHotelById: "+err.Error()), w)
	}
}

// FindAllHotels
// @Summary	Get a list of all hotels
// @Accept json
// @Produce json
// @Success 200 {object} []dto.HotelDto
// @Failure 400 {object} string
// @Failure 500 {object} string
// @Router /hotel [get]
func (handler *HotelHandler) FindAllHotels(w http.ResponseWriter, r *http.Request) {
	hotels, err := handler.service.GetAllHotels()
	if err != nil {
		handler.handleError(err, w)
	}

	err = json.NewEncoder(w).Encode(hotels)
	if err != nil {
		handler.handleError(apperrors.NewParsingError("FindAllHotels: "+err.Error()), w)
	}
}

func (handler *HotelHandler) handleError(err error, w http.ResponseWriter) {
	if errors.Is(err, apperrors.NotFoundErrorInstance) {
		http.Error(w, "Not found", http.StatusNotFound)
	} else if errors.Is(err, apperrors.ParsingErrorInstance) {
		http.Error(w, "Failed to parse", http.StatusBadRequest)
	} else {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
}
