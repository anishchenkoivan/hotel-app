package service

import (
	"fmt"
	"github.com/anishchenkoivan/hotel-app/hotel-service/internal/app/handlers/dto"
	"github.com/anishchenkoivan/hotel-app/hotel-service/internal/apperrors"
	"github.com/anishchenkoivan/hotel-app/hotel-service/internal/model"
	hotelrepository "github.com/anishchenkoivan/hotel-app/hotel-service/internal/service/hotel"
	hotelierrepository "github.com/anishchenkoivan/hotel-app/hotel-service/internal/service/hotelier"
	"github.com/google/uuid"
)

type RoomService struct {
	roomRepository     RoomRepository
	hotelRepository    hotelrepository.HotelRepository
	hotelierRepository hotelierrepository.HotelierRepository
}

func NewRoomService(roomRepository RoomRepository, hotelRepository hotelrepository.HotelRepository, hotelierRepository hotelierrepository.HotelierRepository) *RoomService {
	return &RoomService{roomRepository: roomRepository, hotelRepository: hotelRepository, hotelierRepository: hotelierRepository}
}

func (service *RoomService) GetRoomById(id uuid.UUID) (*dto.RoomDto, error) {
	room, err := service.roomRepository.Get(id)
	if err != nil {
		return nil, fmt.Errorf("GetRoomById %w", err)
	}
	return dto.RoomDtoFromEntity(room), nil
}

func (service *RoomService) GetAllRooms() ([]*dto.RoomDto, error) {
	rooms, err := service.roomRepository.GetAll()
	if err != nil {
		return nil, fmt.Errorf("GetAllRooms %w", err)
	}
	roomDtoList := make([]*dto.RoomDto, len(rooms))
	for i, room := range rooms {
		roomDtoList[i] = dto.RoomDtoFromEntity(room)
	}
	return roomDtoList, nil
}

func (service *RoomService) CreateRoom(roomData dto.RoomModifyDto) (uuid.UUID, error) {
	hotel, err := service.hotelRepository.Get(roomData.HotelId)
	if err != nil {
		return uuid.Nil, fmt.Errorf("CreateRoom %w", err)
	}

	hotelier, err := service.hotelierRepository.Get(hotel.HotelierID)
	if err != nil {
		return uuid.Nil, fmt.Errorf("CreateRoom %w", err)
	}

	if hotelier.TelegramID != roomData.TelegramId {
		return uuid.Nil, apperrors.NewAccessDeniedError("Telegram ID does not match")
	}

	room := model.Room{
		ID:          uuid.Nil,
		Name:        roomData.Name,
		Description: roomData.Description,
		HotelID:     hotel.ID,
		Hotel:       *hotel,
		PricePerDay: roomData.PricePerDay,
		Capacity:    roomData.Capacity,
	}

	id, err := service.roomRepository.Put(&room)
	if err != nil {
		return uuid.Nil, err
	}
	return id, nil
}

func (service *RoomService) UpdateRoom(id uuid.UUID, roomData dto.RoomModifyDto) error {
	room, err := service.roomRepository.Get(id)
	if err != nil {
		return fmt.Errorf("UpdateRoom: %w", err)
	}

	currentHotel, err := service.hotelRepository.Get(room.HotelID)
	if err != nil {
		return fmt.Errorf("UpdateRoom %w", err)
	}

	hotelier, err := service.hotelierRepository.Get(currentHotel.HotelierID)
	if err != nil {
		return fmt.Errorf("UpdateRoom %w", err)
	}

	if hotelier.TelegramID != roomData.TelegramId {
		return apperrors.NewAccessDeniedError("Telegram ID doesn't match")
	}

	hotel, err := service.hotelRepository.Get(roomData.HotelId)
	if err != nil {
		return fmt.Errorf("UpdateRoom %w", err)
	}

	room.Name = roomData.Name
	room.Description = roomData.Description
	room.HotelID = roomData.HotelId
	room.Hotel = *hotel
	room.PricePerDay = roomData.PricePerDay
	room.Capacity = roomData.Capacity

	err = service.roomRepository.Update(room)
	if err != nil {
		return fmt.Errorf("UpdateRoom: %w", err)
	}
	return nil
}

func (service *RoomService) DeleteRoom(id uuid.UUID) error {
	return service.roomRepository.Delete(id)
}
