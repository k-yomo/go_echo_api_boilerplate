package model

import (
	"github.com/k-yomo/go_echo_api_boilerplate/pkg/clock"
	"github.com/k-yomo/go_echo_api_boilerplate/pkg/token_generator"
	"time"
)

var (
	authCodeLength    uint = 6
	authKeyLength     uint = 20
	tempUserExpiresIn      = time.Minute * 30
)

// NewTempUser returns initialized temp user
func NewTempUser(phoneNumber string) *TempUser {
	return &TempUser{
		PhoneNumber: phoneNumber,
		AuthCode:    token_generator.GenerateRandomNumStr(authCodeLength),
		AuthKey:     token_generator.GenerateRandomBase64Token(authKeyLength),
		CreatedAt:   clock.Now(),
		UpdatedAt:   clock.Now(),
	}
}

// IsExpired checks if the temp user confirmation is expired (specified hours have passed)
func (tu *TempUser) IsExpired() bool {
	return clock.Now().After(tu.UpdatedAt.Add(tempUserExpiresIn))
}

// ValidateAuthInfo validates given pair of authCode and authKey
func (tu *TempUser) ValidateAuthInfo(authCode string, authKey string) bool {
	return tu.AuthCode == authCode && tu.AuthKey == authKey
}
