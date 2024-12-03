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

func (s Service) AddReservation(data model.ReservationData) (uuid.UUID, *BookErr) {
	free, err := s.repository.IsAvailible(data.RoomId, data.InTime, data.OutTime)

	if err != nil {
		return uuid.UUID{}, &BookErr{error: err, ErrType: RepositoryError}
	}

	if !free {
		err := NewReservationAlreadyExistsError()
		return uuid.UUID{}, err
	}

	id, err := s.repository.Put(data)

	if err != nil {
		return uuid.UUID{}, &BookErr{error: err, ErrType: RepositoryError}
	}

	return id, nil
}

func (s Service) GetRoomReservations(roomId uuid.UUID) ([]model.Reservation, error) {
	res, err := s.repository.GetRoomReservations(roomId)
	return res, &BookErr{error: err, ErrType: RepositoryError}
}
