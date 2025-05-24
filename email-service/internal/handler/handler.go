package handler

import (
	"context"
	"encoding/json"
	"log"
	"net/http"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"

	"email-service/internal/db"
	"email-service/internal/service"
	emailv1 "github.com/barcek2281/finalProto/gen/go/email"
)

// Request represents the incoming JSON payload
type Request struct {
	Username string `json:"username"`
	Email    string `json:"email"`
}

// Handler manages HTTP routes and email service
type Handler struct {
	emailService *service.EmailService
	client       *db.PostgresClient
}

// NewHandler creates a new Handler instance
func NewHandler(es *service.EmailService, client *db.PostgresClient) *Handler {
	return &Handler{
		emailService: es,
		client:       client,
	}
}

// SetupRoutes configures the HTTP router
func SetupRoutes(es *service.EmailService, client *db.PostgresClient) http.Handler {
	h := NewHandler(es, client)
	router := http.NewServeMux()
	router.HandleFunc("POST /mail/register", h.Register)
	router.Handle("/metrics", promhttp.Handler()) // Added metrics endpoint
	return router
}

// Register handles the registration email request
func (h *Handler) Register(w http.ResponseWriter, r *http.Request) {
	var req Request
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	// Verify user exists in PostgreSQL
	_, err := h.client.GetUserByEmail(req.Email)
	if err != nil {
		log.Printf("Failed to find user: %v", err)
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}

	// Send email
	subject := "Successful Registration"
	body := `<h1>Welcome, ` + req.Username + `!</h1><p>Thank you for registering with Comics.</p><p>Your login details:</p><p>Email: ` + req.Email + `</p>`

	if err := h.emailService.SendEmail(req.Email, subject, body); err != nil {
		log.Printf("Failed to send email: %v", err)
		http.Error(w, "Failed to send email", http.StatusInternalServerError)
		return
	}

	// Increment email sent counter
	emailSentCounter.Inc()

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("User registered successfully and email sent"))
	log.Printf("Successfully sent email to %s", req.Email)
}

type GRPCServer struct {
	emailv1.UnimplementedEmailServiceServer
	DBClient     *db.PostgresClient
	EmailService *service.EmailService
}

func NewGRPCServer(dbClient *db.PostgresClient, emailService *service.EmailService) *GRPCServer {
	return &GRPCServer{
		DBClient:     dbClient,
		EmailService: emailService,
	}
}

func (s *GRPCServer) NotifyComicUploaded(ctx context.Context, req *emailv1.NotifyComicUploadedRequest) (*emailv1.NotifyComicUploadedResponse, error) {
	// Fetch all users
	users, err := s.DBClient.GetAllUsers()
	if err != nil {
		log.Printf("Failed to fetch users: %v", err)
		return nil, err
	}

	// Send email to each user
	subject := "New Comic Uploaded!"
	body := `<h1>New Comic Available!</h1><p>A new comic titled "` + req.Title + `" by ` + req.Author + ` has been uploaded.</p><p>Description: ` + req.Description + `</p>`
	for _, user := range users {
		email := user["email"].(string)
		if err := s.EmailService.SendEmail(email, subject, body); err != nil {
			log.Printf("Failed to send email to %s: %v", email, err)
			continue
		}
		// Increment email sent counter
		emailSentCounter.Inc()
		log.Printf("Successfully sent email to %s", email)
	}

	return &emailv1.NotifyComicUploadedResponse{
		Message: "Emails sent successfully",
	}, nil
}

// Metrics initialization
var (
	emailSentCounter = prometheus.NewCounter(prometheus.CounterOpts{
		Name: "email_sent_total",
		Help: "Total number of emails sent.",
	})
	natsMessagesProcessed = prometheus.NewCounter(prometheus.CounterOpts{
		Name: "nats_messages_processed_total",
		Help: "Total number of NATS messages processed.",
	})
)

func init() {
	prometheus.MustRegister(emailSentCounter)
	prometheus.MustRegister(natsMessagesProcessed)
}
