package service

import (
	"fmt"
	"github.com/anishchenkoivan/hotel-app/hotel-service/internal/model"
	"github.com/google/uuid"
)

type HotelService struct {
	repository HotelRepository
}

func NewHotelService(repository HotelRepository) *HotelService {
	return &HotelService{repository: repository}
}

func (service *HotelService) GetHotelById(id uuid.UUID) (*model.Hotel, error) {
	hotel, err := service.repository.Get(id)
	if err != nil {
		return nil, fmt.Errorf("GetHotelById: %w", err)
	}
	return hotel, nil
}

func (service *HotelService) GetAllHotels() ([]*model.Hotel, error) {
	hotels, err := service.repository.GetAll()
	if err != nil {
		return nil, fmt.Errorf("GetAllHotels: %w", err)
	}
	return hotels, nil
}

func (service *HotelService) CreateHotel(hotelData model.HotelData) (uuid.UUID, error) {
	id, err := uuid.NewRandom()
	if err != nil {
		return uuid.Nil, fmt.Errorf("CreateHotel: %w", err)
	}

	hotel := model.Hotel{
		ID:        id,
		HotelData: hotelData,
	}

	err = service.repository.Put(&hotel)
	if err != nil {
		return uuid.Nil, fmt.Errorf("CreateHotel: %w", err)
	}

	return id, nil
}

func (service *HotelService) UpdateHotel(id uuid.UUID, hotelData model.HotelData) error {
	hotel, err := service.repository.Get(id)
	if err != nil {
		return fmt.Errorf("UpdateHotel: %w", err)
	}

	hotel.HotelData = hotelData
	err = service.repository.Update(hotel)
	if err != nil {
		return fmt.Errorf("UpdateHotel: %w", err)
	}
	return nil
}

func (service *HotelService) DeleteHotel(id uuid.UUID) error {
	return service.repository.Delete(id)
}
