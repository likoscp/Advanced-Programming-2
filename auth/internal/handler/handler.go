package handler

import (
	"github.com/likoscp/Advanced-Programming-2/auth/internal/config"
	"github.com/likoscp/Advanced-Programming-2/auth/internal/store"
)

type Handler struct {
	userHandler  *UserHandler
	adminHandler *AdminHandler
	config       *config.Config
	db           *store.MongoDB
}

func NewHandler(db *store.MongoDB, config *config.Config) *Handler {
	return &Handler{
		db:     db,
		config: config,
	}
}

func (h *Handler) UserHandler() *UserHandler {
	if h.userHandler == nil {
		h.userHandler = NewUserHandler(h.db, h.config)
	}
	return h.userHandler
}

func (h *Handler) AdminHandler() *AdminHandler {
	if h.adminHandler == nil {
		h.adminHandler = NewAdminHandler(h.db, h.config)
	}
	return h.adminHandler
}