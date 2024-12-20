package repository

import (
	"os"
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

func (p GormRepository) Migrate(path string) error {
  _ = p.db.AutoMigrate(&model.ReservationModel{})
	migration, err := os.ReadFile(path)

	if err != nil {
		return err
	}

	if err := p.db.Exec(string(migration)).Error; err != nil {
		return err
	}

	return nil
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

func (p GormRepository) IsAvailable(roomId uuid.UUID, inTime time.Time, outTime time.Time) (bool, error) {
	var count int64
	res := p.db.Model(&model.ReservationModel{}).
		Where("room_id = ?", roomId).
		Where("out_time >= ? and ? >= in_time", inTime, outTime).
		Count(&count)
	return count == 0, res.Error
}

func (p GormRepository) Put(data model.Reservation) (uuid.UUID, error) {
	reserv := model.ReservationModel{Reservation: data}
	res := p.db.Model(&model.ReservationModel{}).Create(&reserv)
	return reserv.ID, res.Error
}

func (p GormRepository) GetReservedDates(roomId uuid.UUID) ([]time.Time, error) {
	var reservs []model.ReservationModel
	res := p.db.Model(&model.ReservationModel{}).Where("room_id = ?", roomId).Find(&reservs)

	return []time.Time{}, res.Error
}

func (p GormRepository) GetRoomReservations(roomId uuid.UUID) ([]model.ReservationModel, error) {
	return make([]model.ReservationModel, 0), nil
}

func (p GormRepository) ConfirmPayment(id uuid.UUID) error {
	res := p.db.Model(&model.ReservationModel{}).
		Where("id = ?", id).
		Update("is_paid", "TRUE")
	return res.Error
}

func (p GormRepository) RemoveReservation(id uuid.UUID) error {
	res := p.db.Model(&model.ReservationModel{}).Delete("id = ?", id)
	return res.Error
}
