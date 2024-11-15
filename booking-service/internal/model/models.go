package model

import (
	"time"

	"github.com/google/uuid"
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
	InTime  time.Time
	OutTime time.Time
	Cost    uint64
}

type Reservation struct {
	Id              uuid.UUID `gorm:"primaryKey"`
	ReservationData `gorm:"embedded"`
}
