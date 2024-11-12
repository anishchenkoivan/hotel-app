package service

import (
	"fmt"
	"github.com/anishchenkoivan/hotel-app/hotel-service/internal/model"
	"github.com/google/uuid"
)

type RoomService struct {
	repository RoomRepository
}

func NewRoomService(repository RoomRepository) *RoomService {
	return &RoomService{repository: repository}
}

func (service *RoomService) GetRoomById(id uuid.UUID) (*model.Room, error) {
	return service.repository.Get(id)
}

func (service *RoomService) GetAllRooms() ([]*model.Room, error) {
	return service.repository.GetAll()
}

func (service *RoomService) CreateRoom(roomData model.RoomData) (uuid.UUID, error) {
	room := model.Room{
		ID:       uuid.Nil,
		RoomData: roomData,
	}

	id, err := service.repository.Put(&room)
	if err != nil {
		return uuid.Nil, err
	}
	return id, nil
}

func (service *RoomService) UpdateRoom(id uuid.UUID, roomData model.RoomData) error {
	room, err := service.repository.Get(id)
	if err != nil {
		return fmt.Errorf("UpdateRoom: %w", err)
	}
	room.RoomData = roomData
	err = service.repository.Update(room)
	if err != nil {
		return fmt.Errorf("UpdateRoom: %w", err)
	}
	return nil
}

func (service *RoomService) DeleteRoom(id uuid.UUID) error {
	return service.repository.Delete(id)
}
