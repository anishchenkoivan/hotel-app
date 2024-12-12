package model

import (
	"github.com/google/uuid"
)

type Hotel struct {
	ID          uuid.UUID `gorm:"primary_key;type:uuid;default:uuid_generate_v4()"`
	Name        string
	Description string
	Location    string
	HotelierID  uuid.UUID `gorm:"type:uuid"`
	Hotelier    Hotelier  `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;foreignKey:HotelierID"`
}

type Hotelier struct {
	ID       uuid.UUID `gorm:"primary_key;type:uuid;default:uuid_generate_v4()"`
	Username string
}

type Room struct {
	ID          uuid.UUID `gorm:"primary_key;type:uuid;default:uuid_generate_v4()"`
	Name        string
	Description string
	HotelID     uuid.UUID `gorm:"type:uuid"`
	Hotel       Hotel     `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;foreignKey:HotelID"`
	PricePerDay int64
	Capacity    int
}
