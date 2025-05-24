package handler

import (
	"encoding/json"
	"log"
	"net/http"

	"email-service/internal/db"
	"email-service/internal/service"
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

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("User registered successfully and email sent"))
	log.Printf("Successfully sent email to %s", req.Email)
}
