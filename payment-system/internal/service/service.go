package service

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"github.com/anishchenkoivan/hotel-app/payment-system/config"
	"github.com/anishchenkoivan/hotel-app/payment-system/internal/model"
	"sync"
	"time"
)

type PaymentSystemService struct {
	mu                  *sync.Mutex
	bookingEntityByHash map[string]model.BookingEntity
	config              config.Config
	bookingService      BookingService
}

func NewPaymentSystemService(config config.Config, bookingService BookingService) *PaymentSystemService {
	return &PaymentSystemService{
		config:              config,
		mu:                  &sync.Mutex{},
		bookingEntityByHash: make(map[string]model.BookingEntity),
		bookingService:      bookingService}
}

func (s *PaymentSystemService) GetBookingEntity(token string) (model.BookingEntity, bool) {
	s.mu.Lock()
	bookingEntity, exists := s.bookingEntityByHash[token]
	s.mu.Unlock()
	return bookingEntity, exists
}

func (s *PaymentSystemService) DeleteBookingEntity(token string) (entity model.BookingEntity, exists bool) {
	s.mu.Lock()
	bookingEntity, exists := s.bookingEntityByHash[token]
	if exists {
		delete(s.bookingEntityByHash, token)
	}
	s.mu.Unlock()
	return bookingEntity, exists
}

func (s *PaymentSystemService) AddPayment(bookingId string, bookingCost float32) string {
	token := GenerateToken(bookingId)
	s.mu.Lock()
	_, exists := s.bookingEntityByHash[token]
	for exists {
		_, exists = s.bookingEntityByHash[token]
	}
	s.bookingEntityByHash[token] = model.BookingEntity{BookingId: bookingId, BookingPrice: bookingCost}
	s.mu.Unlock()
	go func() {
		time.Sleep(s.config.PaymentTimeout)
		s.mu.Lock()
		bookingEntity, exists := s.bookingEntityByHash[token]
		if !exists {
			s.mu.Unlock()
			return
		}
		delete(s.bookingEntityByHash, token)
		s.mu.Unlock()
		_ = s.SendWebhook(bookingEntity.BookingId, false)
	}()

	return fmt.Sprintf("http://%s:%s/payment-system/api/pay/%s", s.config.ServerHost, s.config.ServerPort, token)
}

func (s *PaymentSystemService) SendWebhook(bookingId string, status bool) error {
	return s.bookingService.SendWebhook(bookingId, status)
}

func GenerateToken(bookingId string) string {
	uniqueString := bookingId + time.Now().String()
	hasher := sha256.New()
	hasher.Write([]byte(uniqueString))
	token := hasher.Sum(nil)
	return hex.EncodeToString(token)
}
