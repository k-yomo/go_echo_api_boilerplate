package input

import (
	"github.com/k-yomo/go_echo_boilerplate/pkg/clock"
)

// TempSignUpInput represents parameters for TempSignUp usecase
type TempSignUpInput struct {
	PhoneNumber string `json:"phoneNumber" validate:"required,phoneNumber" example:"08012345678 | validation: phone_number_format"`
	Region      string `json:"region" validate:"required,phoneNumberRegion" example:"JP | validation: region_code(https://github.com/ttacon/libphonenumber/blob/master/countrycodetoregionmap.go)"`
}

// ConfirmTempUserInput represents parameters for ConfirmTempUser usecase
type ConfirmTempUserInput struct {
	TempUserID uint64 `json:"tempUserId" validate:"required,numeric" example:"142 | validation: numeric"`
	AuthCode   string `json:"authCode" validate:"required" example:"123456"`
	AuthKey    string `json:"authKey" validate:"required" example:"o1w.qeTWAXAl1lcueHRH"`
}

// SignUpInput represents parameters for SignUp usecase
type SignUpInput struct {
	Email       string             `json:"email" validate:"required,email" example:"test@example.com | validation: email_format"`
	Password    string             `json:"password" validate:"required,min=6,max=100" example:"password | validation: min=6, max=100"`
	FirstName   string             `json:"firstName" validate:"required,min=1" example:"Taro | validation: min=1"`
	LastName    string             `json:"lastName" validate:"required,min=1" example:"Tanaka | validation: min=1"`
	Gender      string             `json:"gender" validate:"omitempty,oneof=unknown male female" example:"unknown | male | female"`
	DateOfBirth clock.NullableDate `json:"dateOfBirth" validate:"omitempty,date" example:"1995-07-05 | validation: iso8601date"`
}

// SignInInput represents parameters for SignIn usecase
type SignInInput struct {
	PhoneNumber string `json:"phoneNumber" validate:"required,phoneNumber" example:"08012345678 | validation: phone_number_format"`
	Region      string `json:"region" validate:"required,phoneNumberRegion" example:"JP | validation: region_code(https://github.com/ttacon/libphonenumber/blob/master/countrycodetoregionmap.go)"`
	Password    string `json:"password" validate:"required,min=6,max=100" example:"password | validation: min=6, max=100"`
}

// UpdateUnconfirmedPhoneNumberInput represents parameters for UpdatePhoneNumber usecase
type UpdateUnconfirmedPhoneNumberInput struct {
	PhoneNumber string `json:"phoneNumber" validate:"required,phoneNumber" example:"08012345678 | validation: phone_number_format"`
	Region      string `json:"region" validate:"required,phoneNumberRegion" example:"JP | validation: region_code(https://github.com/ttacon/libphonenumber/blob/master/countrycodetoregionmap.go)"`
}

// ConfirmPhoneNumberInput represents parameters for ConfirmPhoneNumber usecase
type ConfirmPhoneNumberInput struct {
	AuthCode string `json:"authCode" validate:"required" example:"123456"`
}
