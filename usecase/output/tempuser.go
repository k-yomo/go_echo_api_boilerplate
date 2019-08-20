package output

import (
	"github.com/k-yomo/go_echo_api_boilerplate/model"
	"github.com/ttacon/libphonenumber"
	"time"
)

// TempUserOutput represents response body for temporary user
type TempUserOutput struct {
	ID          uint64 `json:"id" example:"1"`
	PhoneNumber string `json:"phoneNumber" example:"080-1111-2222"`
	AuthKey     string `json:"authKey" example:"o1w.qeTWAXAl1lcueHRH"`
	CreatedAt   string `json:"createdAt" example:"2020-01-01T00:00:00+09:00"`
}

// NewTempUserOutput returns initialized temporary user
func NewTempUserOutput(u *model.TempUser) *TempUserOutput {
	phoneNumber, _ := libphonenumber.Parse(u.PhoneNumber, "")
	return &TempUserOutput{
		ID:          u.ID,
		PhoneNumber: libphonenumber.Format(phoneNumber, libphonenumber.NATIONAL),
		AuthKey:     u.AuthKey,
		CreatedAt:   u.CreatedAt.Format(time.RFC3339),
	}
}
