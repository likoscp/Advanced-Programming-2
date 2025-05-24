package subscriber

import (
	"encoding/json"
	"log"

	"email-service/internal/db"
	"email-service/internal/service"
	"github.com/nats-io/nats.go"
)

type ComicUploadedEvent struct {
	ComicID     string `json:"comic_id"`
	Title       string `json:"title"`
	Author      string `json:"author"`
	Description string `json:"description"`
}

type NATSClient struct {
	Conn         *nats.Conn
	DBClient     *db.PostgresClient
	EmailService *service.EmailService
}

func NewNATSClient(natsURL string, dbClient *db.PostgresClient, emailService *service.EmailService) (*NATSClient, error) {
	nc, err := nats.Connect(natsURL)
	if err != nil {
		return nil, err
	}
	return &NATSClient{
		Conn:         nc,
		DBClient:     dbClient,
		EmailService: emailService,
	}, nil
}

func (nc *NATSClient) SubscribeComicUploaded() error {
	_, err := nc.Conn.Subscribe("comic.uploaded", func(msg *nats.Msg) {
		var event ComicUploadedEvent
		if err := json.Unmarshal(msg.Data, &event); err != nil {
			log.Printf("Failed to unmarshal comic uploaded event: %v", err)
			return
		}

		// Fetch all users
		users, err := nc.DBClient.GetAllUsers()
		if err != nil {
			log.Printf("Failed to fetch users: %v", err)
			return
		}

		// Send email to each user
		subject := "New Comic Uploaded!"
		body := `<h1>New Comic Available!</h1><p>A new comic titled "` + event.Title + `" by ` + event.Author + ` has been uploaded.</p><p>Description: ` + event.Description + `</p>`
		for _, user := range users {
			email := user["email"].(string)
			if err := nc.EmailService.SendEmail(email, subject, body); err != nil {
				log.Printf("Failed to send email to %s: %v", email, err)
				continue
			}
			log.Printf("Successfully sent email to %s", email)
		}
	})
	if err != nil {
		return err
	}
	return nil
}
