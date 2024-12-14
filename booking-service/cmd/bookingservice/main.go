package main

import (
	"context"
	"log"
	"os/signal"
	"syscall"

	"github.com/anishchenkoivan/hotel-app/booking-service/config"
	"github.com/anishchenkoivan/hotel-app/booking-service/internal/app"
)

func main() {
	conf, err := config.NewConfig()

	if err != nil {
		log.Fatal("Can't create config: ", err)
	}

	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	app, err := app.NewBookingServiceApp(conf)

	if err != nil {
		log.Fatal("Can't initialize app: ", err)
	}

  err = app.Start(ctx)

  if err != nil {
		log.Fatal("Can't start app: ", err)
  }
}
