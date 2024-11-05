package main

import (
	"fmt"
	"github.com/anishchenkoivan/hotel-app/hotel-service/internal/app"
)

func main() {
	hotelApp := app.NewHotelServiceApp()
	if err := hotelApp.Start(); err != nil {
		fmt.Println(err)
	}

	if err := hotelApp.Start(); err != nil {
		return
	}
}
