package repository

import (
	"time"

	"github.com/anishchenkoivan/hotel-app/booking-service/config"
	"github.com/anishchenkoivan/hotel-app/booking-service/internal/model"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type GormRepository struct {
	db     *gorm.DB
	config config.Config
}

func NewGormRepository(db *gorm.DB, cfg config.Config) GormRepository {
	return GormRepository{db: db, config: cfg}
}

func (p GormRepository) GetById(id uuid.UUID) (model.Reservation, error) {
	reserv := model.Reservation{}
	res := p.db.Model(&model.Reservation{}).First(&reserv, id)
	return reserv, res.Error
}

func (p GormRepository) SearchByPhone(phone string) ([]model.Reservation, error) {
	var found []model.Reservation
	res := p.db.Model(&model.Reservation{}).
		Find(&found, model.Reservation{ReservationData: model.ReservationData{Client: model.Client{Phone: phone}}})
	return found, res.Error
}

func (p GormRepository) IsAvailible(roomId uuid.UUID, inTime time.Time, outTime time.Time) (bool, error) {
	var count int64
	res := p.db.Model(&model.Reservation{}).
		Where("room_id = ?", roomId).
		Where("in_time BETWEEN ? AND ?", inTime, outTime).
		Where("out_time BETWEEN ? AND ?", inTime, outTime).
		Count(&count)
	return count == 0, res.Error
}

func (p GormRepository) Put(data model.ReservationData) (uuid.UUID, error) {
	id := uuid.New()
	reserv := model.Reservation{Id: id, ReservationData: data}
	res := p.db.Model(&model.Reservation{}).Create(reserv)
	return id, res.Error
}

func (p GormRepository) GetReservedDates(roomId uuid.UUID) ([]time.Time, error) {
	var reservs []model.Reservation
	res := p.db.Model(&model.Reservation{}).Where("room_id = ?", roomId).Find(&reservs)

	return []time.Time{}, res.Error
}

func (p GormRepository) GetRoomReservations(roomId uuid.UUID) ([]model.Reservation, error) {
	return make([]model.Reservation, 0), nil
}
