package service

import "github.com/anishchenkoivan/hotel-app/hotel-service/internal/repository"

type HotelierService struct {
	repository HotelierRepository
}

func NewHotelierService() HotelierService {
	return HotelierService{repository.NewPostgresHotelierRepository()}
}
