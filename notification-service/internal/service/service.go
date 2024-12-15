package service

import (
	"github.com/anishchenkoivan/hotel-app/notification-service/internal/model"
	"log"
)

type Service struct{}

func NewService() *Service {
	return &Service{}
}

func (s *Service) HandleMessage(key string, message model.Message) error {
	log.Println("Message accepted")
	if err := s.notifyClient(); err != nil {
		log.Println(err)
	}
	if err := s.notifyOwner(); err != nil {
		log.Println(err)
	}
	return nil
}

func (s *Service) notifyClient() error {
	return nil
}

func (s *Service) notifyOwner() error {
	return nil
}
