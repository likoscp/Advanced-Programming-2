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

	s.Configure()

	return http.ListenAndServe(s.config.Addr, s.Router)
}

func (s *Server) Configure() {
	s.Router.Handle("/ping", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(`PONG`))
	})).Methods("GET")
		

}