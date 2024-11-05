package repository

import (
	"github.com/anishchenkoivan/hotel-app/hotel-service/internal/model"
	"github.com/google/uuid"
)

type PostgresRoomRepository struct{}

func (p PostgresRoomRepository) Get(id uuid.UUID) (*model.Room, error) {
	//TODO implement me
	panic("implement me")
}

func (p PostgresRoomRepository) GetAll() ([]*model.Room, error) {
	//TODO implement me
	panic("implement me")
}

func (p PostgresRoomRepository) Put(room *model.Room) error {
	//TODO implement me
	panic("implement me")
}

func (p PostgresRoomRepository) Remove(id uuid.UUID) error {
	//TODO implement me
	panic("implement me")
}

func NewPostgresRoomRepository() PostgresRoomRepository {
	return PostgresRoomRepository{}
}
