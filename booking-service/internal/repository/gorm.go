package repository

import (
	"time"

	"github.com/anishchenkoivan/hotel-app/booking-service/internal/model"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type GormRepository struct {
	db *gorm.DB
}

func NewGormRepository(db *gorm.DB) GormRepository {
	return GormRepository{db: db}
}

func (p GormRepository) GetById(id uuid.UUID) (model.ReservationModel, error) {
	reserv := model.ReservationModel{}
	res := p.db.Model(&model.ReservationModel{}).First(&reserv, id)
	return reserv, res.Error
}

func (p GormRepository) SearchByPhone(phone string) ([]model.ReservationModel, error) {
	var found []model.ReservationModel
	res := p.db.Model(&model.ReservationModel{}).
		Find(&found, model.ReservationModel{Reservation: model.Reservation{Client: model.Client{Phone: phone}}})
	return found, res.Error
}

func (p GormRepository) IsAvailible(roomId uuid.UUID, inTime time.Time, outTime time.Time) (bool, error) {
	var count int64
	res := p.db.Model(&model.ReservationModel{}).
		Where("room_id = ?", roomId).
		Where("in_time BETWEEN ? AND ?", inTime, outTime).
		Where("out_time BETWEEN ? AND ?", inTime, outTime).
		Count(&count)
	return count == 0, res.Error
}

func (p GormRepository) Put(data model.Reservation) (uuid.UUID, error) {
	id := uuid.New()
	reserv := model.ReservationModel{Id: id, Reservation: data}
	res := p.db.Model(&model.ReservationModel{}).Create(&reserv)
	return id, res.Error
}

func (p GormRepository) GetReservedDates(roomId uuid.UUID) ([]time.Time, error) {
	var reservs []model.ReservationModel
	res := p.db.Model(&model.ReservationModel{}).Where("room_id = ?", roomId).Find(&reservs)

	return []time.Time{}, res.Error
}

func (p GormRepository) GetRoomReservations(roomId uuid.UUID) ([]model.ReservationModel, error) {
	return make([]model.ReservationModel, 0), nil
}
