package handlers

import (
	"encoding/json"
	"errors"
	"github.com/anishchenkoivan/hotel-app/hotel-service/internal/apperrors"
	"github.com/anishchenkoivan/hotel-app/hotel-service/internal/model"
	service "github.com/anishchenkoivan/hotel-app/hotel-service/internal/service/room"
	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"net/http"
)

type RoomHandler struct {
	service *service.RoomService
}

func NewRoomHandler(service *service.RoomService) RoomHandler {
	return RoomHandler{service: service}
}

// CreateRoom
// @Summary Create a new room
// @Accept json
// @Produce json
// @Param hotel body model.RoomData true "Room data"
// @Success 201 {object} uuid.UUID
// @Failure 400 {object} string
// @Failure 500 {object} string
// @Router /room [post]
func (handler *RoomHandler) CreateRoom(w http.ResponseWriter, r *http.Request) {
	var roomData model.RoomData
	err := json.NewDecoder(r.Body).Decode(&roomData)
	if err != nil {
		handler.handleError(apperrors.NewParsingError("CreateRoom: "+err.Error()), w)
	}

	id, err := handler.service.CreateRoom(roomData)
	if err != nil {
		handler.handleError(err, w)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	err = json.NewEncoder(w).Encode(id)
	if err != nil {
		handler.handleError(err, w)
	}
}

// UpdateRoom
// @Summary Update a room
// @Accept json
// @Produce json
// @Param id path uuid.UUID true "Room ID"
// @Param hotel body model.RoomData true "Room data"
// @Success 200 "No Content"
// @Failure 400 {object} string
// @Failure 500 {object} string
// @Router /room/{id} [put]
func (handler *RoomHandler) UpdateRoom(w http.ResponseWriter, r *http.Request) {
	roomId, err := uuid.Parse(mux.Vars(r)["id"])
	if err != nil {
		handler.handleError(apperrors.NewParsingError("UpdateRoom: "+err.Error()), w)
	}
	var roomData model.RoomData
	err = json.NewDecoder(r.Body).Decode(&roomData)
	if err != nil {
		handler.handleError(apperrors.NewParsingError("UpdateRoom: "+err.Error()), w)
	}

	err = handler.service.UpdateRoom(roomId, roomData)
	if err != nil {
		handler.handleError(err, w)
	}

	w.WriteHeader(http.StatusOK)
}

// FindRoomById
// @Summary Get a room by ID
// @Accept json
// @Produce json
// @Param id path uuid.UUID true "Room ID"
// @Success 200 {object} model.Room
// @Failure 400 {object} string
// @Failure 500 {object} string
// @Router /room/{id} [get]
func (handler *RoomHandler) FindRoomById(w http.ResponseWriter, r *http.Request) {
	roomId, err := uuid.Parse(mux.Vars(r)["id"])
	if err != nil {
		handler.handleError(apperrors.NewParsingError("FindRoomById: "+err.Error()), w)
	}

	room, err := handler.service.GetRoomById(roomId)
	if err != nil {
		handler.handleError(err, w)
	}

	err = json.NewEncoder(w).Encode(room)
	if err != nil {
		handler.handleError(err, w)
	}
}

// FindAllRooms
// @Summary	Get a list of all rooms
// @Accept json
// @Produce json
// @Success 200 {object} []model.Room
// @Failure 400 {object} string
// @Failure 500 {object} string
// @Router /room [get]
func (handler *RoomHandler) FindAllRooms(w http.ResponseWriter, r *http.Request) {
	rooms, err := handler.service.GetAllRooms()
	if err != nil {
		handler.handleError(err, w)
	}

	err = json.NewEncoder(w).Encode(rooms)
	if err != nil {
		handler.handleError(err, w)
	}
}

// DeleteRoom
// @Summary Delete a room
// @Accept json
// @Produce json
// @Param id path uuid.UUID true "Room ID"
// @Success 204 "No Content"
// @Failure 400 {object} string
// @Failure 500 {object} string
// @Router /room/{id} [delete]
func (handler *RoomHandler) DeleteRoom(w http.ResponseWriter, r *http.Request) {
	roomId, err := uuid.Parse(mux.Vars(r)["id"])
	if err != nil {
		handler.handleError(apperrors.NewParsingError("DeleteRoom: "+err.Error()), w)
	}

	err = handler.service.DeleteRoom(roomId)
	if err != nil {
		handler.handleError(err, w)
	}
	w.WriteHeader(http.StatusNoContent)
}

func (handler *RoomHandler) handleError(err error, w http.ResponseWriter) {
	if errors.Is(err, apperrors.NotFoundErrorInstance) {
		http.Error(w, "Not found", http.StatusNotFound)
	} else if errors.Is(err, apperrors.ParsingErrorInstance) {
		http.Error(w, "Failed to parse", http.StatusBadRequest)
	} else {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
}
