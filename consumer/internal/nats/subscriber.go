package nats

// import (
// 	"context"
// 	"encoding/json"
// 	"log"

// 	"github.com/likoscp/finalAddProgramming/consumer/internal/mail"
// 	pb "github.com/likoscp/finalAddProgramming/finalProto/mail"
// 	"github.com/nats-io/nats.go"
// )

// type Subscriber struct {
// 	conn       *nats.Conn
// 	mailClient *mail.Client
// }

// func NewSubscriber(natsURL string, mailClient *mail.Client) (*Subscriber, error) {
// 	nc, err := nats.Connect(natsURL)
// 	if err != nil {
// 		return nil, err
// 	}
// 	return &Subscriber{conn: nc, mailClient: mailClient}, nil
// }

// func (s *Subscriber) SubscribeToComicUpdated() {
// 	s.conn.Subscribe("comic.updated", func(msg *nats.Msg) {
// 		var comic struct {
// 			Title        string `json:"title"`
// 			TranslatorID uint   `json:"translatorId"`
// 		}

// 		if err := json.Unmarshal(msg.Data, &comic); err != nil {
// 			log.Printf("Error unmarshaling comic: %v", err)
// 			return
// 		}

// 		log.Printf("Mail || Comics: %s, was updated by %d || was send", comic.Title, comic.TranslatorID)

// 	})
// }

// func (s *Subscriber) Close() {
// 	s.conn.Close()
// }
