package handlers

import (
	"encoding/json"
	"errors"
	"github.com/anishchenkoivan/hotel-app/hotel-service/internal/app/handlers/dto"
	"github.com/anishchenkoivan/hotel-app/hotel-service/internal/apperrors"
	service "github.com/anishchenkoivan/hotel-app/hotel-service/internal/service/hotelier"
	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"net/http"
)

type HotelierHandler struct {
	service *service.HotelierService
}

func NewHotelierHandler(service *service.HotelierService) HotelierHandler {
	return HotelierHandler{service: service}
}

// FindHotelierById
// @Summary Get a hotelier by ID
// @Accept json
// @Produce json
// @Param id path string true "Hotelier ID"
// @Success 200 {object} dto.HotelierDto
// @Failure 400 {object} string
// @Failure 500 {object} string
// @Router /hotelier/{id} [get]
func (handler *HotelierHandler) FindHotelierById(w http.ResponseWriter, r *http.Request) {
	hotelierId, err := uuid.Parse(mux.Vars(r)["id"])
	if err != nil {
		handler.handleError(apperrors.NewParsingError("FindHotelierById: "+err.Error()), w)
	}

	hotel, err := handler.service.GetHotelierById(hotelierId)
	if err != nil {
		handler.handleError(err, w)
	}

	err = json.NewEncoder(w).Encode(hotel)
	if err != nil {
		handler.handleError(apperrors.NewParsingError("FindHotelierById: "+err.Error()), w)
	}
}

// FindHotelierByTelegramId
// @Summary Get a hotelier by Telegram ID
// @Accept json
// @Produce json
// @Param telegram-id path string true "Hotelier ID"
// @Success 200 {object} dto.HotelierDto
// @Failure 400 {object} string
// @Failure 500 {object} string
// @Router /hotelier/telegram/{telegram-id} [get]
func (handler *HotelierHandler) FindHotelierByTelegramId(w http.ResponseWriter, r *http.Request) {
	hotelierTelegramId := mux.Vars(r)["telegram-id"]

	hotelier, err := handler.service.GetHotelierByTelegramId(hotelierTelegramId)
	if err != nil {
		handler.handleError(err, w)
	}

	err = json.NewEncoder(w).Encode(hotelier)
	if err != nil {
		handler.handleError(apperrors.NewParsingError("FindHotelierTelegramId: "+err.Error()), w)
	}
}

// CreateHotelier
// @Summary Create a new hotelier
// @Accept json
// @Produce json
// @Param hotel body dto.HotelierModifyDto true "Hotelier data"
// @Success 201 {object} uuid.UUID
// @Failure 400 {object} string
// @Failure 500 {object} string
// @Router /hotelier [post]
func (handler *HotelierHandler) CreateHotelier(w http.ResponseWriter, r *http.Request) {
	var hotelierDto dto.HotelierModifyDto
	err := json.NewDecoder(r.Body).Decode(&hotelierDto)
	if err != nil {
		handler.handleError(apperrors.NewParsingError("CreateHotelier: "+err.Error()), w)
	}

	hotelierId, err := handler.service.CreateHotelier(hotelierDto)
	if err != nil {
		handler.handleError(err, w)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	err = json.NewEncoder(w).Encode(hotelierId)
	if err != nil {
		handler.handleError(apperrors.NewParsingError("CreateHotelier: "+err.Error()), w)
	}
}

// UpdateHotelier
// @Summary Update a hotelier
// @Accept json
// @Produce json
// @Param id path string true "Hotelier ID"
// @Param hotel body dto.HotelModifyDto true "Hotelier data"
// @Success 200 "No Content"
// @Failure 400 {object} string
// @Failure 500 {object} string
// @Router /hotelier/{id} [put]
func (handler *HotelierHandler) UpdateHotelier(w http.ResponseWriter, r *http.Request) {
	hotelierId, err := uuid.Parse(mux.Vars(r)["id"])
	if err != nil {
		handler.handleError(apperrors.NewParsingError("UpdateHotelier: "+err.Error()), w)
	}

	var hotelierModifyDto dto.HotelierModifyDto
	err = json.NewDecoder(r.Body).Decode(&hotelierModifyDto)
	if err != nil {
		handler.handleError(apperrors.NewParsingError("UpdateHotelier: "+err.Error()), w)
	}

	err = handler.service.UpdateHotelier(hotelierId, hotelierModifyDto)
	if err != nil {
		handler.handleError(err, w)
	}

	w.WriteHeader(http.StatusOK)
}

// DeleteHotelier
// @Summary Delete a hotelier
// @Accept json
// @Produce json
// @Param id path string true "Hotelier ID"
// @Success 204 "No Content"
// @Failure 400 {object} string
// @Failure 500 {object} string
// @Router /hotelier/{id} [delete]
func (handler *HotelierHandler) DeleteHotelier(w http.ResponseWriter, r *http.Request) {
	hotelierId, err := uuid.Parse(mux.Vars(r)["id"])
	if err != nil {
		handler.handleError(apperrors.NewParsingError("DeleteHotelier: "+err.Error()), w)
	}

	err = handler.service.DeleteHotelier(hotelierId)
	if err != nil {
		handler.handleError(err, w)
	}
	w.WriteHeader(http.StatusNoContent)
}

func (handler *HotelierHandler) handleError(err error, w http.ResponseWriter) {
	if errors.Is(err, apperrors.NotFoundErrorInstance) {
		http.Error(w, "Not found", http.StatusNotFound)
	} else if errors.Is(err, apperrors.ParsingErrorInstance) {
		http.Error(w, "Failed to parse", http.StatusBadRequest)
	} else if errors.Is(err, apperrors.AccessDeniedErrorInstance) {
		http.Error(w, "Access denied", http.StatusForbidden)
	} else {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
}
