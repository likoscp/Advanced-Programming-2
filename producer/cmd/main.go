package main

// import (
// 	"log"
// 	"net"

// 	gr "google.golang.org/grpc"
// 	// "github.com/nats-io/nats.go"
// 	"github.com/likoscp/final/producer/internal/grpc"
// 	"github.com/likoscp/final/producer/pkg/nats"

// 	pb "github.com/likoscp/final/proto/order"
// )

// func main() {
// 	publisher, err := nats.NewPublisher("nats://nats:4222")
// 	if err != nil {
// 		log.Fatalf("Failed to connect to NATS: %v", err)
// 	}
// 	defer publisher.Close()

// 	lis, err := net.Listen("tcp", ":50051")
// 	if err != nil {
// 		log.Fatalf("failed to listen: %v", err)
// 	}

// 	s := gr.NewServer()
// 	pb.RegisterOrderServiceServer(s, grpc.NewOrderHandler(publisher))

// 	log.Println("Producer service started on :50051")
// 	if err := s.Serve(lis); err != nil {
// 		log.Fatalf("failed to serve: %v", err)
// 	}
// }
