package service

import (
	"github.com/anishchenkoivan/hotel-app/hotel-service/internal/model"
	"github.com/google/uuid"
)

type HotelRepository interface {
	Get(id uuid.UUID) (*model.Hotel, error)
	GetAll() ([]*model.Hotel, error)
	Put(hotel *model.Hotel) error
	Remove(id uuid.UUID) error
}
