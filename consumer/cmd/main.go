package main

// import (
// 	"log"
// 	"os"
// 	"os/signal"
// 	"syscall"

// 	"github.com/likoscp/final/consumer/internal/nats"
// 	"github.com/likoscp/final/consumer/internal/mail"
// )

// func main() {
// 	mailClient, err := mail.NewMailClient("mail-service:8081")
// 	if err != nil {
// 		log.Fatalf("Failed to connect to mail-service: %v", err)
// 	}
// 	defer mailClient.Close()

// 	subscriber, err := nats.NewSubscriber("nats://nats:4222", mailClient)
// 	if err != nil {
// 		log.Fatalf("Failed to connect to NATS: %v", err)
// 	}
// 	defer subscriber.Close()

// 	subscriber.SubscribeToComicUpdated()
// 	log.Println("Consumer service subscribed to comic.updated events")

// 	sigChan := make(chan os.Signal, 1)
// 	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)
// 	<-sigChan
// 	log.Println("Shutting down consumer service...")
// }
