package service

import "github.com/anishchenkoivan/hotel-app/hotel-service/internal/repository"

type HotelService struct {
	repository HotelRepository
}

func NewHotelService() HotelService {
	return HotelService{repository: repository.NewPostgresHotelRepository()}
}
