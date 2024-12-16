package clients

import (
	"context"
	"encoding/json"
	"github.com/anishchenkoivan/hotel-app/booking-service/config"
	"github.com/segmentio/kafka-go"
)

type NotificationService struct {
	conf   config.KafkaConfig
	writer *kafka.Writer
}

func NewNotificationService(conf config.KafkaConfig) (*NotificationService, error) {
	writer := kafka.Writer{
		Addr:     kafka.TCP(conf.BrokerHost + ":" + conf.BrokerPort),
		Topic:    conf.Topic,
		Balancer: &kafka.LeastBytes{},
	}
	return &NotificationService{conf: conf, writer: &writer}, nil
}

func (s NotificationService) SendNotification(message Message) error {
	stringMessage, err := json.Marshal(message)
	if err != nil {
		return err
	}

	kafkaMessage := kafka.Message{
		Key:   []byte(message.TelegramId),
		Value: []byte(stringMessage),
	}

	err = s.writer.WriteMessages(context.Background(), kafkaMessage)
	if err != nil {
		return err
	}

	return nil
}
