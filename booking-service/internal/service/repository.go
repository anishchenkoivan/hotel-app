package service

import (
	"github.com/anishchenkoivan/hotel-app/booking-service/internal/model"
  "github.com/google/uuid"
)

type Repository interface {
	GetById(id uuid.UUID) (model.Reservation, error)
	SearchByPhone(phone string) ([]model.Reservation, error)
	Put(reserv model.ReservationData) (uuid.UUID, error)
}
