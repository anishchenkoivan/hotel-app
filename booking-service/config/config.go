package config

import (
	"time"

	"github.com/kelseyhightower/envconfig"
)

type ServerConfig struct {
	Host string `envconfig:"HOST"`
	Port string `envconfig:"PORT"`
}

type DbConfig struct {
	Host       string `envconfig:"HOST"`
	Port       string `envconfig:"PORT"`
	Name       string `envconfig:"NAME"`
	User       string `envconfig:"USER"`
	Password   string `envconfig:"PASSWORD"`
	Migrations string `envconfig:"MIGRATIONS"`
}

type KafkaConfig struct {
	Topic      string `envconfig:"PRODUCE_TOPIC"`
	BrokerHost string `envconfig:"BROKER_HOST"`
	BrokerPort string `envconfig:"BROKER_PORT"`
	GroupId    string `envconfig:"CONSUMER_GROUP_ID"`
}

type AppConfig struct {
	ShutdownTimeout time.Duration `envconfig:"SHUTDOWN_TIMEOUT"`
}

type Config struct {
	App    AppConfig
	Server ServerConfig
	Db     DbConfig

	BookingService ServerConfig
	HotelService   ServerConfig

	Kafka KafkaConfig
	PaymentSystem  ServerConfig
}

func NewConfig() (Config, error) {
	cfg := Config{}
	err := envconfig.Process("BOOKING_SERVICE_SERVER", &cfg.Server)

	if err != nil {
		return cfg, err
	}

	err = envconfig.Process("BOOKING_SERVICE_DB", &cfg.Db)

	if err != nil {
		return cfg, err
	}

	err = envconfig.Process("BOOKING_SERVICE_APP", &cfg.App)

	if err != nil {
		return cfg, err
	}

	err = envconfig.Process("HOTEL_SERVICE_GRPC", &cfg.HotelService)

  if err != nil {
		return cfg, err
	}

	err = envconfig.Process("PAYMENT_SYSTEM_GRPC", &cfg.PaymentSystem)

	if err != nil {
		return cfg, err
	}

	err = envconfig.Process("BOOKING_SERVICE_GRPC", &cfg.BookingService)

	if err != nil {
		return cfg, err
	}

	err = envconfig.Process("KAFKA", &cfg.Kafka)

	return cfg, err
}
