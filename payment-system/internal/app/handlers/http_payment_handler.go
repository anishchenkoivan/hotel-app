package handlers

import (
	"encoding/json"
	"github.com/anishchenkoivan/hotel-app/payment-system/internal/service"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

type HTTPPaymentHandler struct {
	service *service.PaymentSystemService
}

func NewHTTPPaymentHandler(service *service.PaymentSystemService) *HTTPPaymentHandler {
	return &HTTPPaymentHandler{service: service}
}

func (s *HTTPPaymentHandler) PaymentHandle(w http.ResponseWriter, r *http.Request) {
	token, exists := mux.Vars(r)["token"]
	if !exists {
		w.WriteHeader(http.StatusBadRequest)
		log.Println("Payment system handler error: no token in url")
		return
	}
	if r.Method == http.MethodGet {
		bookingEntity, exists := s.service.GetBookingEntity(token)
		if !exists {
			w.WriteHeader(http.StatusNotFound)
			return
		}
		data := map[string]float32{"price": bookingEntity.BookingPrice}
		_ = json.NewEncoder(w).Encode(data)
		//TODO info about payment
	} else if r.Method == http.MethodPost {
		bookingEntity, exists := s.service.DeleteBookingEntity(token)
		if !exists {
			w.WriteHeader(http.StatusNotFound)
			return
		}
		err := s.service.SendWebhook(bookingEntity.BookingId, true)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
		}
	}
}
