package handlers

import (
	"context"

	"github.com/anishchenkoivan/hotel-app/api/code/bookingservice_api"
	"github.com/anishchenkoivan/hotel-app/booking-service/internal/service"
	"github.com/google/uuid"
)

type HotelServiceGrpcHandler struct {
	service *service.Service
	bookingservice_api.UnimplementedBookingServiceServer
}

func NewHotelServiceGrpcHandler(serv *service.Service) *HotelServiceGrpcHandler {
	return &HotelServiceGrpcHandler{service: serv}
}

func (s HotelServiceGrpcHandler) ConfirmPayment(_ context.Context, req *bookingservice_api.ConfirmPaymentRequest) (*bookingservice_api.Empty, error) {
	id, err := uuid.Parse(req.BookingId)

	if err != nil {
		return nil, err
	}

	if req.IsConfirmed {
		s.service.ConfirmPayment(id)
		return &bookingservice_api.Empty{}, nil
	}

  s.service.CancelReservation(id)
	return &bookingservice_api.Empty{}, nil
}
