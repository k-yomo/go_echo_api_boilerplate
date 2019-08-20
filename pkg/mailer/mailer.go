package mailer

import (
	"github.com/k-yomo/go_echo_api_boilerplate/config"
	"github.com/pkg/errors"
	"github.com/sendgrid/sendgrid-go/helpers/mail"
)

var (
	fromAddress = "info@go-echo-api-boilerplate.com"
	fromName    = "THIS_SHOULD_BE_REPLACED"
)

// Mailer represents email sender
type Mailer interface {
	SendEmail(m *Mail) error
}

// Mail represents email
type Mail struct {
	To               string
	Subject          string
	PlainTextContent string
	HtmlContent      string
}

// EmailAddress represents email address
type EmailAddress struct {
	Name    string
	Address string
}

type sendgridMailer struct{}

// NewMailer returns initialized mailer
func NewMailer() *sendgridMailer {
	return &sendgridMailer{}
}

// SendEmail sends email
func (s *sendgridMailer) SendEmail(m *Mail) error {
	from := mail.NewEmail(fromName, fromAddress)
	to := mail.NewEmail(m.To, m.To)
	message := mail.NewSingleEmail(from, m.Subject, to, m.PlainTextContent, m.HtmlContent)
	client := config.NewSendgridClient()
	_, err := client.Send(message)
	if err != nil {
		return errors.Wrap(err, "send email failed")
	}
	return nil
}
