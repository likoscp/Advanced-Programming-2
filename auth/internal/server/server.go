package server

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/likoscp/Advanced-Programming-2/auth/internal/config"
)

type Server struct {
	config *config.Config
	Router *mux.Router
}

func NewServer(config *config.Config) *Server {
	router := mux.NewRouter()

	s := Server{Router: router,config: config}

	return &s
}


func (s *Server) Run() error {

	// s.Router.HandleFunc("/")

	return http.ListenAndServe(s.config.Addr, s.Router)
}
