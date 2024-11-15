package service

import (
	"fmt"
	dto "github.com/anishchenkoivan/hotel-app/hotel-service/internal/app/handlers"
	"github.com/anishchenkoivan/hotel-app/hotel-service/internal/model"
	"github.com/google/uuid"
)

type HotelierService struct {
	hotelierRepository HotelierRepository
}

func NewHotelierService(repository HotelierRepository) *HotelierService {
	return &HotelierService{hotelierRepository: repository}
}

func (service *HotelierService) GetHotelierById(id uuid.UUID) (*dto.HotelierDto, error) {
	hotelier, err := service.hotelierRepository.Get(id)
	if err != nil {
		return nil, fmt.Errorf("GetHotelierById: %w", err)
	}
	return dto.HotelierDtoFromEntity(hotelier), nil
}

func (service *HotelierService) GetAllHoteliers() ([]*dto.HotelierDto, error) {
	hoteliers, err := service.hotelierRepository.GetAll()
	if err != nil {
		return nil, fmt.Errorf("GetAllHoteliers: %w", err)
	}
	hotelierDtoList := make([]*dto.HotelierDto, len(hoteliers))
	for i, hotelier := range hoteliers {
		hotelierDtoList[i] = dto.HotelierDtoFromEntity(hotelier)
	}
	return hotelierDtoList, nil
}

func (service *HotelierService) CreateHotelier(hotelierData dto.HotelierModifyDto) (uuid.UUID, error) {
	hotelier := model.Hotelier{
		ID:       uuid.Nil,
		Username: hotelierData.Username,
	}

	id, err := service.hotelierRepository.Put(&hotelier)
	if err != nil {
		return uuid.Nil, fmt.Errorf("CreateHotelier: %w", err)
	}

	return id, nil
}

func (service *HotelierService) UpdateHotelier(id uuid.UUID, hotelierData dto.HotelierModifyDto) error {
	hotelier, err := service.hotelierRepository.Get(id)
	if err != nil {
		return fmt.Errorf("UpdateHotelier: %w", err)
	}

	hotelier.Username = hotelierData.Username

	err = service.hotelierRepository.Update(hotelier)
	if err != nil {
		return fmt.Errorf("UpdateHotelier: %w", err)
	}
	return nil
}

func (service *HotelierService) DeleteHotelier(id uuid.UUID) error {
	return service.hotelierRepository.Delete(id)
}
