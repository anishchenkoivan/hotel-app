package main

import (
	"context"
	"fmt"
	"log"
	"os/signal"
	"syscall"

	"github.com/anishchenkoivan/hotel-app/booking-service/config"
	"github.com/anishchenkoivan/hotel-app/booking-service/internal/app"
	"github.com/anishchenkoivan/hotel-app/booking-service/internal/model"
	"github.com/anishchenkoivan/hotel-app/booking-service/internal/repository"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	conf, err := config.NewConfig()

	if err != nil {
		log.Fatal("Unable to create config: ", err)
	}

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s",
		conf.Db.Host, conf.Db.User, conf.Db.Password, conf.Db.Name, conf.Db.Port)
	gorm_cfg := gorm.Config{}
	db, err := gorm.Open(postgres.Open(dsn), &gorm_cfg)

	if err != nil {
		log.Fatal("Unable to create postgres repository: ", err)
	}

	db.AutoMigrate(&model.Reservation{})

	repo, err := repository.NewGormRepository(db, conf), err

	if err != nil {
		log.Fatal("Unable to create postgres repository: ", err)
	}

	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	app := app.NewBookingServiceApp(repo, conf)
	app.Start(ctx)
}
