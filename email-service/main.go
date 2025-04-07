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

func main() {
    // Load .env file
    if err := godotenv.Load(); err != nil {
        log.Fatalf("Error loading .env file: %v", err)
    }

    // Load configuration
    cfg, err := config.LoadConfig()
    if err != nil {
        log.Fatalf("Failed to load configuration: %v", err)
    }

    // Connect to MongoDB
    mongoClient, err := db.ConnectMongo(cfg.MongoURI, cfg.DBName)
    if err != nil {
        log.Fatalf("Failed to connect to MongoDB: %v", err)
    }
    defer mongoClient.Client.Disconnect(nil)

    // Initialize services
    emailService := service.NewEmailService(cfg.SMTPHost, cfg.SMTPPort, cfg.SMTPEmail, cfg.SMTPPassword)
    handler := handler.NewHandler(emailService, mongoClient, cfg.Collection)

    // Start HTTP server
    server := &http.Server{
        Addr:    cfg.Addr,
        Handler: handler.SetupRoutes(),
    }

    log.Printf("Starting server on %s", cfg.Addr)
    if err := server.ListenAndServe(); err != nil {
        log.Fatalf("Server failed: %v", err)
    }
}