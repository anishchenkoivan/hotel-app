package main

import (
	"errors"
	"testing"

	"github.com/anishchenkoivan/hotel-app/hotel-service/internal/app/handlers/dto"
	"github.com/anishchenkoivan/hotel-app/hotel-service/internal/model"
	service "github.com/anishchenkoivan/hotel-app/hotel-service/internal/service/hotelier"
	hotelier_mocks "github.com/anishchenkoivan/hotel-app/hotel-service/tests/mocks/hotelier"
	"github.com/google/uuid"
)

func TestGetHotelierByIDNormal(t *testing.T) {
	id := uuid.New()

	expectedDto := dto.HotelierDto{ID: id}
  expectedModel := model.Hotelier{ID: id}

	hotelierRepo := hotelier_mocks.MockHotelierRepository{
    GetReturnValue: &expectedModel,
  }

  serv := service.NewHotelierService(&hotelierRepo)
	hotel, err := serv.GetHotelierById(id)

	if err != nil {
		t.Fatalf("err is not nil. %v", err)
	}
	if *hotel != expectedDto {
		t.Fatalf("Hotels not match")
	}
}

func TestGetHotelierByIDRepositoryError(t *testing.T) {
	expectedErr := errors.New("test error")
	id := uuid.New()

	hotelierRepo := hotelier_mocks.MockHotelierRepository{
    GetReturnError: expectedErr,
  }

	serv := service.NewHotelierService(&hotelierRepo)
	hotel, err := serv.GetHotelierById(id)

	if err == nil {
		t.Fatalf("err is nil")
	}
	if !errors.Is(err, expectedErr) {
		t.Fatalf("!errors.Is(%v, %v)", err, expectedErr)
	}
	if hotel != nil {
		t.Fatalf("Hotel is not nil")
	}
}
