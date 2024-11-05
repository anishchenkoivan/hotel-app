package model

import (
	"time"

	"github.com/Rhymond/go-money"
	"github.com/google/uuid"
)

type Client struct {
	name    string
	surname string
	phone   string
	email   string
}

type Reservation struct {
	id       uuid.UUID
	client   Client
	room_id  uuid.UUID
	in_time  time.Time
	out_time time.Time
	cost     money.Money
}
