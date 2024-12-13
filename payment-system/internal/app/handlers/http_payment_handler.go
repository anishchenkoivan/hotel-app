package handlers

import (
	"context"
	"encoding/json"
	"fmt"
	pb "github.com/anishchenkoivan/hotel-app/api/generated/bookingservice"
	"github.com/anishchenkoivan/hotel-app/payment-system/config"
	"github.com/anishchenkoivan/hotel-app/payment-system/internal/model"
	"github.com/gorilla/mux"
	"google.golang.org/grpc"
	"log"
	"net/http"
	"sync"
)

type HTTPPaymentHandler struct {
	mu                  *sync.Mutex
	bookingEntityByHash map[string]model.BookingEntity
	config              config.Config
}

func NewHTTPPaymentHandler(config config.Config, mu *sync.Mutex, bookingEntityByHash map[string]model.BookingEntity) *HTTPPaymentHandler {
	return &HTTPPaymentHandler{config: config, mu: mu, bookingEntityByHash: bookingEntityByHash}
}

func (s *HTTPPaymentHandler) PaymentHandle(w http.ResponseWriter, r *http.Request) {
	token, exists := mux.Vars(r)["token"]
	if !exists {
		w.WriteHeader(http.StatusBadRequest)
		log.Println("Payment system handler error: no token in url")
		return
	}
	s.mu.Lock()
	bookingEntity, exists := s.bookingEntityByHash[token]
	s.mu.Unlock()
	if !exists {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	if r.Method == http.MethodPost {
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

		client := pb.NewBookingServiceClient(conn)
		req := &pb.ConfirmPaymentRequest{BookingId: bookingEntity.BookingId, IsConfirmed: true}
		_, err = client.ConfirmPayment(context.Background(), req)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			log.Println("Payment system paymentHandle error with grpc request:", err)
		}
		s.mu.Lock()
		delete(s.bookingEntityByHash, token)
		s.mu.Unlock()
	} else if r.Method == http.MethodGet {
		data := map[string]float32{"price": bookingEntity.BookingPrice}
		_ = json.NewEncoder(w).Encode(data)
		//TODO info about payment
	}
}
