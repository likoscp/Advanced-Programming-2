package main

import (
	emailv1 "github.com/barcek2281/finalProto/gen/go/email"
	"log"
	"net"
	"net/http"

	"github.com/joho/godotenv"
	"google.golang.org/grpc"

	"email-service/internal/config"
	"email-service/internal/db"
	"email-service/internal/handler"
	"email-service/internal/service"
	"email-service/internal/subscriber"
)

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

	// Start HTTP server
	go func() {
		log.Printf("Starting email HTTP server on %s", cfg.Addr)
		if err := http.ListenAndServe(cfg.Addr, router); err != nil {
			log.Fatalf("Failed to start HTTP server: %v", err)
		}
	}()

	// Initialize NATS client and subscribe to events
	natsClient, err := subscriber.NewNATSClient(cfg.NATSURL, postgresClient, emailService)
	if err != nil {
		log.Fatalf("Failed to connect to NATS: %v", err)
	}
	defer natsClient.Conn.Close()

	if err := natsClient.SubscribeComicUploaded(); err != nil {
		log.Fatalf("Failed to subscribe to comic.uploaded: %v", err)
	}
	if err := natsClient.SubscribeChapterUpdated(); err != nil {
		log.Fatalf("Failed to subscribe to chapter.updated: %v", err)
	}

	// Start gRPC server
	grpcAddr := cfg.GRPCAddr
	lis, err := net.Listen("tcp", grpcAddr)
	if err != nil {
		log.Fatalf("Failed to listen on gRPC port %s: %v", grpcAddr, err)
	}
	grpcServer := grpc.NewServer()
	emailv1.RegisterEmailServiceServer(grpcServer, handler.NewGRPCServer(postgresClient, emailService))

	log.Printf("Starting email gRPC server on %s", grpcAddr)
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to serve gRPC: %v", err)
	}
}
