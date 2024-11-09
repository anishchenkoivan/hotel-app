package main

import (
	"context"
	"fmt"
	"log"
	"os/signal"
	"syscall"

	"github.com/anishchenkoivan/hotel-app/booking-service/internal/app"
	"github.com/anishchenkoivan/hotel-app/booking-service/internal/repository"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	dsn := "host=localhost user=postgres password=postgres dbname=postgres port=5432"
	config := gorm.Config{}
	dialector := postgres.Open(dsn)
	db, err := gorm.Open(dialector, &config)

	if err != nil {
		log.Fatal("Unable to open postgres db:", err)
	}

	repo := repository.NewGormRepository(db)

	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	app := app.NewBookingServiceApp(repo)
	app.Start(ctx)
}
