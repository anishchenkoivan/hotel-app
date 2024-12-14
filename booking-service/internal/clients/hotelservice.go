package clients

import (
	"context"
	"fmt"
	"log"

	"github.com/anishchenkoivan/hotel-app/api/code/hotelservice_api"
	"github.com/anishchenkoivan/hotel-app/booking-service/config"
	"github.com/google/uuid"
	"google.golang.org/grpc"
)

type HotelService struct {
	service hotelservice_api.HotelServiceClient
	conf    config.ServerConfig
}

func NewHotelService(conf config.ServerConfig) (*HotelService, error) {
	conn, err := grpc.NewClient(fmt.Sprintf("%s:%s", conf.Host, conf.Port), grpc.WithInsecure())

	if err != nil {
		return nil, fmt.Errorf("Can't connect: %v", err)
	}

	defer func(conn *grpc.ClientConn) {
		err := conn.Close()
		if err != nil {
			log.Fatal("Could not close connection: ", err)
		}
	}(conn)

	return &HotelService{hotelservice_api.NewHotelServiceClient(conn), conf}, nil
}

func (s HotelService) GetPrice(id uuid.UUID) (int64, error) {
	conn, err := grpc.NewClient(fmt.Sprintf("%s:%s", s.conf.Host, s.conf.Port), grpc.WithInsecure())

	if err != nil {
		return 0, fmt.Errorf("Can't connect: %v", err)
	}

	defer func(conn *grpc.ClientConn) {
		err := conn.Close()
		if err != nil {
			log.Fatal("Could not close connection: ", err)
		}
	}(conn)

	client := hotelservice_api.NewHotelServiceClient(conn)

	req := &hotelservice_api.GetRoomRequest{Id: id.String()}
	resp, err := client.GetRoom(context.Background(), req)

	if err != nil {
		return 0, err
	}

	return resp.Price, err
}
