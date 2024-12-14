package handlers

import (
	"context"
	pb "github.com/anishchenkoivan/hotel-app/api/code/hotelservice_api"
	service "github.com/anishchenkoivan/hotel-app/hotel-service/internal/service/room"
	"github.com/google/uuid"
)

type RoomGrpcHandler struct {
	service *service.RoomService
	pb.UnimplementedHotelServiceServer
}

func NewRoomGrpcHandler(service *service.RoomService) *RoomGrpcHandler {
	return &RoomGrpcHandler{service: service}
}

func (s *RoomGrpcHandler) GetRoom(_ context.Context, req *pb.GetRoomRequest) (*pb.GetRoomResponse, error) {
	id, err := uuid.Parse(req.Id)
	if err != nil {
		return nil, err
	}
	roomDto, err := s.service.GetRoomById(id)
	if err != nil {
		return nil, err
	}
	return &pb.GetRoomResponse{Price: roomDto.PricePerDay}, nil
}
