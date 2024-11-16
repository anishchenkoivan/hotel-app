package service

import (
	"time"

	"github.com/anishchenkoivan/hotel-app/booking-service/internal/model"
	"github.com/google/uuid"
)

type Repository interface {
	GetById(id uuid.UUID) (model.Reservation, error)
	SearchByPhone(phone string) ([]model.Reservation, error)
	Put(reserv model.ReservationData) (uuid.UUID, error)
  IsAvailible(roomId uuid.UUID, inTime time.Time, outTime time.Time) bool
}
