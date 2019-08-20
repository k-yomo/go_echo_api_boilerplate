package config

import (
	"os"
)

// NewSendgridApiKey returns sendgrid api key
func NewSendgridApiKey() string {
	sendgridApiKey := os.Getenv("SENDGRID_API_KEY")
	if sendgridApiKey == "" {
		panic("load SENDGRID_API_KEY from env variable failed")
	}
	return sendgridApiKey
}
