package repository

import (
	"github.com/anishchenkoivan/hotel-app/hotel-service/internal/model"
	"github.com/google/uuid"
)

type PostgresHotelierRepository struct{}

func (p PostgresHotelierRepository) Get(id uuid.UUID) (*model.Hotelier, error) {
	//TODO implement me
	panic("implement me")
}

func (p PostgresHotelierRepository) GetAll() ([]*model.Hotelier, error) {
	//TODO implement me
	panic("implement me")
}

func (p PostgresHotelierRepository) Put(hotelier *model.Hotelier) error {
	//TODO implement me
	panic("implement me")
}

func (p PostgresHotelierRepository) Remove(hotelier *model.Hotelier) error {
	//TODO implement me
	panic("implement me")
}

func NewPostgresHotelierRepository() PostgresHotelierRepository {
	return PostgresHotelierRepository{}
}
