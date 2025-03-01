package config

import (
	"log"

	"github.com/caarlos0/env/v11"
)

type Config struct {
	Dev bool `env:"DEV" envDefault:"false"`

	Database Database
	Telegram Telegram
}

type Database struct {
	EnableSQLLog bool   `env:"ENABLE_SQL_LOG" envDefault:"false"`
	Host         string `env:"DB_HOST" envDefault:"localhost"`
	Port         string `env:"DB_PORT" envDefault:"5432"`
	User         string `env:"DB_USER" envDefault:"postgres"`
	Password     string `env:"DB_PASSWORD" envDefault:"postgres"`
	DB           string `env:"DB_NAME" envDefault:"postgres"`
	SSLMode      string `env:"DB_SSL_MODE" envDefault:"disable"`
}

type Telegram struct {
	Token string `env:"TELEGRAM_TOKEN" envDefault:""`
}

func New() *Config {
	cfg := &Config{}
	opts := env.Options{RequiredIfNoDef: true}
	if err := env.ParseWithOptions(cfg, opts); err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}
	return cfg
}
