package handler

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
	// "time"

	// "github.com/google/uuid"
	"github.com/likoscp/Advanced-Programming-2/forum/internal/service"
	"github.com/likoscp/Advanced-Programming-2/forum/models"

	"github.com/golang-jwt/jwt/v5"
)

type ForumHandler struct {
	service *service.ForumService
}

func (h *ForumHandler) AddReply(w http.ResponseWriter, r *http.Request) {
	panic("unimplemented")
}

func NewForumHandler(s *service.ForumService) *ForumHandler {
	return &ForumHandler{service: s}
}

func (h *ForumHandler) CreatePost(w http.ResponseWriter, r *http.Request) {
	author := getAuthorFromCookie(r)

	var req struct {
		Content  string `json:"content"`
		ParentID string `json:"parent_id"`
	}
	_ = json.NewDecoder(r.Body).Decode(&req)

	if req.ParentID == "" {
		thread := models.Thread{
			Author:  author,
			Content: req.Content,
		}
		h.service.CreateThread(r.Context(), thread)
	} 
	// else {
	// 	reply := models.Reply{
	// 		ID:        uuid.New().String(),
	// 		Author:    author,
	// 		Content:   req.Content,
	// 		CreatedAt: time.Now(),
	// 	}
	// 	h.service.AddReply(r.Context(), req.ParentID, reply)
	// }

	w.WriteHeader(http.StatusCreated)
}

func (h *ForumHandler) GetThread(w http.ResponseWriter, r *http.Request) {
	id := strings.TrimPrefix(r.URL.Path, "/forum/")
	thread, err := h.service.GetThread(r.Context(), id)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	json.NewEncoder(w).Encode(thread)
}

func (h *ForumHandler) UpdateThread(w http.ResponseWriter, r *http.Request) {
	id := strings.TrimPrefix(r.URL.Path, "/forum/")
	var req struct {
		Content string `json:"content"`
	}
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, "invalid request body", http.StatusBadRequest)
		return
	}

	err = h.service.UpdateThread(r.Context(), id, req.Content)
	if err != nil {
		http.Error(w, "thread not found", http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (h *ForumHandler) DeleteThread(w http.ResponseWriter, r *http.Request) {
	id := strings.TrimPrefix(r.URL.Path, "/forum/")
	err := h.service.DeleteThread(r.Context(), id)
	if err != nil {
		http.Error(w, "thread not found", http.StatusNotFound)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func getAuthorFromCookie(r *http.Request) string {
	cookie, err := r.Cookie("token")
	if err != nil {
		log.Println("No token in cookie: ", err)
		return "anonymous"
	}

	tokenStr := cookie.Value
	log.Println("Token from cookie: ", tokenStr)

	secret := os.Getenv("SECRET")
	token, err := jwt.Parse(tokenStr, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			log.Println("Unexpected signing method: ", t.Header["alg"])
			return nil, fmt.Errorf("unexpected signing method: %v", t.Header["alg"])
		}
		return []byte(secret), nil
	})

	if err != nil || !token.Valid {
		log.Println("Invalid token: ", err)
		return "anonymous"
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		log.Println("Invalid token claims: ", err)
		return "anonymous"
	}

	userID, ok := claims["_id"].(string)
	if !ok {
		log.Println("No _id in token claims: ", err)
		return "anonymous"
	}

	log.Println("Authenticated user with _id: ", userID)
	return userID
}
