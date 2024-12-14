package service

import (
	"fmt"

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

func (s Service) AddReservation(data model.Reservation) (uuid.UUID, error) {
	if !data.InTime.Before(data.OutTime) {
		return uuid.UUID{}, InvalidReservation
	}

	free, err := s.repository.IsAvailable(data.RoomId, data.InTime, data.OutTime)

	if err != nil {
		return uuid.UUID{}, fmt.Errorf("%w: %w", RepositoryError, err)
	}

	if !free {
		return uuid.UUID{}, ReservationAlreadyExists
	}

	data.Cost, err = s.hotelService.GetPrice(data.RoomId)

	if err != nil {
		return uuid.UUID{}, fmt.Errorf("%w: %w", GrpcError, err)
	}

	id, err := s.repository.Put(data)

	if err != nil {
		return uuid.UUID{}, fmt.Errorf("%w: %w", RepositoryError, err)
	}

	return id, nil
}

func (s Service) GetRoomReservations(roomId uuid.UUID) ([]model.ReservationModel, error) {
	return s.repository.GetRoomReservations(roomId)
}
