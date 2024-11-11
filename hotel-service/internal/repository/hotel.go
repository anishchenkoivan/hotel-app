package repository

import (
	"errors"
	"github.com/anishchenkoivan/hotel-app/hotel-service/internal/apperrors"
	"github.com/anishchenkoivan/hotel-app/hotel-service/internal/model"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type PostgresHotelRepository struct {
	db *gorm.DB
}

func NewPostgresHotelRepository(db *gorm.DB) *PostgresHotelRepository {
	return &PostgresHotelRepository{
		db: db,
	}
}

func (p *PostgresHotelRepository) Get(id uuid.UUID) (*model.Hotel, error) {
	var hotel model.Hotel
	if err := p.db.First(&hotel, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, apperrors.NewNotFoundError("Failed to find hotel with id=" + id.String())
		}
		return nil, err
	}
	return &hotel, nil
}

func (p *PostgresHotelRepository) GetAll() ([]*model.Hotel, error) {
	var hotels []*model.Hotel
	if err := p.db.Find(&hotels).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, apperrors.NewNotFoundError("Failed to find any hotels")
		}
		return nil, err
	}
	return hotels, nil
}

func (p *PostgresHotelRepository) Put(hotel *model.Hotel) error {
	if err := p.db.Save(hotel).Error; err != nil {
		return err
	}
	return nil
}

func (p *PostgresHotelRepository) Update(hotel *model.Hotel) error {
	if err := p.db.Save(hotel).Error; err != nil {
		return err
	}
	return nil
}

func (p *PostgresHotelRepository) Delete(id uuid.UUID) error {
	if err := p.db.Delete(&model.Hotel{}, id).Error; err != nil {
		return err
	}
	return nil
}
