package service

import (
	"github.com/anishchenkoivan/hotel-app/notification-service/internal/client"
	"github.com/anishchenkoivan/hotel-app/notification-service/internal/model"
	"log"
)

type Service struct {
	client client.BotClient
}

func NewService(client client.BotClient) *Service {
	return &Service{client: client}
}

func (s *Service) HandleMessage(key string, message model.Message) error {
	log.Println("Message accepted: " + message.TelegramId)
	if err := s.notifyClient(message); err != nil {
		log.Println(err)
	}
	if err := s.notifyOwner(); err != nil {
		log.Println(err)
	}
	return nil
}

func (s *Service) notifyClient(message model.Message) error {
	err := s.client.SendMessage(message)
	if err != nil {
		return err
	}
	return nil
}

func (s *Service) notifyOwner() error {
	return nil
}
