package sms

import (
	"github.com/pkg/errors"
	"github.com/sfreiberg/gotwilio"
)

// SMSMessenger represents sms messenger
type SMSMessenger interface {
	SendSMSMessage(phoneNumber string, body string) error
}

type twilioMessenger struct {
	phoneNumber string
	*gotwilio.Twilio
}

// NewMailer returns initialized sms messenger
func NewSMSMessenger(twilioAccountSid, twilioAuthToken, phoneNumber string) *twilioMessenger {
	return &twilioMessenger{phoneNumber, gotwilio.NewTwilioClient(twilioAccountSid, twilioAuthToken)}
}

// SendSMSMessage sends sms message
func (s *twilioMessenger) SendSMSMessage(toPhoneNumber string, body string) error {
	_, exception, err := s.SendSMS(s.phoneNumber, toPhoneNumber, body, "", "")
	if err != nil {
		return errors.Wrap(err, "send sms message failed")
	}
	if exception != nil {
		return errors.Errorf("send sms message with twilio failed: status=%d, code=%d, message=%s", exception.Status, exception.Code, exception.Message)
	}
	return nil
}
