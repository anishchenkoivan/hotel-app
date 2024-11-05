package service

import "github.com/anishchenkoivan/hotel-app/hotel-service/internal/repository"

type RoomService struct {
	repository RoomRepository
}

func NewRoomService() RoomService {
	return RoomService{
		repository.NewPostgresRoomRepository(),
	}
}
