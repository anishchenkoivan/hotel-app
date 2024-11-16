package service

import "errors"

type ServiceErrorType int

const (
	RepositoryError          ServiceErrorType = 0
	ReservationAlreadyExists ServiceErrorType = 1
)

type ServiceError struct {
	error
	ErrType ServiceErrorType
}

func NewReservationAlreadyExistsError() *ServiceError {
  msg := "Reservation already exists"
  err := errors.New(msg)
  return &ServiceError{error: err, ErrType: ReservationAlreadyExists}
}
