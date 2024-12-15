package service

type BookingService interface {
	SendWebhook(bookingId string, status bool) error
}
