package main

import (
	"log/slog"

	"github.com/likoscp/Advanced-Programming-2/backend/internal/config"
	"github.com/likoscp/Advanced-Programming-2/backend/internal/server"
)

func main() {
	cnf := config.NewConfig()

	s := server.NewServer(cnf)

	if err := s.Run(); err != nil {
		slog.Error("error to start server", "error", err)
	}
}
