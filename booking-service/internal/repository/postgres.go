package repository

import (
	"fmt"

	"github.com/anishchenkoivan/hotel-app/booking-service/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewPostgresRepository(cfg config.Config) (GormRepository, error) {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s",
		cfg.DbHost, cfg.DbUser, cfg.DbPassword, cfg.DbName, cfg.DbPort)
	gorm_cfg := gorm.Config{}
	dialector := postgres.Open(dsn)
	db, err := gorm.Open(dialector, &gorm_cfg)
	return NewGormRepository(db, cfg), err
}
