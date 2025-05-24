package config

import (
	"os"
)

type Config struct {
	Addr        string
	PGHost      string
	PGPort      string
	PGUser      string
	PGPassword  string
	PGDB        string
	Secret      string
	PostgresDSN string
}

func NewConfig() (*Config, error) {
	return &Config{
		Addr:        os.Getenv("ADDR"),
		PGHost:      os.Getenv("PGHOST"),
		PGPort:      os.Getenv("PGPORT"),
		PGUser:      os.Getenv("PGUSER"),
		PGPassword:  os.Getenv("PGPASSWORD"),
		PGDB:        os.Getenv("PGDATABASE"),
		Secret:      os.Getenv("SECRET"),
		PostgresDSN: os.Getenv("POSTGRES_DSN"),
	}, nil
}
