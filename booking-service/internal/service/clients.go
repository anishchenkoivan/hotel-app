package service

import "github.com/google/uuid"

type HotelService interface {
	GetPrice(id uuid.UUID) (int64, error)
}
