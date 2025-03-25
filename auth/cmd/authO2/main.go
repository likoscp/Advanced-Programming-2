package main

import (
	"flag"
	"os"

	"github.com/likoscp/Advanced-Programming-2/auth/internal/config"
	"github.com/likoscp/Advanced-Programming-2/auth/internal/server"
	log "github.com/sirupsen/logrus"
	prefixed "github.com/x-cray/logrus-prefixed-formatter"
)

var (
	port int
)

func init() {
	flag.IntVar(&port, "port", 8081, "number port")
}

func main() {
	logger := log.New()

	logger.SetFormatter(&prefixed.TextFormatter{
		DisableColors:   false,
		TimestampFormat: "2006-01-02 15:04:05",
		FullTimestamp:   true,
		ForceFormatting: true,
	})

	c, err := config.NewConfig()
	if err != nil {
		log.Error(err)
		os.Exit(1)
	}

	s := server.NewServer(c)

	if err := s.Run(); err != nil {
		log.Error(err)
		os.Exit(1)
	}
}
