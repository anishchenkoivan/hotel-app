package clients

import (
	"github.com/google/uuid"
	"time"
)

type Message struct {
	FirstName     string    `json:"firstName"`
	LastName      string    `json:"lastName"`
	TelegramId    string    `json:"telegramId"`
	ReservationId uuid.UUID `json:"reservationId"`
	RoomId        uuid.UUID `json:"roomId"`
	InTime        time.Time `json:"inTime"`
	OutTime       time.Time `json:"outTime"`
	Cost          int64     `json:"cost"`
}
