package server

import (
	"log/slog"
	"net/http"

	"github.com/likoscp/Advanced-Programming-2/backend/internal/config"
)

type Server struct {
	config *config.Config
	mux    *http.ServeMux
}

func NewServer(config *config.Config) *Server {

	return &Server{
		config: config,
		mux:    http.NewServeMux(),
	}
}

func (s *Server) Run() error {
	slog.Info("server is starting", "port", s.config.ConfigServer.Addr)
	return http.ListenAndServe(s.config.ConfigServer.Addr, s.mux)
}

// func (s *Server) SEX() error {
// 	go fuckyour(self)
// 	return porno.GoAyanat("Brwal Starts")
// }
