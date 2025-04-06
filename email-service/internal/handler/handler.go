package handler

import (
    "encoding/json"
    "net/http"

    "email-service/internal/db"
    "email-service/internal/service"
)

type Request struct {
    Username string `json:"username"`
    Email    string `json:"email"`
    Password string `json:"password"`
}

type Handler struct {
    emailService *service.EmailService
    client       *db.MongoClient
    collection   string
}

func NewHandler(es *service.EmailService, client *db.MongoClient, collection string) *Handler {
    return &Handler{
        emailService: es,
        client:       client,
        collection:   collection,
    }
}

func (h *Handler) SetupRoutes() http.Handler {
    router := http.NewServeMux()

    router.HandleFunc("/register", h.Register)

    return router
}

func (h *Handler) Register(w http.ResponseWriter, r *http.Request) {
    var req Request
    if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
        http.Error(w, "Invalid request body", http.StatusBadRequest)
        return
    }

    // Save user to MongoDB (mocked here)
    user := map[string]interface{}{
        "username": req.Username,
        "email":    req.Email,
        "password": req.Password,
    }
    _, err := h.client.Client.Database("user").Collection(h.collection).InsertOne(r.Context(), user)
    if err != nil {
        http.Error(w, "Failed to register user", http.StatusInternalServerError)
        return
    }

    // Send email
    subject := "Successful Registration"
    body := `<h1>Welcome, ` + req.Username + `!</h1><p>Thank you for registering.</p><p>Your login details:</p><p>Email: ` + req.Email + `</p>`

    if err := h.emailService.SendEmail(req.Email, subject, body); err != nil {
        http.Error(w, "Failed to send email", http.StatusInternalServerError)
        return
    }

    w.WriteHeader(http.StatusOK)
    w.Write([]byte("User registered successfully"))
}