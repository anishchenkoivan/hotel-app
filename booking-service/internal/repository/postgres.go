package repository

import (
	"fmt"
	"github.com/anishchenkoivan/hotel-app/booking-service/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewPostgresRepository(conf config.DbConfig) (GormRepository, error) {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s",
		conf.Host, conf.User, conf.Password, conf.Name, conf.Port)
	gorm_cfg := gorm.Config{}
	db, err := gorm.Open(postgres.Open(dsn), &gorm_cfg)

	if err != nil {
		return GormRepository{}, err
	}

	repo := NewGormRepository(db)

	return repo, nil
}
