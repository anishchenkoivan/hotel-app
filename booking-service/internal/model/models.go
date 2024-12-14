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
	ID          uuid.UUID `gorm:"primary_key;type:uuid;default:uuid_generate_v4()"`
	Reservation `gorm:"embedded"`
}
