package service

import (
	"errors"

	"github.com/anishchenkoivan/hotel-app/booking-service/internal/clients"
	"github.com/anishchenkoivan/hotel-app/booking-service/internal/model"
	"github.com/google/uuid"
)

type Service struct {
	repository   Repository
	hotelService clients.HotelService
}

func NewService(repo Repository, hotel clients.HotelService) Service {
	return Service{repo, hotel}
}

func (s Service) GetById(id uuid.UUID) (model.ReservationModel, error) {
	return s.repository.GetById(id)
}

func (s Service) SearchByPhone(phone string) ([]model.ReservationModel, error) {
	return s.repository.SearchByPhone(phone)
}

func (s Service) AddReservation(data model.Reservation) (uuid.UUID, error, ErrType) {
  if !data.InTime.Before(data.OutTime) {
    return uuid.UUID{}, errors.New("Invalid reservation"), BadReservation
  }

	free, err := s.repository.IsAvailible(data.RoomId, data.InTime, data.OutTime)

	if err != nil {
		return uuid.UUID{}, err, RepositoryError
	}

	if !free {
		return uuid.UUID{}, errors.New("Reservation already exists"), ReservationAlreadyExists
	}

	data.Cost, err = s.hotelService.GetPrice(data.RoomId)

  if err != nil {
		return uuid.UUID{}, err, GrpcError
	}

	id, err := s.repository.Put(data)

	if err != nil {
		return uuid.UUID{}, err, RepositoryError
	}

	return id, nil, NoError
}

func (s Service) GetRoomReservations(roomId uuid.UUID) ([]model.ReservationModel, error) {
	return s.repository.GetRoomReservations(roomId)
}
