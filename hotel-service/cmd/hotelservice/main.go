package main

import (
	"github.com/anishchenkoivan/hotel-app/hotel-service/config"
	"github.com/anishchenkoivan/hotel-app/hotel-service/internal/app"
	"log"
)

func main() {
	hotelApp := app.NewHotelServiceApp(*config.NewConfig())

	if err := hotelApp.Start(); err != nil {
		log.Fatal("App start failed")
	}
}
