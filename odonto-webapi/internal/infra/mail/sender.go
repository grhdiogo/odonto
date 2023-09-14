package mail

import (
	"bytes"
	"fmt"
	"net/smtp"
)

// --------------------------------------------------------------
// Configuration
// --------------------------------------------------------------

// SMTPConfig smtp configuration
type SMTPConfig struct {
	Username string
	Password string
	Host     string
	Port     int
}

// --------------------------------------------------------------
// Mail Sender
// --------------------------------------------------------------

// Sender mail sender
type Sender interface {
	SendEmail(from, to, subject, message string) error
}

// --------------------------------------------------------------
// Default Mail Sender
// --------------------------------------------------------------

type senderImpl struct {
	config SMTPConfig
}

func (s *senderImpl) bodyMessage(from, receiver, subject, message string) string {
	body := &bytes.Buffer{}
	// write mail header
	fmt.Fprintf(body, "From: NoReply <%s>\r\n", from)
	fmt.Fprintf(body, "To: Recebedor <%s>\r\n", receiver)
	fmt.Fprintf(body, "Subject: %s\r\n\r\n", subject)
	// write email body
	fmt.Fprint(body, message)
	// result
	return body.String()
}

func (s *senderImpl) SendEmail(from, to, subject, message string) error {
	// create body message
	msg := s.bodyMessage(from, to, subject, message)
	// convert in bytes
	body := []byte(msg)
	auth := smtp.PlainAuth("", s.config.Username, s.config.Password, s.config.Host)
	// send mail
	err := smtp.SendMail(fmt.Sprintf("%s:%d", s.config.Host, s.config.Port), auth, from, []string{to}, body)
	// handling the errors
	if err != nil {
		return err
	}
	// success
	return nil
}

func NewSender(config SMTPConfig) Sender {
	return &senderImpl{
		config: config,
	}
}

