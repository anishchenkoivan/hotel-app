package service

import (
	"github.com/anishchenkoivan/hotel-app/hotel-service/internal/model"
	"github.com/google/uuid"
)

type RoomRepository interface {
	Get(id uuid.UUID) (*model.Room, error)
	GetAll() ([]*model.Room, error)
	Put(room *model.Room) (uuid.UUID, error)
	Update(room *model.Room) error
	Delete(id uuid.UUID) error
}
