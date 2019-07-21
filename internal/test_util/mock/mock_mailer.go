package mock

import "github.com/k-yomo/go_echo_boilerplate/pkg/mailer"

// Mailer represents mocked mailer
type Mailer struct {
}

// SendEmail mocks sending email
func (s *Mailer) SendEmail(m *mailer.Mail) error {
	return nil
}
