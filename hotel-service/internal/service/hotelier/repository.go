package service

import (
	"github.com/anishchenkoivan/hotel-app/hotel-service/internal/model"
	"github.com/google/uuid"
)

type HotelierRepository interface {
	Get(id uuid.UUID) (*model.Hotelier, error)
	GetAll() ([]*model.Hotelier, error)
	Put(*model.Hotelier) error
	Remove(*model.Hotelier) error
}
