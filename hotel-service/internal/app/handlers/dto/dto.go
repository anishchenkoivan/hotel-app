package dto

import (
	"github.com/anishchenkoivan/hotel-app/hotel-service/internal/model"
	"github.com/google/uuid"
)

type HotelDto struct {
	ID          uuid.UUID `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Location    string    `json:"location"`
	HotelierId  uuid.UUID `json:"hotelierId"`
}

type HotelModifyDto struct {
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Location    string    `json:"location"`
	HotelierId  uuid.UUID `json:"hotelierId"`
}

type RoomDto struct {
	ID          uuid.UUID `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	HotelId     uuid.UUID `json:"hotelId"`
	PricePerDay int64     `json:"pricePerDay"`
	Capacity    int       `json:"capacity"`
}

type RoomModifyDto struct {
	Name        string    `json:"name"`
	Description string    `json:"description"`
	HotelId     uuid.UUID `json:"hotelId"`
	PricePerDay int64     `json:"pricePerDay"`
	Capacity    int       `json:"capacity"`
}

type HotelierDto struct {
	ID       uuid.UUID `json:"id"`
	Username string    `json:"username"`
}

type HotelierModifyDto struct {
	Username string `json:"name"`
}

func HotelDtoFromEntity(hotel *model.Hotel) *HotelDto {
	return &HotelDto{
		ID:          hotel.ID,
		Name:        hotel.Name,
		Description: hotel.Description,
		Location:    hotel.Location,
		HotelierId:  hotel.HotelierID,
	}
}

func HotelierDtoFromEntity(hotelier *model.Hotelier) *HotelierDto {
	return &HotelierDto{
		ID:       hotelier.ID,
		Username: hotelier.Username,
	}
}

func RoomDtoFromEntity(room *model.Room) *RoomDto {
	return &RoomDto{
		ID:          room.ID,
		Name:        room.Name,
		Description: room.Description,
		HotelId:     room.HotelID,
		PricePerDay: room.PricePerDay,
		Capacity:    room.Capacity,
	}
}
