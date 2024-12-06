package main

import (
	"github.com/anishchenkoivan/hotel-app/payment-system/config"
	"github.com/anishchenkoivan/hotel-app/payment-system/internal/app"
	"log"
)

func main() {
	paymentSystem := app.NewPaymentSystemApp(*config.NewConfig())

	if err := paymentSystem.Start(); err != nil {
		log.Fatal("Payment System start failed")
	}
}
