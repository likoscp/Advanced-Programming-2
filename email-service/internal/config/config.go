package config

import (
	"log"
	"os"
	"strconv"
)

// Config holds configuration for the email microservice
type Config struct {
	Addr         string
	SMTPHost     string
	SMTPPort     int
	SMTPEmail    string
	SMTPPassword string
	DBHost       string
	DBAddr       string
	DBUser       string
	DBPassword   string
	DBName       string
}

// LoadConfig loads environment variables into Config
func LoadConfig() (*Config, error) {
	addr := os.Getenv("ADDR")
	smtpHost := os.Getenv("SMTP_HOST")
	smtpPortStr := os.Getenv("SMTP_PORT")
	smtpEmail := os.Getenv("SMTP_EMAIL")
	smtpPassword := os.Getenv("SMTP_PASSWORD")
	dbHost := os.Getenv("DB_HOST")
	dbAddr := os.Getenv("DB_ADDR")
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")

	if addr == "" || smtpHost == "" || smtpPortStr == "" || smtpEmail == "" || smtpPassword == "" ||
		dbHost == "" || dbAddr == "" || dbUser == "" || dbPassword == "" || dbName == "" {
		log.Fatal("Missing required environment variables")
	}

	smtpPort, err := strconv.Atoi(smtpPortStr)
	if err != nil {
		log.Fatalf("Invalid SMTP_PORT value: %v", err)
	}

	return &Config{
		Addr:         addr,
		SMTPHost:     smtpHost,
		SMTPPort:     smtpPort,
		SMTPEmail:    smtpEmail,
		SMTPPassword: smtpPassword,
		DBHost:       dbHost,
		DBAddr:       dbAddr,
		DBUser:       dbUser,
		DBPassword:   dbPassword,
		DBName:       dbName,
	}, nil
}
