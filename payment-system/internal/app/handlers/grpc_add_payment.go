package handlers

import (
	"context"
	pb "github.com/anishchenkoivan/hotel-app/payment-system/api/api_v1pb"
	"log"
	"net/http"
	"strings"
	"time"
)

type AddPaymentGrpcHandler struct {
	CallbackTimeout time.Duration
	pb.UnimplementedPaymentSystemServer
}

func NewAddPaymentGrpcHandler(callbackTimeout time.Duration) *AddPaymentGrpcHandler {
	return &AddPaymentGrpcHandler{CallbackTimeout: callbackTimeout}
}

func (s *AddPaymentGrpcHandler) AddPayment(_ context.Context, req *pb.AddPaymentRequest) (*pb.AddPaymentResponse, error) {
	go func() {
		time.Sleep(s.CallbackTimeout)

		_, err := http.Post(req.GetCallbackUrl(), "application/json", strings.NewReader("todo")) //TODO
		if err != nil {
			log.Printf("Failed to send callback: %v", err)
		}
	}()

	return &pb.AddPaymentResponse{IsSuccessful: true}, nil

}
