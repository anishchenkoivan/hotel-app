package service

import (
	"github.com/anishchenkoivan/hotel-app/hotel-service/internal/model"
	"github.com/google/uuid"
)

type RoomRepository interface {
	Get(id uuid.UUID) (*model.Room, error)
	GetAll() ([]*model.Room, error)
	Put(room *model.Room) error
	Remove(id uuid.UUID) error
}
