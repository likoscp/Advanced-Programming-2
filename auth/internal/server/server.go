package server

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/likoscp/Advanced-Programming-2/auth/internal/config"
	"github.com/likoscp/Advanced-Programming-2/auth/internal/handler"
	"github.com/likoscp/Advanced-Programming-2/auth/internal/store"
)

type Server struct {
	userHandler *handler.UserHandler
	config      *config.Config
	Router      *mux.Router
}

func NewServer(config *config.Config) *Server {
	router := mux.NewRouter()
	storage, err := store.NewMongoDB(config)

	if err != nil {
		log.Fatalf("cannot connect to db: %v", err)
	}

	userHanlder := handler.NewUserHandler(storage)
	s := Server{Router: router, config: config, userHandler: userHanlder}

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

	s.Router.HandleFunc("/user/register", s.userHandler.RegisterUser()).Methods("POST")
}
