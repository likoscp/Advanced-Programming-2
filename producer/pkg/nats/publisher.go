package nats

import (
	"encoding/json"
	"github.com/nats-io/nats.go"
	"log"
	// pb "github.com/likoscp/final/proto/producer"
)

type Publisher struct {
	conn *nats.Conn
}

func NewPublisher(natsURL string) (*Publisher, error) {
	nc, err := nats.Connect(natsURL)
	if err != nil {
		return nil, err
	}
	return &Publisher{conn: nc}, nil
}

func (p *Publisher) PublishOrderCreated(order interface{}) error {
	jsonData, err := json.Marshal(order)
	if err != nil {
		return err
	}

	if err := p.conn.Publish("order.created", jsonData); err != nil {
		return err
	}

	log.Println("Published order.created event")
	return nil
}

func (p *Publisher) Close() {
	p.conn.Close()
}
