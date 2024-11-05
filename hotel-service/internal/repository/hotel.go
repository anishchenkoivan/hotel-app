package repository

import (
	"github.com/anishchenkoivan/hotel-app/hotel-service/internal/model"
	"github.com/google/uuid"
)

type PostgresHotelRepository struct{}

func (p PostgresHotelRepository) Get(id uuid.UUID) (*model.Hotel, error) {
	//TODO implement me
	panic("implement me")
}

func (p PostgresHotelRepository) GetAll() ([]*model.Hotel, error) {
	//TODO implement me
	panic("implement me")
}

func (p PostgresHotelRepository) Put(hotel *model.Hotel) error {
	//TODO implement me
	panic("implement me")
}

func (p PostgresHotelRepository) Remove(id uuid.UUID) error {
	//TODO implement me
	panic("implement me")
}

func NewPostgresHotelRepository() PostgresHotelRepository {
	return PostgresHotelRepository{}
}
