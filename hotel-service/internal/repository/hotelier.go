package repository

import (
	"errors"
	"github.com/anishchenkoivan/hotel-app/hotel-service/internal/apperrors"
	"github.com/anishchenkoivan/hotel-app/hotel-service/internal/model"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type PostgresHotelierRepository struct {
	db *gorm.DB
}

func NewPostgresHotelierRepository(db *gorm.DB) *PostgresHotelierRepository {
	return &PostgresHotelierRepository{db: db}
}

func (p *PostgresHotelierRepository) Get(id uuid.UUID) (*model.Hotelier, error) {
	var hotelier model.Hotelier
	if err := p.db.First(&hotelier, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, apperrors.NewNotFoundError("Failed to find hotel with id=" + id.String())
		}
		return nil, err
	}
	return &hotelier, nil
}

func (p *PostgresHotelierRepository) GetAll() ([]*model.Hotelier, error) {
	var hoteliers []*model.Hotelier
	if err := p.db.Find(&hoteliers).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, apperrors.NewNotFoundError("Failed to find any hotels")
		}
		return nil, err
	}
	return hoteliers, nil
}

func (p *PostgresHotelierRepository) Put(hotelier *model.Hotelier) error {
	if err := p.db.Save(hotelier).Error; err != nil {
		return err
	}
	return nil
}

func (p *PostgresHotelierRepository) Update(hotelier *model.Hotelier) error {
	if err := p.db.Save(hotelier).Error; err != nil {
		return err
	}
	return nil
}

func (p *PostgresHotelierRepository) Delete(uuid.UUID) error {
	if err := p.db.Delete(&model.Hotelier{}, &model.Hotelier{}).Error; err != nil {
		return err
	}
	return nil
}
