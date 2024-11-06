package main

import (
	"github.com/anishchenkoivan/hotel-app/hotel-service/internal/app"
	"log"
)

func main() {
	hotelApp := app.NewHotelServiceApp()

	if err := hotelApp.Start(); err != nil {
		log.Fatal("App start failed")
	}
}
