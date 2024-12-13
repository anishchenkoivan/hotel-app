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
	Host     string `envconfig:"HOST"`
	Port     string `envconfig:"PORT"`
	Name     string `envconfig:"NAME"`
	User     string `envconfig:"USER"`
	Password string `envconfig:"PASSWORD"`
}

type AppConfig struct {
	ShutdownTimeout time.Duration `envconfig:"APP_SHUTDOWN_TIMEOUT"`
}

type Config struct {
	App    AppConfig
	Server ServerConfig
	Db     DbConfig
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

	return cfg, err
}
