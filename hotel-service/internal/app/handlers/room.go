package handlers

import (
	service "github.com/anishchenkoivan/hotel-app/hotel-service/internal/service/room"
	"net/http"
)

type RoomHandler struct {
	service service.RoomService
}

func NewRoomHandler() RoomHandler {
	return RoomHandler{service: service.NewRoomService()}
}

func (handler *RoomHandler) CreateRoom(w http.ResponseWriter, r *http.Request) {

}

func (handler *RoomHandler) UpdateRoom(w http.ResponseWriter, r *http.Request) {

}

func (handler *RoomHandler) FindRoomById(w http.ResponseWriter, r *http.Request) {

}

func (handler *RoomHandler) FindAllRooms(w http.ResponseWriter, r *http.Request) {

}

func (handler *RoomHandler) DeleteRoomById(w http.ResponseWriter, r *http.Request) {

}
