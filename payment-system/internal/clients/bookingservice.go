package clients

import (
	"context"
	"fmt"
	"google.golang.org/grpc/credentials/insecure"
	"log"

	bookingpb "github.com/anishchenkoivan/hotel-app/api/code/bookingservice_api"
	"github.com/anishchenkoivan/hotel-app/payment-system/config"
	"google.golang.org/grpc"
)

type BookingService struct {
	config config.Config
}

func NewBookingService(config config.Config) *BookingService {
	return &BookingService{config}
}

func (s *BookingService) SendWebhook(bookingId string, status bool) error {
	conn, err := grpc.NewClient(
		fmt.Sprintf("%s:%s", s.config.BookingServiceHost, s.config.BookingServicePort),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		log.Fatalf("Payment system can't connect to grpc client: %v", err)
		return err
	}
	defer func(conn *grpc.ClientConn) {
		err := conn.Close()
		if err != nil {
			log.Fatalf("Payment system can't close connection: %v", err)
		}
	}(conn)

	client := bookingpb.NewBookingServiceClient(conn)
	req := &bookingpb.ConfirmPaymentRequest{BookingId: bookingId, IsConfirmed: status}
	_, err = client.ConfirmPayment(context.Background(), req)
	if err != nil {
		log.Println("Payment system grpc webhook error with grpc request:", err)
		return err
	}
	return nil
}
