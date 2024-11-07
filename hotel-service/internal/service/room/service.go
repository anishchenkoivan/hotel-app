package service

import (
	"fmt"
	"github.com/anishchenkoivan/hotel-app/hotel-service/internal/model"
	"github.com/anishchenkoivan/hotel-app/hotel-service/internal/repository"
	"github.com/google/uuid"
)

type RoomService struct {
	repository RoomRepository
}

func NewRoomService() RoomService {
	return RoomService{
		repository.NewPostgresRoomRepository(),
	}
}

func (service *RoomService) GetRoomById(id uuid.UUID) (*model.Room, error) {
	return service.repository.Get(id)
}

func (service *RoomService) GetAllRooms() ([]*model.Room, error) {
	return service.repository.GetAll()
}

func (service *RoomService) CreateRoom(roomData model.RoomData) (*model.Room, error) {
	id, err := uuid.NewRandom()
	if err != nil {
		return nil, err
	}

	room := model.Room{
		Id:       id,
		RoomData: roomData,
	}

	err = service.repository.Put(&room)
	if err != nil {
		return nil, err
	}
	return &room, nil
}

func (service *RoomService) UpdateRoom(id uuid.UUID, roomData model.RoomData) error {
	room, err := service.repository.Get(id)
	if err != nil {
		return fmt.Errorf("UpdateRoom: %w", err)
	}
	room.RoomData = roomData
	err = service.repository.Put(room)
	if err != nil {
		return fmt.Errorf("UpdateRoom: %w", err)
	}
	return nil
}

func (service *RoomService) DeleteRoom(id uuid.UUID) error {
	return service.repository.Remove(id)
}
