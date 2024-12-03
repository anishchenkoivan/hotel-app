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
	Client           Client `gorm:"embedded;embeddedPrefix:cliend_"`
	RoomId           uuid.UUID
	Cost             uint64
	InTime  time.Time
	OutTime time.Time
}

type Reservation struct {
	Id              uuid.UUID `gorm:"primaryKey"`
	ReservationData `gorm:"embedded"`
}
