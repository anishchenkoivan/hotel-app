package handlers

import (
	"context"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	booking_pb "github.com/anishchenkoivan/hotel-app/api/generated/bookingservice"
	pb "github.com/anishchenkoivan/hotel-app/payment-system/api/api_v1pb"
	"github.com/anishchenkoivan/hotel-app/payment-system/config"
	"github.com/anishchenkoivan/hotel-app/payment-system/internal/model"
	"google.golang.org/grpc"
	"log"
	"sync"
	"time"
)

type AddPaymentGrpcHandler struct {
	mu                  *sync.Mutex
	bookingEntityByHash map[string]model.BookingEntity
	config              config.Config
	pb.UnimplementedPaymentSystemServer
}

func NewAddPaymentGrpcHandler(config config.Config, mu *sync.Mutex, bookingEntityByHash map[string]model.BookingEntity) *AddPaymentGrpcHandler {
	return &AddPaymentGrpcHandler{config: config, mu: mu, bookingEntityByHash: bookingEntityByHash}
}

func (s *AddPaymentGrpcHandler) AddPayment(_ context.Context, req *pb.AddPaymentRequest) (*pb.AddPaymentResponse, error) {
	token := GenerateToken(req.BookingId)
	s.mu.Lock()
	bookingEntity, exists := s.bookingEntityByHash[token]
	for exists {
		_, exists = s.bookingEntityByHash[token]
	}
	s.bookingEntityByHash[token] = model.BookingEntity{BookingId: req.BookingId, BookingPrice: req.BookingCost}
	s.mu.Unlock()
	go func() {
		time.Sleep(s.config.PaymentTimeout)
		s.mu.Lock()
		_, exists := s.bookingEntityByHash[token]
		if exists {
			conn, err := grpc.NewClient(fmt.Sprintf("%s:%s", s.config.BookingServiceHost, s.config.BookingServicePort))
			if err != nil {
				log.Fatalf("did not connect: %v", err)
			}
			defer func(conn *grpc.ClientConn) {
				err := conn.Close()
				if err != nil {
					log.Fatalf("could not close connection: %v", err)
				}
			}(conn)

			client := booking_pb.NewBookingServiceClient(conn)
			req := &booking_pb.ConfirmPaymentRequest{BookingId: bookingEntity.BookingId, IsConfirmed: false}
			_, err = client.ConfirmPayment(context.Background(), req)
			if err != nil {
				log.Println("Payment system grpc webhook error with grpc request:", err)
			}
			s.mu.Lock()
			delete(s.bookingEntityByHash, token)
			s.mu.Unlock()
		}
	}()

	return &pb.AddPaymentResponse{UrlForPay: fmt.Sprintf("http://%s:%s/payment-system/api/pay/%s", s.config.ServerHost, s.config.ServerPort, token)}, nil

}

func GenerateToken(bookingId string) string {
	uniqueString := bookingId + time.Now().String()
	hasher := sha256.New()
	hasher.Write([]byte(uniqueString))
	token := hasher.Sum(nil)
	return hex.EncodeToString(token)
}
