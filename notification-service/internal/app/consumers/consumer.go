package consumers

import (
	"context"
	"encoding/json"
	"github.com/anishchenkoivan/hotel-app/notification-service/internal/model"
	"github.com/anishchenkoivan/hotel-app/notification-service/internal/service"
	"github.com/segmentio/kafka-go"
	"log"
)

type Consumer struct {
	reader  *kafka.Reader
	service *service.Service
}

func NewConsumer(reader *kafka.Reader, service *service.Service) *Consumer {
	return &Consumer{reader: reader, service: service}
}

func (c *Consumer) Consume(ctx context.Context) {
	defer func(reader *kafka.Reader) {
		err := reader.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(c.reader)

	for {
		select {
		case <-ctx.Done():
			log.Println("Consumer shutting down")
			return
		default:
			messageData, err := c.reader.ReadMessage(context.Background())
			if err != nil {
				log.Printf("Error reading message: %s\n", err)
				continue
			}

			var message model.Message
			if err := json.Unmarshal(messageData.Value, &message); err != nil {
				log.Printf("Error unmarshaling message: %s\n", err)
				continue
			}

			if err := c.service.HandleMessage(string(messageData.Key), message); err != nil {
				log.Printf("Error handling message: %s\n", err)
			}
		}
	}
}
