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

type Reservation struct {
	Client  Client `gorm:"embedded;embeddedPrefix:cliend_"`
	RoomId  uuid.UUID
	InTime  time.Time
	OutTime time.Time
	Cost    int64
	IsPaid  bool
}

type ReservationModel struct {
	Id          uuid.UUID `gorm:"primaryKey"`
	Reservation `gorm:"embedded"`
}
