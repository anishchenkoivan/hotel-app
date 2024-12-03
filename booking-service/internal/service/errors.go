package service

import "errors"

type BookErrType int

const (
	RepositoryError          BookErrType = 0
	ReservationAlreadyExists BookErrType = 1
)

type BookErr struct {
	error
	ErrType BookErrType
}

func NewReservationAlreadyExistsError() *BookErr {
  msg := "Reservation already exists"
  err := errors.New(msg)
  return &BookErr{error: err, ErrType: ReservationAlreadyExists}
}
