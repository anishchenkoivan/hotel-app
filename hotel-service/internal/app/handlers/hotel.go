package handlers

import (
	"encoding/json"
	"errors"
	"github.com/anishchenkoivan/hotel-app/hotel-service/internal/apperrors"
	"github.com/anishchenkoivan/hotel-app/hotel-service/internal/model"
	service "github.com/anishchenkoivan/hotel-app/hotel-service/internal/service/hotel"
	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"net/http"
)

type HotelHandler struct {
	service service.HotelService
}

func NewHotelHandler() HotelHandler {
	return HotelHandler{service: service.NewHotelService()}
}

func (handler *HotelHandler) CreateHotel(w http.ResponseWriter, r *http.Request) {
	var hotelData model.HotelData
	err := json.NewDecoder(r.Body).Decode(&hotelData)
	if err != nil {
		handler.handleError(apperrors.NewParsingError("CreateHotel: "+err.Error()), w)
	}

	hotelId, err := handler.service.CreateHotel(hotelData)
	if err != nil {
		handler.handleError(err, w)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	err = json.NewEncoder(w).Encode(hotelId)
	if err != nil {
		handler.handleError(err, w)
	}
}

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

func (handler *HotelHandler) UpdateHotel(w http.ResponseWriter, r *http.Request) {
	hotelId, err := uuid.Parse(mux.Vars(r)["id"])
	if err != nil {
		handler.handleError(apperrors.NewParsingError("DeleteHotel: "+err.Error()), w)
	}
	var hotelData model.HotelData
	err = json.NewDecoder(r.Body).Decode(&hotelData)
	if err != nil {
		handler.handleError(apperrors.NewParsingError("DeleteHotel: "+err.Error()), w)
	}

	err = handler.service.UpdateHotel(hotelId, hotelData)
	if err != nil {
		handler.handleError(err, w)
	}

	w.WriteHeader(http.StatusOK)
}

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
