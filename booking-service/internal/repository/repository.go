package repository

import (
	"github.com/anishchenkoivan/hotel-app/booking-service/internal/model"
	"github.com/google/uuid"
)

type PostgresRepository struct{}

func (p PostgresRepository) GetById(id uuid.UUID) (*model.Reservation, error) {
	//TODO implement me
	panic("implement me")
}

func (p PostgresRepository) GetByPhone(phone string) (*model.Reservation, error) {
	//TODO implement me
	panic("implement me")
}

func (p PostgresRepository) GetByName(name string, surname string) (*model.Reservation, error) {
	//TODO implement me
	panic("implement me")
}

func (p PostgresRepository) Put(reserv model.Reservation) error {
	//TODO implement me
	panic("implement me")
}

func NewPostgresRepository() PostgresRepository {
  return PostgresRepository{}
}
