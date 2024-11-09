package main

import (
	"context"
	"github.com/anishchenkoivan/hotel-app/booking-service/internal/app"
	"os/signal"
	"syscall"
)

func main() {
	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	app := app.NewBookingServiceApp()
	app.Start(ctx)
}
