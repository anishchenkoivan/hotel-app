package mocks

import "github.com/google/uuid"

type MockHotelService struct {
	GetPriceReturnValue int64
	GetPriceReturnError error
	LastId      uuid.UUID
}

func (h *MockHotelService) GetPrice(id uuid.UUID) (int64, error) {
	h.LastId = id
	return h.GetPriceReturnValue, h.GetPriceReturnError
}
