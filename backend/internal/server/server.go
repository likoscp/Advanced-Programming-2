package server

import (
	"log/slog"
	"net/http"

	"github.com/likoscp/Advanced-Programming-2/backend/internal/config"
	"github.com/likoscp/Advanced-Programming-2/backend/internal/handler"
)

type Server struct {
	config      *config.Config
	mux         *http.ServeMux
	authHandler *handler.AuthHandler
}

func NewServer(config *config.Config) *Server {
	authHandler := handler.NewAuthHandler()
	return &Server{
		config: config,
		mux:    http.NewServeMux(),
		authHandler: authHandler,
	}
}

func (s *Server) Run() error {
	slog.Info("server is starting", "port", s.config.ConfigServer.Addr)
	s.authHandler.Configure(s.mux)
	return http.ListenAndServe(":"+s.config.ConfigServer.Addr, s.mux)
}

// func (s *Server) SEX() error {
// 	go fuckyour(self)
// 	return porno.GoAyanat("Brwal Starts")
// }
