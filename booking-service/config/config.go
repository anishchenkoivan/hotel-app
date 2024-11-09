package config

import (
	"time"

	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	ServerHost         string        `envconfig:"SERVER_HOST"`
	ServerPort         string        `envconfig:"SERVER_PORT"`

	DbHost             string        `envconfig:"DB_HOST"`
	DbPort             string        `envconfig:"DB_PORT"`
	DbName             string        `envconfig:"DB_NAME"`
	DbUser             string        `envconfig:"DB_USER"`
	DbPassword         string        `envconfig:"DB_PASSWORD"`

	AppShutdownTimeout time.Duration `envconfig:"APP_SHUTDOWN_TIMEOUT"`
}

func NewConfig() (Config, error) {
	cfg := Config{}
	err := envconfig.Process("BOOKING_SERVICE", &cfg)
	return cfg, err
}
