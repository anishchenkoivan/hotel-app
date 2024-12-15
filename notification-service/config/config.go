package config

import (
	"github.com/joho/godotenv"
	"log"
	"os"
	"time"
)

type Config struct {
	Topic      string
	BrokerHost string
	BrokerPort string
	GroupId    string

	AppShutdownTimeout time.Duration
}

func getTimeout(timeoutString string) time.Duration {
	timeout, err := time.ParseDuration(timeoutString)
	if err != nil {
		log.Fatal("Error parsing timeout string", err)
	}
	return timeout
}

func NewConfig() *Config {
	err := godotenv.Load("config/config.env")
	if err != nil {
		log.Fatal("Error loading .env file", err)
		return nil
	}

	return &Config{
		Topic:      os.Getenv("KAFKA_CONSUME_TOPIC"),
		BrokerHost: os.Getenv("KAFKA_BROKER_HOST"),
		BrokerPort: os.Getenv("KAFKA_BROKER_PORT"),
		GroupId:    os.Getenv("KAFKA_GROUP_ID"),

		AppShutdownTimeout: getTimeout(os.Getenv("NOTIFICATION_SERVICE_SHUTDOWN_TIMEOUT")),
	}
}
