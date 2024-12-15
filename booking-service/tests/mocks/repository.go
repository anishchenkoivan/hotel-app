package mocks

import (
	"time"

	"github.com/anishchenkoivan/hotel-app/booking-service/internal/model"
	"github.com/google/uuid"
)

type MockRepository struct {
	LastPutted             *model.Reservation
	PutReturnValue         uuid.UUID
	PutReturnError         error
	IsAvailableReturnValue bool
	IsAvailableReturnError error
}

func NewMockRepository() *MockRepository {
	return &MockRepository{}
}

func (p MockRepository) GetById(id uuid.UUID) (model.ReservationModel, error) {
	return model.ReservationModel{}, nil
}

func (p MockRepository) SearchByPhone(phone string) ([]model.ReservationModel, error) {
	return []model.ReservationModel{}, nil
}

func (p MockRepository) IsAvailable(roomId uuid.UUID, inTime time.Time, outTime time.Time) (bool, error) {
	return p.IsAvailableReturnValue, p.IsAvailableReturnError
}

func (p *MockRepository) Put(data model.Reservation) (uuid.UUID, error) {
	p.LastPutted = &data
	return p.PutReturnValue, p.PutReturnError
}

func (p MockRepository) GetReservedDates(roomId uuid.UUID) ([]time.Time, error) {
	return []time.Time{}, nil
}

func (p MockRepository) GetRoomReservations(roomId uuid.UUID) ([]model.ReservationModel, error) {
	return []model.ReservationModel{}, nil
}

func (p *MockRepository) ConfirmPayment(id uuid.UUID) error {
	return nil
}

func (p *MockRepository) RemoveReservation(id uuid.UUID) error {
	return nil
}
