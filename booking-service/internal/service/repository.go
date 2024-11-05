package service

import (
	"github.com/anishchenkoivan/hotel-app/booking-service/internal/model"
	"github.com/google/uuid"
)

type Repository interface {
	GetById(id uuid.UUID) (*model.Reservation, error)
  GetByPhone(phone string) (*model.Reservation, error)
  GetByName(name string, surname string) (*model.Reservation, error)
  Put(reserv model.Reservation) error
}
