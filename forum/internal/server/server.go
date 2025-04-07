package server

import (
	"context"
	"log"
	"net/http"
	"time"

	"github.com/likoscp/Advanced-Programming-2/forum/internal/config"
	"github.com/likoscp/Advanced-Programming-2/forum/internal/handler"
	"github.com/likoscp/Advanced-Programming-2/forum/internal/repository"
	"github.com/likoscp/Advanced-Programming-2/forum/internal/service"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Server struct {
	cfg    *config.Config
	router *http.ServeMux
}

func NewServer(cfg *config.Config) *Server {
	return &Server{
		cfg:    cfg,
		router: http.NewServeMux(),
	}
}

func (s *Server) Run() error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(s.cfg.MongoUri))
	if err != nil {
		return err
	}

	database := client.Database(s.cfg.DBname)
	repo := repository.NewForumRepository(database)
	svc := service.NewForumService(repo)
	h := handler.NewForumHandler(svc)

s.router.HandleFunc("/forum/", func(w http.ResponseWriter, r *http.Request) {
    switch r.Method {
    case http.MethodPost:
        h.CreatePost(w, r) 
    default:
        http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
    }
})


	s.router.HandleFunc("/forum/{id}", func(w http.ResponseWriter, r *http.Request) {
    	switch r.Method {
    	case http.MethodGet:
    		h.GetThread(w, r) 
    	case http.MethodPut:
        	h.UpdateThread(w, r) 
		case http.MethodDelete:
			h.DeleteThread(w, r)
		default:
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})

	s.router.HandleFunc("/forum/replies", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPost {
			h.AddReply(w, r)
		} else {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})

	log.Printf("Server running on %s\n", s.cfg.Addr)
	return http.ListenAndServe(s.cfg.Addr, s.router)
}
