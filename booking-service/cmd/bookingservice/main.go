package main

import (
	"context"
	"log"
	"os/signal"
	"syscall"

	"github.com/anishchenkoivan/hotel-app/booking-service/config"
	"github.com/anishchenkoivan/hotel-app/booking-service/internal/app"
	"github.com/anishchenkoivan/hotel-app/booking-service/internal/repository"
)

func main() {
	conf, err := config.NewConfig()

	if err != nil {
		log.Fatal("Unable to create config: ", err)
	}

	repo, err := repository.NewPostgresRepository(conf.Db)

  if err != nil {
		log.Fatal("Unable to create repository: ", err)
	}

	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	app := app.NewBookingServiceApp(repo, conf)
	app.Start(ctx)
}
