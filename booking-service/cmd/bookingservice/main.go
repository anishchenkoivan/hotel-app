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
		log.Fatal("Unable to create config: ", err)
	}

	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	app, err := app.NewBookingServiceApp(conf)

  if err != nil {
    log.Fatal("Unable to create app: ", err)
  }

	app.Start(ctx)
}
