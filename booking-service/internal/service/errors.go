package service

type ErrType int

const (
  NoError                  ErrType = 0
	RepositoryError          ErrType = 1
	ReservationAlreadyExists ErrType = 2
  GrpcError                ErrType = 3
  BadReservation           ErrType = 4
)
