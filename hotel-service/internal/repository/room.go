package repository

import (
	"errors"
	"github.com/anishchenkoivan/hotel-app/hotel-service/internal/apperrors"
	"github.com/anishchenkoivan/hotel-app/hotel-service/internal/model"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type PostgresRoomRepository struct {
	db *gorm.DB
}

func NewPostgresRoomRepository(db *gorm.DB) *PostgresRoomRepository {
	return &PostgresRoomRepository{
		db: db,
	}
}

func (p *PostgresRoomRepository) Get(id uuid.UUID) (*model.Room, error) {
	var room model.Room
	if err := p.db.First(&room, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, apperrors.NewNotFoundError("Failed to find room with id=" + id.String())
		}
		return nil, err
	}
	return &room, nil
}

func (p *PostgresRoomRepository) GetAll() ([]*model.Room, error) {
	var rooms []*model.Room
	if err := p.db.Find(&rooms).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, apperrors.NewNotFoundError("Failed to find any rooms")
		}
		return nil, err
	}
	return rooms, nil
}

func (p *PostgresRoomRepository) Put(room *model.Room) (uuid.UUID, error) {
	if err := p.db.Save(room).Error; err != nil {
		return uuid.Nil, err
	}
	return room.ID, nil
}

func (p *PostgresRoomRepository) Update(room *model.Room) error {
	if err := p.db.Save(room).Error; err != nil {
		return err
	}
	return nil
}

func (p *PostgresRoomRepository) Delete(id uuid.UUID) error {
	if err := p.db.Delete(&model.Room{}, id).Error; err != nil {
		return err
	}
	return nil
}
