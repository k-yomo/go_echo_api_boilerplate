package config

import (
	"github.com/sendgrid/sendgrid-go"
	"os"
)

var sendgridApiKey = os.Getenv("SENDGRID_API_KEY")

func init() {
	if sendgridApiKey == "" {
		panic("load SENDGRID_API_KEY from env variable failed")
	}
}

// NewSendgridClient returns initialized sendgrid client
func NewSendgridClient() *sendgrid.Client {
	return sendgrid.NewSendClient(sendgridApiKey)
}
