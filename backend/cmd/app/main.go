package main

import (
	"log/slog"

	_ "github.com/likoscp/Advanced-Programming-2/backend/cmd/app/docs"
	"github.com/likoscp/Advanced-Programming-2/backend/internal/config"
	"github.com/likoscp/Advanced-Programming-2/backend/internal/server"
)

//	@title			Swagger Store API
//	@version		1.0
//	@description	This is a sample server Comics-viewer.
//	@host			localhost:8080
func main() {
	cnf := config.NewConfig()

	s, err := server.NewServer(cnf)
	if err != nil {
		slog.Error("lol", "error", err)
	}

	if err := s.Run(); err != nil {
		slog.Error("error to start server", "error", err)
	}
}
