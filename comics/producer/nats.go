package producer

import (
	"encoding/json"
	"github.com/nats-io/nats.go"
	"log"
)

// ChapterUpdatedEvent represents the structure of the chapter.updated event
type ChapterUpdatedEvent struct {
	ChapterID string  `json:"chapter_id"`
	ComicID   string  `json:"comic_id"`
	Title     string  `json:"title"`
	Number    float64 `json:"number"`
}

// Publisher handles NATS publishing
type Publisher struct {
	Conn *nats.Conn
}

// NewPublisher creates a new NATS publisher
func NewPublisher(url string) (*Publisher, error) {
	conn, err := nats.Connect(url)
	if err != nil {
		return nil, err
	}
	return &Publisher{Conn: conn}, nil
}

// PublishChapterUpdated publishes a chapter updated event
func (p *Publisher) PublishChapterUpdated(event ChapterUpdatedEvent) error {
	data, err := json.Marshal(event)
	if err != nil {
		return err
	}
	err = p.Conn.Publish("chapter.updated", data)
	if err != nil {
		return err
	}
	log.Printf("Published chapter.updated event for chapter ID: %s", event.ChapterID)
	return nil
}

// Close closes the NATS connection
func (p *Publisher) Close() {
	p.Conn.Close()
}
