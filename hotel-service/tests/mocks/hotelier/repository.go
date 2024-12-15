package hotelier_mocks

import (
	"github.com/anishchenkoivan/hotel-app/hotel-service/internal/model"
	"github.com/google/uuid"
)

type MockHotelierRepository struct {
	GetReturnValue *model.Hotelier
	GetReturnError error
}

func (p *MockHotelierRepository) Get(id uuid.UUID) (*model.Hotelier, error) {
	return p.GetReturnValue, p.GetReturnError
}

func (p *MockHotelierRepository) GetByTelegramId(telegramId string) (*model.Hotelier, error) {
	return nil, nil
}

func (p *MockHotelierRepository) GetAll() ([]*model.Hotelier, error) {
	return []*model.Hotelier{}, nil
}

func (p *MockHotelierRepository) Put(hotelier *model.Hotelier) (uuid.UUID, error) {
	return uuid.New(), nil
}

func (p *MockHotelierRepository) Update(hotelier *model.Hotelier) error {
	return nil
}

func (p *MockHotelierRepository) Delete(uuid.UUID) error {
	return nil
}
