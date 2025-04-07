package server

import (
    "github.com/likoscp/Advanced-Programming-2/subscription/internal/config"
    "net/http"
)

type Server struct {
    config *config.Config
}

func NewServer(c *config.Config) *Server {
    return &Server{config: c}
}

func (s *Server) Run() error {
    return http.ListenAndServe(s.config.Addr, nil)
}