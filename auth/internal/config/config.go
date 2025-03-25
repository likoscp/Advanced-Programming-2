package config

import "github.com/joho/godotenv"

type Config struct {
	Addr string
}

func NewConfig() (*Config, error) {
	c, err := godotenv.Read(".env")
	if err != nil {
		return nil, err
	}

	return &Config{Addr: c["ADDR"]}, nil
}
