package service

import "errors"

var (
	RepositoryError          error = errors.New("Repository error")
	ReservationAlreadyExists error = errors.New("Reservation already exists")
	HotelServiceError        error = errors.New("Hotel service Error")
	PayemntSystemError       error = errors.New("Payment system Error")
	InvalidReservation       error = errors.New("Invalid reservation")
)
