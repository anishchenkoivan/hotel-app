package mocks

import "github.com/google/uuid"

type MockHotelService struct {
	GetPriceReturnValue int64
	GetPriceReturnError error
	LastId              uuid.UUID
}

func (h *MockHotelService) GetPrice(id uuid.UUID) (int64, error) {
	h.LastId = id
	return h.GetPriceReturnValue, h.GetPriceReturnError
}

type MockPaymentSystem struct {
	AddPaymentReturnValue string
	AddPaymentReturnError error
	LastId                uuid.UUID
}

func (p *MockPaymentSystem) AddPayment(id uuid.UUID, cost int64) (string, error) {
  p.LastId = id
  return p.AddPaymentReturnValue, p.AddPaymentReturnError
}
