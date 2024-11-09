package repository

import (
	"github.com/anishchenkoivan/hotel-app/booking-service/internal/model"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type GormRepository struct {
	db *gorm.DB
}

func NewGormRepository(db *gorm.DB) GormRepository {
	db.AutoMigrate(&model.Reservation{})
	return GormRepository{db: db}
}

func (p GormRepository) GetById(id uuid.UUID) (model.Reservation, error) {
	reserv := model.Reservation{}
	res := p.db.Model(&model.Reservation{}).First(&reserv, id)
	return reserv, res.Error
}

func (p GormRepository) SearchByPhone(phone string) ([]model.Reservation, error) {
	var found []model.Reservation
  res := p.db.Model(&model.Reservation{}).Find(&found, model.Reservation{ReservationData: model.ReservationData{Client: model.Client{Phone: phone}}})
	return found, res.Error
}

func (p GormRepository) Put(data model.ReservationData) (uuid.UUID, error) {
	id := uuid.New()
	reserv := model.Reservation{Id: id, ReservationData: data}
	res := p.db.Model(&model.Reservation{}).Create(reserv)
	return id, res.Error
}
