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

	DbHost     string
	DbPort     string
	DbName     string
	DbUser     string
	DbPassword string

	AppShutdownTimeout time.Duration
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
		ServerHost: os.Getenv("HOTEL_SERVICE_SERVER_HOST"),
		ServerPort: os.Getenv("HOTEL_SERVICE_SERVER_PORT"),

		ServerGrpcHost: os.Getenv("HOTEL_SERVICE_SERVER_GRPC_HOST"),
		ServerGrpcPort: os.Getenv("HOTEL_SERVICE_SERVER_GRPC_PORT"),

		DbHost:     os.Getenv("HOTEL_SERVICE_DB_HOST"),
		DbPort:     os.Getenv("HOTEL_SERVICE_DB_PORT"),
		DbName:     os.Getenv("HOTEL_SERVICE_DB_NAME"),
		DbUser:     os.Getenv("HOTEL_SERVICE_DB_USER"),
		DbPassword: os.Getenv("HOTEL_SERVICE_DB_PASSWORD"),

		AppShutdownTimeout: getTimeout(os.Getenv("HOTEL_SERVICE_APP_SHUTDOWN_TIMEOUT")),
	}
}
