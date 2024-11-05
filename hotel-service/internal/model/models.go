package model

import (
	"github.com/google/uuid"
)

type Hotel struct {
	ID          uuid.UUID
	Name        string
	Description string
	Hotelier    Hotelier
	Location    string
}

type Hotelier struct {
	Id       uuid.UUID
	Username string
}

type Room struct {
	Id          uuid.UUID
	IsAvailable bool
	Name        string
	Description string
	Hotel       Hotel
	PricePerDay float64
	Capacity    int
}
