package config

import (
	"github.com/joho/godotenv"
	"log"
	"os"
	"time"
)

type Config struct {
	ServerHost string
	ServerPort string

	ServerGrpcHost string
	ServerGrpcPort string

	AppShutdownTimeout time.Duration
	PaymentTimeout     time.Duration

	BookingServiceHost string
	BookingServicePort string

	PaymentUrl string
}

func getTimeout(timeoutString string) time.Duration {
	timeout, err := time.ParseDuration(timeoutString)
	if err != nil {
		log.Fatal("Error parsing timeout string")
	}
	return timeout
}

func NewConfig() *Config {
	err := godotenv.Load("config/config.env")
	if err != nil {
		log.Fatal("Error loading .env file")
		return nil
	}

	return &Config{
		ServerHost: os.Getenv("PAYMENT_SYSTEM_HOST"),
		ServerPort: os.Getenv("PAYMENT_SYSTEM_PORT"),

		ServerGrpcHost: os.Getenv("PAYMENT_SYSTEM_GRPC_HOST"),
		ServerGrpcPort: os.Getenv("PAYMENT_SYSTEM_GRPC_PORT"),

		AppShutdownTimeout: getTimeout(os.Getenv("PAYMENT_SYSTEM_APP_SHUTDOWN_TIMEOUT")),
		PaymentTimeout:     getTimeout(os.Getenv("PAYMENT_SYSTEM_PAYMENT_TIMEOUT")),

		BookingServiceHost: os.Getenv("BOOKING_SERVICE_HOST"),
		BookingServicePort: os.Getenv("BOOKING_SERVICE_PORT"),

		PaymentUrl: os.Getenv("PAYMENT_URL"),
	}
}
