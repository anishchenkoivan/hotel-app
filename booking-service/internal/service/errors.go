package service

import "errors"

var (
	RepositoryError          error = errors.New("Repository error")
	ReservationAlreadyExists error = errors.New("Reservation already exists")
	GrpcError                error = errors.New("Grpc Error")
	InvalidReservation       error = errors.New("Invalid reservation")
)
