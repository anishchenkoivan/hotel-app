package main

import (
	"github.com/anishchenkoivan/hotel-app/notification-service/config"
	"github.com/anishchenkoivan/hotel-app/notification-service/internal/app"
	"log"
)

func main() {
	notificationApp := app.NewNotificationApp(*config.NewConfig())

	if err := notificationApp.Start(); err != nil {
		log.Fatal("App start failed")
	}
}
