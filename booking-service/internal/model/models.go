package model

import (
	"github.com/google/uuid"
	"gorm.io/datatypes"
)

type Client struct {
	Name    string
	Surname string
	Phone   string
	Email   string
}

type ReservationData struct {
	Client  Client `gorm:"embedded;embeddedPrefix:cliend_"`
	RoomId  uuid.UUID
	InTime  datatypes.Date
	OutTime datatypes.Date
	Cost    uint64 
}

type Reservation struct {
	Id              uuid.UUID `gorm:"primaryKey"`
	ReservationData `gorm:"embedded"`
}
