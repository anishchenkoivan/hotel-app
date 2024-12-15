package clients

import (
	"context"
	"fmt"
	"log"

	"github.com/anishchenkoivan/hotel-app/api/code/paymentsystem_api"
	"github.com/anishchenkoivan/hotel-app/booking-service/config"
	"github.com/google/uuid"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type PaymentSystem struct {
	conf config.ServerConfig
}

func NewPayementSystem(conf config.ServerConfig) (*PaymentSystem, error) {
	return &PaymentSystem{conf}, nil
}

func (s PaymentSystem) AddPayment(id uuid.UUID, cost int64) (string, error) {
	conn, err := grpc.NewClient(fmt.Sprintf("%s:%s", s.conf.Host, s.conf.Port), grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		return "", fmt.Errorf("Can't connect: %v", err)
	}

	defer func(conn *grpc.ClientConn) {
		err := conn.Close()
		if err != nil {
			log.Fatal("Could not close connection: ", err)
		}
	}(conn)

	client := paymentsystem_api.NewPaymentSystemClient(conn)

	req := &paymentsystem_api.AddPaymentRequest{BookingId: id.String(), BookingCost: float32(cost)}
	resp, err := client.AddPayment(context.Background(), req)

	if err != nil {
		return "", err
	}

	return resp.UrlForPay, err
}
