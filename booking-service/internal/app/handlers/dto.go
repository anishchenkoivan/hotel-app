package handlers

import (
	"time"

	"github.com/anishchenkoivan/hotel-app/booking-service/internal/model"
	"github.com/google/uuid"
)

const timeLayout = "02.01.2006"

type CreateReservationDto struct {
	ClientName    string
	ClientSurname string
	ClientPhone   string
	ClientEmail   string
	RoomId        string
	InTime        string
	OutTime       string
}

type ReservationDto struct {
	ClientName    string
	ClientSurname string
	ClientPhone   string
	ClientEmail   string
	RoomId        string
	InTime        string
	OutTime       string
	Cost          int64
}

type ReservationsArrayDto struct {
	Reservations []ReservationDto
}

type ReservationIdDto struct {
	Id uuid.UUID
}

type RoomIdDto struct {
	Id uuid.UUID
}

func ReservationDtoFromModel(data model.ReservationModel) ReservationDto {
	return ReservationDto{
		ClientName:    data.Client.Name,
		ClientSurname: data.Client.Surname,
		ClientPhone:   data.Client.Phone,
		ClientEmail:   data.Client.Email,
		RoomId:        data.RoomId.String(),
		InTime:        data.InTime.Format("02.01.2006"),
		OutTime:       data.OutTime.Format("02.01.2006"),
		Cost:          data.Cost,
	}
}

func ReservationsArrayDtoFromModelsArray(data []model.ReservationModel) ReservationsArrayDto {
	reservs := make([]ReservationDto, len(data))

	for i := range data {
		reservs[i] = ReservationDto{
			ClientName:    data[i].Client.Name,
			ClientSurname: data[i].Client.Surname,
			ClientPhone:   data[i].Client.Phone,
			ClientEmail:   data[i].Client.Email,
			RoomId:        data[i].RoomId.String(),
			InTime:        data[i].InTime.Format("02.01.2006"),
			OutTime:       data[i].OutTime.Format("02.01.2006"),
			Cost:          data[i].Cost,
		}
	}

	return ReservationsArrayDto{Reservations: reservs}
}

func ReservationFromDto(dto CreateReservationDto) (model.Reservation, error) {
	uuid, err := uuid.Parse(dto.RoomId)

	if err != nil {
		return model.Reservation{}, err
	}

	inTime, err := time.Parse(timeLayout, dto.InTime)

	if err != nil {
		return model.Reservation{}, err
	}

	outTime, err := time.Parse(timeLayout, dto.OutTime)

	if err != nil {
		return model.Reservation{}, err
	}

	return model.Reservation{
		Client: model.Client{
			Name:    dto.ClientName,
			Surname: dto.ClientSurname,
			Phone:   dto.ClientPhone,
			Email:   dto.ClientEmail,
		},
		RoomId:  uuid,
		InTime:  inTime,
		OutTime: outTime,
	}, nil
}
