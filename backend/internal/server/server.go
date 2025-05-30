package server

import (
	"log/slog"
	"net/http"

	_ "github.com/likoscp/Advanced-Programming-2/backend/cmd/app/docs"
	"github.com/likoscp/Advanced-Programming-2/backend/internal/config"
	"github.com/likoscp/Advanced-Programming-2/backend/internal/handler"
	httpSwagger "github.com/swaggo/http-swagger"
)

type Server struct {
	config      *config.Config
	mux         *http.ServeMux
	minioHandler *handler.S3Handler
	authHandler *handler.AuthHandler
}

func NewServer(config *config.Config) (*Server, error) {
	authHandler, err := handler.NewAuthHandler(config)
	if err != nil {
		return nil, err
	}

	mini, err := handler.NewS3Handler(config.ConfigS3)
	if err != nil {
		return nil, err
	}
	return &Server{
		config:      config,
		mux:         http.NewServeMux(),
		authHandler: authHandler,
		minioHandler: mini,
	}, nil
}

func (s *Server) Run() error {
	slog.Info("server is starting", "port", s.config.ConfigServer.Addr)
	s.authHandler.Configure(s.mux)

	s.mux.Handle("/swagger/", httpSwagger.WrapHandler)
	s.mux.HandleFunc("POST /comics/", s.minioHandler.UploadHandler)
	return http.ListenAndServe(":"+s.config.ConfigServer.Addr, s.mux)
}
