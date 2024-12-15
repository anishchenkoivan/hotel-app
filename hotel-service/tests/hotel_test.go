package main

import (
	"errors"
	"testing"

	"github.com/anishchenkoivan/hotel-app/hotel-service/internal/app/handlers/dto"
	"github.com/anishchenkoivan/hotel-app/hotel-service/internal/model"
	service "github.com/anishchenkoivan/hotel-app/hotel-service/internal/service/hotel"
	hotel_mocks "github.com/anishchenkoivan/hotel-app/hotel-service/tests/mocks/hotel"
	hotelier_mocks "github.com/anishchenkoivan/hotel-app/hotel-service/tests/mocks/hotelier"
	"github.com/google/uuid"
)

func TestGetHotelByIDNormal(t *testing.T) {
	id := uuid.New()

	expectedDto := dto.HotelDto{ID: id}
	expectedModel := model.Hotel{ID: id}

	hotelRepo := hotel_mocks.MockHotelRepository{
		GetReturnValue: &expectedModel,
		GetReturnError: nil,
	}
	hotelierRepo := hotelier_mocks.MockHotelierRepository{}

	serv := service.NewHotelService(&hotelRepo, &hotelierRepo)
	hotel, err := serv.GetHotelById(id)

	if err != nil {
		t.Fatalf("err is not nil. %v", err)
	}
	if *hotel != expectedDto {
		t.Fatalf("Hotels not match")
	}
}

func TestGetHotelByIDRepositoryError(t *testing.T) {
	expectedErr := errors.New("test error")
	id := uuid.New()

	hotelRepo := hotel_mocks.MockHotelRepository{
		GetReturnValue: nil,
		GetReturnError: expectedErr,
	}
	hotelierRepo := hotelier_mocks.MockHotelierRepository{}

	serv := service.NewHotelService(&hotelRepo, &hotelierRepo)
	hotel, err := serv.GetHotelById(id)

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

func TestGetAllHotelsNormal(t *testing.T) {
	id := []uuid.UUID{uuid.New(), uuid.New(), uuid.New()}

	expectedDto := []dto.HotelDto{{ID: id[0]}, {ID: id[1]}, {ID: id[2]}}
	expectedModel := []*model.Hotel{{ID: id[0]}, {ID: id[1]}, {ID: id[2]}}

	hotelRepo := hotel_mocks.MockHotelRepository{
		GetAllReturnValue: expectedModel,
		GetReturnError:    nil,
	}
	hotelierRepo := hotelier_mocks.MockHotelierRepository{}

	serv := service.NewHotelService(&hotelRepo, &hotelierRepo)

	hotels, err := serv.GetAllHotels()

	if err != nil {
		t.Fatalf("err is not nil. %v", err)
	}
	if hotels == nil {
		t.Fatalf("Returned value is nil")
	}
	for i := range hotels {
		if *hotels[i] != expectedDto[i] {
			t.Fatalf("Returned list of hotels does not match")
		}
	}
}

func TestGetAllHotelsRepositoryError(t *testing.T) {
	expectedErr := errors.New("test error")

	hotelRepo := hotel_mocks.MockHotelRepository{
		GetAllReturnValue: nil,
		GetAllReturnError: expectedErr,
	}
	hotelierRepo := hotelier_mocks.MockHotelierRepository{}
	serv := service.NewHotelService(&hotelRepo, &hotelierRepo)
	hotels, err := serv.GetAllHotels()

	if err == nil {
		t.Fatalf("Err is nil")
	}
	if !errors.Is(err, expectedErr) {
		t.Fatalf("!errors.Is(%v, %v)", err, expectedErr)
	}
	if hotels != nil {
		t.Fatalf("Hotels is not nil")
	}
}

func TestCreateHotelNormal(t *testing.T) {
	hotelId := uuid.New()
  hotelierId := uuid.New()

	hotelDto := dto.HotelModifyDto{
		Name: "test",
    Description: "test",
    Location: "test",
    HotelierId: hotelierId,
	}
  hotelierModel := model.Hotelier{
    ID: hotelierId,
    Username: "test",
  }
  hotelModel := model.Hotel{
    ID: uuid.Nil,
		Name: "test",
    Description: "test",
    Location: "test",
    HotelierID: hotelierId,
    Hotelier: hotelierModel,
  }

	hotelRepo := hotel_mocks.MockHotelRepository{
		PutReturnValue: hotelId,
		PutReturnError: nil,
	}
	hotelierRepo := hotelier_mocks.MockHotelierRepository{
    GetReturnValue: &hotelierModel,
    GetReturnError: nil,
  }

	serv := service.NewHotelService(&hotelRepo, &hotelierRepo)
  id, err := serv.CreateHotel(hotelDto)

  if err != nil {
		t.Fatalf("err is not nil. %v", err)
  }
  if id != hotelId {
    t.Fatalf("Id not match. Expected %s, got %s", id.String(), hotelId.String())
  }
  if hotelRepo.LastPut == nil {
		t.Fatalf("Nothing is put in repository")
  }
  if *hotelRepo.LastPut != hotelModel {
		t.Fatalf("Last put value does not match")
  }
}

func TestCreateHotelHotelRepositoryError(t *testing.T) {
  expectedErr := errors.New("test error")

  hotelierId := uuid.New()

	hotelDto := dto.HotelModifyDto{
		Name: "test",
    Description: "test",
    Location: "test",
    HotelierId: hotelierId,
	}
  hotelierModel := model.Hotelier{
    ID: hotelierId,
    Username: "test",
  }

	hotelRepo := hotel_mocks.MockHotelRepository{
		PutReturnError: expectedErr,
	}
	hotelierRepo := hotelier_mocks.MockHotelierRepository{
    GetReturnValue: &hotelierModel,
    GetReturnError: nil,
  }

	serv := service.NewHotelService(&hotelRepo, &hotelierRepo)
  _, err := serv.CreateHotel(hotelDto)

  if err == nil {
		t.Fatalf("err is nil")
  }
	if !errors.Is(err, expectedErr) {
		t.Fatalf("!errors.Is(%v, %v)", err, expectedErr)
	}
}

func TestCreateHotelHotelierRepositoryError(t *testing.T) {
  expectedErr := errors.New("test error")

  hotelierId := uuid.New()

	hotelDto := dto.HotelModifyDto{
		Name: "test",
    Description: "test",
    Location: "test",
    HotelierId: hotelierId,
	}

	hotelRepo := hotel_mocks.MockHotelRepository{}
	hotelierRepo := hotelier_mocks.MockHotelierRepository{
    GetReturnValue: nil,
    GetReturnError: expectedErr,
  }

	serv := service.NewHotelService(&hotelRepo, &hotelierRepo)
  _, err := serv.CreateHotel(hotelDto)

  if err == nil {
		t.Fatalf("err is nil")
  }
	if !errors.Is(err, expectedErr) {
		t.Fatalf("!errors.Is(%v, %v)", err, expectedErr)
	}
  if hotelRepo.LastPut != nil {
    t.Fatalf("Something is put")
  }
}
