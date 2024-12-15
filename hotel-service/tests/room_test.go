package main

import (
	"errors"
	"testing"

	"github.com/anishchenkoivan/hotel-app/hotel-service/internal/app/handlers/dto"
	"github.com/anishchenkoivan/hotel-app/hotel-service/internal/model"
	service "github.com/anishchenkoivan/hotel-app/hotel-service/internal/service/room"
	hotel_mocks "github.com/anishchenkoivan/hotel-app/hotel-service/tests/mocks/hotel"
	room_mocks "github.com/anishchenkoivan/hotel-app/hotel-service/tests/mocks/room"
	"github.com/google/uuid"
)

func TestGetRoomByIDNormal(t *testing.T) {
	id := uuid.New()

	expectedDto := dto.RoomDto{ID: id}
  expectedModel := model.Room{ID: id}

	roomRepo := room_mocks.MockRoomRepository{
    GetReturnValue: &expectedModel,
  }
  hotelRepo := hotel_mocks.MockHotelRepository{}

  serv := service.NewRoomService(&roomRepo, &hotelRepo)
	hotel, err := serv.GetRoomById(id)

	if err != nil {
		t.Fatalf("err is not nil. %v", err)
	}
	if *hotel != expectedDto {
		t.Fatalf("Rooms not match")
	}
}

func TestGetRoomByIDRepositoryError(t *testing.T) {
	expectedErr := errors.New("test error")
	id := uuid.New()

	hotelierRepo := room_mocks.MockRoomRepository{
    GetReturnError: expectedErr,
  }
  hotelRepo := hotel_mocks.MockHotelRepository{}

	serv := service.NewRoomService(&hotelierRepo, &hotelRepo)
	hotel, err := serv.GetRoomById(id)

	if err == nil {
		t.Fatalf("err is nil")
	}
	if !errors.Is(err, expectedErr) {
		t.Fatalf("!errors.Is(%v, %v)", err, expectedErr)
	}
	if hotel != nil {
		t.Fatalf("Room is not nil")
	}
}

func TestGetAllRoomsNormal(t *testing.T) {
	id := []uuid.UUID{uuid.New(), uuid.New(), uuid.New()}

	expectedDto := []dto.RoomDto{{ID: id[0]}, {ID: id[1]}, {ID: id[2]}}
	expectedModel := []*model.Room{{ID: id[0]}, {ID: id[1]}, {ID: id[2]}}

	hotelRepo := room_mocks.MockRoomRepository{
		GetAllReturnValue: expectedModel,
		GetReturnError:    nil,
	}
	hotelierRepo := hotel_mocks.MockHotelRepository{}

	serv := service.NewRoomService(&hotelRepo, &hotelierRepo)

	hotels, err := serv.GetAllRooms()

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

func TestGetAllRoomsRepositoryError(t *testing.T) {
	expectedErr := errors.New("test error")

	hotelRepo := room_mocks.MockRoomRepository{
		GetAllReturnValue: nil,
		GetAllReturnError: expectedErr,
	}
	hotelierRepo := hotel_mocks.MockHotelRepository{}
	serv := service.NewRoomService(&hotelRepo, &hotelierRepo)
	hotels, err := serv.GetAllRooms()

	if err == nil {
		t.Fatalf("Err is nil")
	}
	if !errors.Is(err, expectedErr) {
		t.Fatalf("!errors.Is(%v, %v)", err, expectedErr)
	}
	if hotels != nil {
		t.Fatalf("Rooms is not nil")
	}
}
