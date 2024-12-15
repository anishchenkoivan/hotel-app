package handlers

import (
	"context"
	pb "github.com/anishchenkoivan/hotel-app/payment-system/api/api_v1pb"
	"github.com/anishchenkoivan/hotel-app/payment-system/internal/service"
)

type AddPaymentGrpcHandler struct {
	service *service.PaymentSystemService
	pb.UnimplementedPaymentSystemServer
}

func NewAddPaymentGrpcHandler(service *service.PaymentSystemService) *AddPaymentGrpcHandler {
	return &AddPaymentGrpcHandler{service: service}
}

func (s *AddPaymentGrpcHandler) AddPayment(_ context.Context, req *pb.AddPaymentRequest) (*pb.AddPaymentResponse, error) {
	urlForPay := s.service.AddPayment(req.BookingId, req.BookingCost)
	return &pb.AddPaymentResponse{UrlForPay: urlForPay}, nil
}
