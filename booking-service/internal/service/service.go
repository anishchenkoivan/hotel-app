package service

import (
	"fmt"

	"github.com/anishchenkoivan/hotel-app/booking-service/internal/model"
	"github.com/google/uuid"
)

type Service struct {
	repository    Repository
	hotelService  HotelService
	paymentSystem PaymentSystem
}

func NewService(repo Repository, hotel HotelService, payment PaymentSystem) *Service {
	return &Service{repo, hotel, payment}
}

func (s Service) GetById(id uuid.UUID) (model.ReservationModel, error) {
	return s.repository.GetById(id)
}

func (s Service) SearchByPhone(phone string) ([]model.ReservationModel, error) {
	return s.repository.SearchByPhone(phone)
}

func (s Service) AddReservation(data model.Reservation) (uuid.UUID, string, error) {
	if !data.InTime.Before(data.OutTime) {
		return uuid.UUID{}, "", InvalidReservation
	}

	free, err := s.repository.IsAvailable(data.RoomId, data.InTime, data.OutTime)

	if err != nil {
		return uuid.UUID{}, "", fmt.Errorf("%w: %w", RepositoryError, err)
	}

	if !free {
		return uuid.UUID{}, "", ReservationAlreadyExists
	}

	data.Cost, err = s.hotelService.GetPrice(data.RoomId)

	if err != nil {
		return uuid.UUID{}, "", fmt.Errorf("%w: %w", HotelServiceError, err)
	}

	id, err := s.repository.Put(data)

	if err != nil {
		return uuid.UUID{}, "", fmt.Errorf("%w: %w", RepositoryError, err)
	}

  payURL, err := s.paymentSystem.AddPayment(id, data.Cost)

  if err != nil {
		return uuid.UUID{}, "", fmt.Errorf("%w: %w", PayemntSystemError, err)
	}

	return id, payURL, nil
}

func (s Service) GetRoomReservations(roomId uuid.UUID) ([]model.ReservationModel, error) {
	return s.repository.GetRoomReservations(roomId)
}

func (s Service) ConfirmPayment(id uuid.UUID) error {
	return s.repository.ConfirmPayment(id)
}

func (s Service) CancelReservation(id uuid.UUID) error {
	return s.repository.RemoveReservation(id)
}
