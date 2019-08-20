package mock

// SMSMessenger represents mocked mailer
type SMSMessenger struct {
}

// SendSMSMessage mocks sending sms message
func (s *SMSMessenger) SendSMSMessage(toPhoneNumber string, body string) error {
	return nil
}
