package hotel_mocks

import (
	"github.com/anishchenkoivan/hotel-app/hotel-service/internal/model"
	"github.com/google/uuid"
)

type MockHotelRepository struct {
	GetReturnValue *model.Hotel
	GetReturnError error

	GetAllReturnValue []*model.Hotel
	GetAllReturnError error

	PutReturnValue uuid.UUID
	PutReturnError error
	LastPut        *model.Hotel
}

func (p *MockHotelRepository) Get(id uuid.UUID) (*model.Hotel, error) {
	return p.GetReturnValue, p.GetReturnError
}

func (p *MockHotelRepository) GetAll() ([]*model.Hotel, error) {
	return p.GetAllReturnValue, p.GetAllReturnError
}

func (p *MockHotelRepository) Put(hotel *model.Hotel) (uuid.UUID, error) {
  p.LastPut = hotel
	return p.PutReturnValue, p.PutReturnError
}

func (p *MockHotelRepository) Update(hotel *model.Hotel) error {
	return nil
}

func (p *MockHotelRepository) Delete(id uuid.UUID) error {
	return nil
}
