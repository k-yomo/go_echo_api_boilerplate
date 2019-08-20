package config

import (
	"os"
)

var (
	fromPhoneNumber = "SET_YOUR_TWILIO_PHONE_NUMBER"
)

// NewTwilioTokens returns twilio account sid and auth token
func NewTwilioTokens() (twilioAccountSid, twilioAuthToken, phoneNumber string) {
	twilioAccountSid = os.Getenv("TWILIO_ACCOUNT_SID")
	twilioAuthToken = os.Getenv("TWILIO_AUTH_TOKEN")
	phoneNumber = fromPhoneNumber

	if twilioAccountSid == "" {
		panic("load TWILIO_ACCOUNT_SID from env variable failed")
	}
	if twilioAuthToken == "" {
		panic("load TWILIO_AUTH_TOKEN from env variable failed")
	}

	return
}
