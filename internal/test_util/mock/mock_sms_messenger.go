package mock

// SMSMessenger represents mocked mailer
type SMSMessenger struct {
}

// SendSMS mocks sending sms message
func (s *SMSMessenger) SendSMS(toPhoneNumber string, body string) error {
	return nil
}
