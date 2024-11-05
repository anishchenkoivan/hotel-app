package service

import (
	"github.com/anishchenkoivan/hotel-app/booking-service/internal/repository"
)

type Service struct {
	repository Repository
}

func NewService() Service {
	return Service{repository.NewPostgresRepository()}
}
