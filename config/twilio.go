package config

import (
	"github.com/sfreiberg/gotwilio"
	"os"
)

var twilioAccountSid = os.Getenv("TWILIO_ACCOUNT_SID")
var twilioAuthToken = os.Getenv("TWILIO_AUTH_TOKEN")

func init() {
	if twilioAccountSid == "" {
		panic("load TWILIO_ACCOUNT_SID from env variable failed")
	}
	if twilioAuthToken == "" {
		panic("load TWILIO_AUTH_TOKEN from env variable failed")
	}
}

// NewTwilioClient returns initialized twilio client
func NewTwilioClient() *gotwilio.Twilio {
	return gotwilio.NewTwilioClient(twilioAccountSid, twilioAuthToken)
}
