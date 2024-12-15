package main

import (
	"errors"
	"testing"
	"time"

	"github.com/anishchenkoivan/hotel-app/booking-service/internal/model"
	"github.com/anishchenkoivan/hotel-app/booking-service/internal/service"
	"github.com/anishchenkoivan/hotel-app/booking-service/tests/mocks"
	"github.com/google/uuid"
)

func TestAddReservation(t *testing.T) {
	expectedId := uuid.New()

	hotel := mocks.MockHotelService{GetPriceReturnValue: 10, GetPriceReturnError: nil}
  payment := mocks.MockPaymentSystem{}
	repo := mocks.MockRepository{
    IsAvailableReturnValue: true,
    IsAvailableReturnError: nil,
    PutReturnValue: expectedId,
    PutReturnError: nil,
  }

	data := model.Reservation{
    InTime: time.Unix(0, 0),
    OutTime: time.Unix(3, 0),
		Cost:   4,
		IsPaid: false,
	}

	serv := service.NewService(&repo, &hotel, &payment)

	id, _, err := serv.AddReservation(data)

  if err != nil {
    t.Fatalf("AddReservation returns not nil err = %s", err)
  }
  if repo.LastPutted == nil {
    t.Fatal("Data is not put in the repo")
  }
  if repo.LastPutted.Cost != hotel.GetPriceReturnValue {
    t.Fatalf("Putted cost does not match given from client. Expected %d, got %d", hotel.GetPriceReturnValue, repo.LastPutted.Cost)
  }
  if id != expectedId {
    t.Fatalf("Inserted id not match. Expected %s, got %s", expectedId, id)
  }
}

func TestGrpcError(t *testing.T) {
  expectedErr := errors.New("test error")

	hotel := mocks.MockHotelService{GetPriceReturnValue: 10, GetPriceReturnError: expectedErr}
  payment := mocks.MockPaymentSystem{}
	repo := mocks.MockRepository{
    IsAvailableReturnValue: true,
    IsAvailableReturnError: nil,
    PutReturnValue: uuid.New(),
    PutReturnError: nil,
  }

	data := model.Reservation{
    InTime: time.Unix(0, 0),
    OutTime: time.Unix(3, 0),
  }

	serv := service.NewService(&repo, &hotel, &payment)

	_, _, err := serv.AddReservation(data)

  if repo.LastPutted != nil {
    t.Errorf("Something is put in the repo when unable to get cost")
  }
  if err == nil {
    t.Errorf("Returned err is nil")
  }
  if !errors.Is(err, expectedErr) {
    t.Errorf("!errors.Is(%v, %v)", err, expectedErr)
  }
}

func TestAvailabilityCheckError(t *testing.T) {
  expectedErr := errors.New("test error")

	hotel := mocks.MockHotelService{GetPriceReturnValue: 10, GetPriceReturnError: nil}
  payment := mocks.MockPaymentSystem{}
	repo := mocks.MockRepository{
    IsAvailableReturnValue: true,
    IsAvailableReturnError: expectedErr,
    PutReturnValue: uuid.New(),
    PutReturnError: nil,
  }

	data := model.Reservation{
    InTime: time.Unix(0, 0),
    OutTime: time.Unix(3, 0),
  }

	serv := service.NewService(&repo, &hotel, &payment)

	_, _, err := serv.AddReservation(data)

  if repo.LastPutted != nil {
    t.Errorf("Something is putted in the repo when unable to get cost")
  }
  if err == nil {
    t.Errorf("Returned err is nil")
  }
  if !errors.Is(err, expectedErr) {
    t.Errorf("!errors.Is(%v, %v)", err, expectedErr)
  }
}

func TestPutError(t *testing.T) {
  expectedErr := errors.New("test error")

	hotel := mocks.MockHotelService{GetPriceReturnValue: 10, GetPriceReturnError: nil}
  payment := mocks.MockPaymentSystem{}
	repo := mocks.MockRepository{
    IsAvailableReturnValue: true,
    IsAvailableReturnError: nil,
    PutReturnValue: uuid.New(),
    PutReturnError: expectedErr,
  }

	data := model.Reservation{
    InTime: time.Unix(0, 0),
    OutTime: time.Unix(3, 0),
  }

	serv := service.NewService(&repo, &hotel, &payment)

	_, _, err := serv.AddReservation(data)

  if err == nil {
    t.Errorf("Returned err is nil")
  }
  if !errors.Is(err, expectedErr) {
    t.Errorf("!errors.Is(%v, %v)", err, expectedErr)
  }
}

func TestIsNotAvailable(t *testing.T) {
	hotel := mocks.MockHotelService{GetPriceReturnValue: 10, GetPriceReturnError: nil}
  payment := mocks.MockPaymentSystem{}
	repo := mocks.MockRepository{
    IsAvailableReturnValue: false,
    IsAvailableReturnError: nil,
    PutReturnValue: uuid.New(),
    PutReturnError: nil,
  }

	data := model.Reservation{
    InTime: time.Unix(0, 0),
    OutTime: time.Unix(3, 0),
  }

	serv := service.NewService(&repo, &hotel, &payment)

	_, _, err := serv.AddReservation(data)

  if err == nil {
    t.Errorf("Returned err is nil")
  }
  if !errors.Is(err, service.ReservationAlreadyExists) {
    t.Errorf("!errors.Is(%v, %v)", err, service.ReservationAlreadyExists)
  }
}

func TestInvalidReservation(t *testing.T) {
	hotel := mocks.MockHotelService{GetPriceReturnValue: 10, GetPriceReturnError: nil}
  payment := mocks.MockPaymentSystem{}
	repo := mocks.MockRepository{
    IsAvailableReturnValue: true,
    IsAvailableReturnError: nil,
    PutReturnValue: uuid.New(),
    PutReturnError: nil,
  }

	data := model.Reservation{
    InTime: time.Unix(3, 0),
    OutTime: time.Unix(0, 0),
  }

	serv := service.NewService(&repo, &hotel, &payment)

	_, _, err := serv.AddReservation(data)

  if err == nil {
    t.Errorf("Returned err is nil")
  }
  if !errors.Is(err, service.InvalidReservation) {
    t.Errorf("!errors.Is(%v, %v)", err, service.InvalidReservation)
  }
}
