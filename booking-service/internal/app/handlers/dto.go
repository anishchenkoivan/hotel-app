package handlers

import (
	"github.com/anishchenkoivan/hotel-app/booking-service/internal/model"
	"github.com/google/uuid"
)

type GetByIdQuery struct {
	Id uuid.UUID
}

type SearchByPhoneQuery struct {
	phone string
}

type AddReservationQuery struct {
  model.ReservationData
}

type GetByIdResponse struct {
  model.Reservation
}

type SearchByPhoneResponse struct {
  reservations []model.Reservation
}

type AddReservationResponse struct {
  Id uuid.UUID
}
