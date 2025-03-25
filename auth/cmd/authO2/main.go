package main

import (
	"flag"

	"github.com/likoscp/Advanced-Programming-2/auth/internal/config"
	"github.com/likoscp/Advanced-Programming-2/auth/internal/server"
	log "github.com/sirupsen/logrus"
	prefixed "github.com/x-cray/logrus-prefixed-formatter"
)

var (
	port int
)

func init() {
	flag.IntVar(&port, "port", 8081, "number of port")
}

func main() {
	c := config.Config{Addr: ":8081"}
	s := server.NewServer(&c)

	logger := log.New()

	logger.SetFormatter(&prefixed.TextFormatter{
		DisableColors:   false,
		TimestampFormat: "2006-01-02 15:04:05",
		FullTimestamp:   true,
		ForceFormatting: true,
	})

	log.Error("faeioifajo")

	if err := s.Run(); err != nil {

	}
}
