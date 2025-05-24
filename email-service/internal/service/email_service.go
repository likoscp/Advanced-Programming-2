package service

import (
	"gopkg.in/gomail.v2"
)

// EmailService handles email sending
type EmailService struct {
	host     string
	port     int
	email    string
	password string
}

// NewEmailService creates a new EmailService instance
func NewEmailService(host string, port int, email, password string) *EmailService {
	return &EmailService{
		host:     host,
		port:     port,
		email:    email,
		password: password,
	}
}

// SendEmail sends an email to the specified recipient
func (es *EmailService) SendEmail(to, subject, body string) error {
	m := gomail.NewMessage()
	m.SetHeader("From", es.email)
	m.SetHeader("To", to)
	m.SetHeader("Subject", subject)
	m.SetBody("text/html", body)

	d := gomail.NewDialer(es.host, es.port, es.email, es.password)
	if err := d.DialAndSend(m); err != nil {
		return err
	}

	return nil
}
