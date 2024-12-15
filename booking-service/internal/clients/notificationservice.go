package clients

import "github.com/anishchenkoivan/hotel-app/booking-service/config"

type NotificationService struct {
	conf config.KafkaConfig
}

func NewNotificationService(conf config.KafkaConfig) (*NotificationService, error) {
	return &NotificationService{conf: conf}, nil
}

func (s NotificationService) SendNotification(message Message) error {
	return nil
}
