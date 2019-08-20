package model

import (
	"github.com/k-yomo/go_echo_api_boilerplate/pkg/clock"
	"github.com/k-yomo/go_echo_api_boilerplate/pkg/token_generator"
	"time"
)

var smsReconfirmationExpiresIn = time.Minute * 30

// NewSMSConfirmation returns initialized sms reconfirmation
func NewSMSReconfirmation(userId uint64, phoneNumber string) *SmsReconfirmation {
	return &SmsReconfirmation{
		UserID:      userId,
		PhoneNumber: phoneNumber,
		AuthCode:    token_generator.GenerateRandomNumStr(authCodeLength),
		CreatedAt:   clock.Now(),
		UpdatedAt:   clock.Now(),
	}
}

// IsExpired checks if the sms reconfirmation is expired (specified hours have passed)
func (sr *SmsReconfirmation) IsExpired() bool {
	return clock.Now().After(sr.UpdatedAt.Add(smsReconfirmationExpiresIn))
}
