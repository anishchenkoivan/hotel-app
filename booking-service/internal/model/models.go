package model

import (
	"time"

	"github.com/Rhymond/go-money"
	"github.com/google/uuid"
)

type Client struct {
	Name    string
	Surname string
	Phone   string
	Email   string
}

type ReservationData struct {
	Client  Client
	RoomId  uuid.UUID
	InTime  time.Time
	OutTime time.Time
	Cost    money.Money
}

type Reservation struct {
	Id uuid.UUID
	ReservationData
}
