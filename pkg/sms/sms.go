package sms

import (
	"github.com/k-yomo/go_echo_api_boilerplate/config"
	"github.com/pkg/errors"
)

var (
	fromPhoneNumber = "+14156550110"
)

// SMSMessenger represents sms messenger
type SMSMessenger interface {
	SendSMS(phoneNumber string, body string) error
}

type twilioMessenger struct{}

// NewMailer returns initialized sms messenger
func NewSMSMessenger() *twilioMessenger {
	return &twilioMessenger{}
}

// SendSMS sends sms message
func (s *twilioMessenger) SendSMS(toPhoneNumber string, body string) error {
	client := config.NewTwilioClient()

	_, exception, err := client.SendSMS(fromPhoneNumber, toPhoneNumber, body, "", "")
	if err != nil {
		return errors.Wrap(err, "send sms message failed")
	}
	if exception != nil {
		return errors.Errorf("send sms message with twilio failed: status=%d, code=%d, message=%s", exception.Status, exception.Code, exception.Message)
	}
	return nil
}
