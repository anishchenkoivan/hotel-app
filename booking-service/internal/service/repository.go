package service

import (
	"time"

	"github.com/anishchenkoivan/hotel-app/booking-service/internal/model"
	"github.com/google/uuid"
)

type Repository interface {
	GetById(id uuid.UUID) (model.ReservationModel, error)
	SearchByPhone(phone string) ([]model.ReservationModel, error)
	Put(reserv model.Reservation) (uuid.UUID, error)
  IsAvailable(roomId uuid.UUID, inTime time.Time, outTime time.Time) (bool, error)
  GetRoomReservations(roomId uuid.UUID) ([]model.ReservationModel, error)
}
