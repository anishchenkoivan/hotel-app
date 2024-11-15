package service

import (
	"fmt"
	dto "github.com/anishchenkoivan/hotel-app/hotel-service/internal/app/handlers"
	"github.com/anishchenkoivan/hotel-app/hotel-service/internal/model"
	hotelierrepository "github.com/anishchenkoivan/hotel-app/hotel-service/internal/service/hotelier"
	"github.com/google/uuid"
)

type HotelService struct {
	hotelRepository    HotelRepository
	hotelierRepository hotelierrepository.HotelierRepository
}

func NewHotelService(repository HotelRepository) *HotelService {
	return &HotelService{hotelRepository: repository}
}

func (service *HotelService) GetHotelById(id uuid.UUID) (*dto.HotelDto, error) {
	hotel, err := service.hotelRepository.Get(id)
	if err != nil {
		return nil, fmt.Errorf("GetHotelById: %w", err)
	}
	return dto.HotelDtoFromEntity(hotel), nil
}

func (service *HotelService) GetAllHotels() ([]*dto.HotelDto, error) {
	hotels, err := service.hotelRepository.GetAll()
	if err != nil {
		return nil, fmt.Errorf("GetAllHotels: %w", err)
	}
	hotelDtoList := make([]*dto.HotelDto, len(hotels))
	for i, hotel := range hotels {
		hotelDtoList[i] = dto.HotelDtoFromEntity(hotel)
	}
	return hotelDtoList, nil
}

func (service *HotelService) CreateHotel(hotelData dto.HotelModifyDto) (uuid.UUID, error) {
	hotelier, err := service.hotelierRepository.Get(hotelData.HotelierId)
	if err != nil {
		return uuid.Nil, fmt.Errorf("CreateHotel: %w", err)
	}

	hotel := model.Hotel{
		ID:          uuid.Nil,
		Name:        hotelData.Name,
		Description: hotelData.Description,
		Location:    hotelData.Location,
		HotelierID:  hotelData.HotelierId,
		Hotelier:    *hotelier,
	}

	id, err := service.hotelRepository.Put(&hotel)
	if err != nil {
		return uuid.Nil, fmt.Errorf("CreateHotel: %w", err)
	}

	return id, nil
}

func (service *HotelService) UpdateHotel(id uuid.UUID, hotelData dto.HotelModifyDto) error {
	hotel, err := service.hotelRepository.Get(id)
	if err != nil {
		return fmt.Errorf("UpdateHotel: %w", err)
	}
	hotelier, err := service.hotelierRepository.Get(hotelData.HotelierId)
	if err != nil {
		return fmt.Errorf("CreateHotel: %w", err)
	}

	hotel.Name = hotelData.Name
	hotel.Description = hotelData.Description
	hotel.Location = hotelData.Location
	hotel.HotelierID = hotelData.HotelierId
	hotel.Hotelier = *hotelier

	err = service.hotelRepository.Update(hotel)
	if err != nil {
		return fmt.Errorf("UpdateHotel: %w", err)
	}
	return nil
}

func (service *HotelService) DeleteHotel(id uuid.UUID) error {
	return service.hotelRepository.Delete(id)
}
