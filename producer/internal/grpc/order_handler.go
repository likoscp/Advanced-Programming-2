package grpc

// import (
// 	"context"
// 	"log"

// 	// "github.com/nats-io/nats.go"
// 	"github.com/likoscp/final/producer/pkg/nats"
// 	pb "github.com/likoscp/final/proto/order"
// )

// type OrderHandler struct {
// 	pb.UnimplementedOrderServiceServer
// 	publisher *nats.Publisher
// }

// func NewOrderHandler(publisher *nats.Publisher) *OrderHandler {
// 	return &OrderHandler{
// 		publisher: publisher,
// 	}
// }

// func (h *OrderHandler) CreateOrder(ctx context.Context, req *pb.CreateOrderRequest) (*pb.CreateOrderResponse, error) {
// 	if err := h.publisher.PublishOrderCreated(req); err != nil {
// 		log.Printf("Failed to publish order: %v", err)
// 		return nil, err
// 	}

// 	return &pb.CreateOrderResponse{}, nil
// }
