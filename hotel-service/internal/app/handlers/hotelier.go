package handlers

import (
	"encoding/json"
	"errors"
	"github.com/anishchenkoivan/hotel-app/hotel-service/internal/apperrors"
	"github.com/anishchenkoivan/hotel-app/hotel-service/internal/model"
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

func (handler *HotelierHandler) CreateHotelier(w http.ResponseWriter, r *http.Request) {
	var hotelierData model.HotelierData
	err := json.NewDecoder(r.Body).Decode(&hotelierData)
	if err != nil {
		handler.handleError(apperrors.NewParsingError("CreateHotelier: "+err.Error()), w)
	}

	hotelierId, err := handler.service.CreateHotelier(hotelierData)
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

func (handler *HotelierHandler) UpdateHotelier(w http.ResponseWriter, r *http.Request) {
	hotelierId, err := uuid.Parse(mux.Vars(r)["id"])
	if err != nil {
		handler.handleError(apperrors.NewParsingError("UpdateHotelier: "+err.Error()), w)
	}

	var hotelierData model.HotelierData
	err = json.NewDecoder(r.Body).Decode(&hotelierData)
	if err != nil {
		handler.handleError(apperrors.NewParsingError("UpdateHotelier: "+err.Error()), w)
	}

	err = handler.service.UpdateHotelier(hotelierId, hotelierData)
	if err != nil {
		handler.handleError(err, w)
	}

	w.WriteHeader(http.StatusOK)
}

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
	} else {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
}