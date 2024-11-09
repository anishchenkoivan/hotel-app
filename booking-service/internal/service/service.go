package service

import (
	"github.com/anishchenkoivan/hotel-app/booking-service/internal/model"
	"github.com/google/uuid"
)

type Service struct {
	repository Repository
}

func NewService(repo Repository) Service {
	return Service{repo}
}

func (s Service) GetById(id uuid.UUID) (model.Reservation, error) {
	return s.repository.GetById(id)
}

func (s Service) SearchByPhone(phone string) ([]model.Reservation, error) {
	return s.repository.SearchByPhone(phone)
}

func (s Service) AddReservation(data model.ReservationData) (uuid.UUID, error) {
	return s.repository.Put(data)
}
