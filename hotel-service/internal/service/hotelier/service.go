package service

import (
	"fmt"
	"github.com/anishchenkoivan/hotel-app/hotel-service/internal/model"
	"github.com/google/uuid"
)

type HotelierService struct {
	repository HotelierRepository
}

func NewHotelierService(repository HotelierRepository) *HotelierService {
	return &HotelierService{repository: repository}
}

func (service *HotelierService) GetHotelierById(id uuid.UUID) (*model.Hotelier, error) {
	return service.repository.Get(id)
}

func (service *HotelierService) GetAllHoteliers() ([]*model.Hotelier, error) {
	return service.repository.GetAll()
}

func (service *HotelierService) CreateHotelier(hotelierData model.HotelierData) (uuid.UUID, error) {
	id, err := uuid.NewRandom()
	if err != nil {
		return uuid.Nil, fmt.Errorf("CreateHotelier: %w", err)
	}

	hotelier := model.Hotelier{
		ID:           id,
		HotelierData: hotelierData,
	}

	err = service.repository.Put(&hotelier)
	if err != nil {
		return uuid.Nil, fmt.Errorf("CreateHotelier: %w", err)
	}

	return id, nil
}

func (service *HotelierService) UpdateHotelier(id uuid.UUID, hotelierData model.HotelierData) error {
	hotelier, err := service.repository.Get(id)
	if err != nil {
		return fmt.Errorf("UpdateHotelier: %w", err)
	}

	hotelier.HotelierData = hotelierData
	err = service.repository.Update(hotelier)
	if err != nil {
		return fmt.Errorf("UpdateHotelier: %w", err)
	}
	return nil
}

func (service *HotelierService) DeleteHotelier(id uuid.UUID) error {
	return service.repository.Delete(id)
}
