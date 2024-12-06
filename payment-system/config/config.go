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
	CallbackTimeout    time.Duration
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
		CallbackTimeout:    getTimeout(os.Getenv("PAYMENT_SYSTEM_CALLBACK_TIMEOUT")),
	}
}
