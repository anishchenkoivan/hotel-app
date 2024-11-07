package model

import (
	"github.com/google/uuid"
)

type HotelData struct {
	Name        string
	Description string
	Hotelier    Hotelier
	Location    string
}

type Hotel struct {
	Id uuid.UUID
	HotelData
}

type HotelierData struct {
	Username string
}

type Hotelier struct {
	Id uuid.UUID
	HotelierData
}

type RoomData struct {
	Id          uuid.UUID
	IsAvailable bool
	Name        string
	Description string
	Hotel       Hotel
	PricePerDay float64
	Capacity    int
}

type Room struct {
	Id uuid.UUID
	RoomData
}
