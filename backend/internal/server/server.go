package server

import "github.com/likoscp/Advanced-Programming-2/backend/internal/config"

type Server struct {
	config *config.Config
}


func NewServer(config *config.Config) *Server {

	return &Server{
		config: config,
	}
}