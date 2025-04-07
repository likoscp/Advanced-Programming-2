package main

import (
	"flag"
	"os"

	"github.com/likoscp/Advanced-Programming-2/subscription/internal/config"
	"github.com/likoscp/Advanced-Programming-2/subscription/internal/server"
	log "github.com/sirupsen/logrus"
	prefixed "github.com/x-cray/logrus-prefixed-formatter"
)

var (
	port int
)

func init() {
	flag.IntVar(&port, "port", 8085, "number port")
	logger := log.New()
	
	logger.SetFormatter(&prefixed.TextFormatter{
		DisableColors:   false,
		TimestampFormat: "2006-01-02 15:04:05",
		FullTimestamp:   true,
		ForceFormatting: true,
	})
}

func main() {

	c, err := config.NewConfig()
	if err != nil {
		log.Error(err)
		os.Exit(1)
	}

	s := server.NewServer(c)
	log.Info("server start on port:", c.Addr)

	if err := s.Run(); err != nil {
		log.Error(err)
		os.Exit(1)
	}
}
