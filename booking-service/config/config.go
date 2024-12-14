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

type AppConfig struct {
	ShutdownTimeout time.Duration `envconfig:"APP_SHUTDOWN_TIMEOUT"`
}

type Config struct {
	App    AppConfig
	Server ServerConfig
	Db     DbConfig

	HotelService ServerConfig
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

	err = envconfig.Process("HOTEL_SERVICE_SERVER", &cfg.HotelService)

	return cfg, err
}
