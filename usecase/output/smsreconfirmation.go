package output

import (
	"github.com/k-yomo/go_echo_api_boilerplate/model"
	"github.com/ttacon/libphonenumber"
	"time"
)

// TempUserOutput represents response body for temporary user
type SMSReconfirmationOutput struct {
	ID          uint64 `json:"id" example:"1"`
	PhoneNumber string `json:"phoneNumber" example:"080-1111-2222"`
	CreatedAt   string `json:"createdAt" example:"2020-01-01T00:00:00+09:00"`
}

// NewSMSReconfirmationOutput returns initialized sms reconfirmation
func NewSMSReconfirmationOutput(s *model.SmsReconfirmation) *SMSReconfirmationOutput {
	phoneNumber, _ := libphonenumber.Parse(s.PhoneNumber, "")
	return &SMSReconfirmationOutput{
		ID:          s.ID,
		PhoneNumber: libphonenumber.Format(phoneNumber, libphonenumber.NATIONAL),
		CreatedAt:   s.CreatedAt.Format(time.RFC3339),
	}
}
