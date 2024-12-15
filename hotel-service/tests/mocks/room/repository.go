package room_mocks

import (
	"github.com/anishchenkoivan/hotel-app/hotel-service/internal/model"
	"github.com/google/uuid"
)

type MockRoomRepository struct {
  GetReturnValue *model.Room
  GetReturnError error

  GetAllReturnValue []*model.Room
  GetAllReturnError error
}

func (p *MockRoomRepository) Get(id uuid.UUID) (*model.Room, error) {
	return p.GetReturnValue, p.GetReturnError
}

func (p *MockRoomRepository) GetAll() ([]*model.Room, error) {
	return p.GetAllReturnValue, p.GetAllReturnError
}

func (p *MockRoomRepository) Put(room *model.Room) (uuid.UUID, error) {
	return uuid.New(), nil
}

func (p *MockRoomRepository) Update(room *model.Room) error {
	return nil
}

func (p *MockRoomRepository) Delete(id uuid.UUID) error {
	return nil
}
