package model

import (
	"github.com/google/uuid"
)

type HotelData struct {
	Name        string
	Description string
	Location    string
	HotelierID  uuid.UUID `gorm:"type:uuid"`
	Hotelier    Hotelier  `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;foreignKey:HotelierID"`
}

type Hotel struct {
	ID        uuid.UUID `gorm:"primary_key;type:uuid;default:uuid_generate_v4()"`
	HotelData `gorm:"embedded"`
}

type HotelierData struct {
	Username string
}

type Hotelier struct {
	ID           uuid.UUID `gorm:"primary_key;type:uuid;default:uuid_generate_v4()"`
	HotelierData `gorm:"embedded"`
}

type RoomData struct {
	IsAvailable bool
	Name        string
	Description string
	HotelID     uuid.UUID `gorm:"type:uuid"`
	Hotel       Hotel     `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;foreignKey:HotelID"`
	PricePerDay float64
	Capacity    int
}

type Room struct {
	ID       uuid.UUID `gorm:"primary_key;type:uuid;default:uuid_generate_v4()"`
	RoomData `gorm:"embedded"`
}
