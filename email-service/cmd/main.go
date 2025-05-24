package main

import (
	"log"
	"net/http"

	"github.com/joho/godotenv"

	"email-service/internal/config"
	"email-service/internal/db"
	"email-service/internal/handler"
	"email-service/internal/service"
)

// main starts the email microservice
func main() {
	// Load .env file
	if err := godotenv.Load("./.env"); err != nil {
		log.Fatalf("Failed to load .env file: %v", err)
	}

	// Load configuration
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

	// Initialize PostgreSQL client
	postgresClient, err := db.ConnectPostgres(cfg)
	if err != nil {
		log.Fatalf("Failed to connect to PostgreSQL: %v", err)
	}
	defer postgresClient.DB.Close()

	// Initialize email service
	emailService := service.NewEmailService(cfg.SMTPHost, cfg.SMTPPort, cfg.SMTPEmail, cfg.SMTPPassword)

	// Setup HTTP routes
	router := handler.SetupRoutes(emailService, postgresClient)

	// Start the server
	log.Printf("Starting email microservice on %s", cfg.Addr)
	if err := http.ListenAndServe(cfg.Addr, router); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
