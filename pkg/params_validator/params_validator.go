package params_validator

import (
	"github.com/k-yomo/go_echo_api_boilerplate/pkg/clock"
	"github.com/ttacon/libphonenumber"
	"gopkg.in/go-playground/validator.v9"
	"time"
)

type Validator struct {
	validator *validator.Validate
}

// NewValidator returns Validator
func NewValidator() *Validator {
	v := validator.New()
	_ = v.RegisterValidation("date", ValidateDate)
	_ = v.RegisterValidation("rfc3339DateTime", ValidateRfc3339Datetime)
	_ = v.RegisterValidation("phoneNumber", ValidatePhoneNumber)
	_ = v.RegisterValidation("phoneNumberRegion", ValidatePhoneNumberRegion)
	return &Validator{validator: v}
}

// Validate validates parameters
func (v *Validator) Validate(i interface{}) error {
	return v.validator.Struct(i)
}

// ValidateDate checks if the the value is valid date format
func ValidateDate(fl validator.FieldLevel) bool {
	if _, err := time.Parse(clock.RFC3339DATE, fl.Field().String()); err != nil {
		return false
	}
	return true
}

// ValidateRfc3339Datetime checks if the the value is valid RFC3339 format
func ValidateRfc3339Datetime(fl validator.FieldLevel) bool {
	if _, err := time.Parse(time.RFC3339, fl.Field().String()); err != nil {
		return false
	}
	return true
}

// ValidatePhoneNumber checks if the the value is valid phone number
// this validation must be used in a struct that has Region field
func ValidatePhoneNumber(fl validator.FieldLevel) bool {
	region := fl.Parent().Elem().FieldByName("Region")
	if !region.IsValid() {
		return false
	}
	num, err := libphonenumber.Parse(fl.Field().String(), region.String())
	if err != nil {
		return false
	}
	return libphonenumber.IsValidNumber(num)
}

// ValidatePhoneNumberRegion checks if the the value is valid region
func ValidatePhoneNumberRegion(fl validator.FieldLevel) bool {
	regionMap := libphonenumber.GetSupportedRegions()
	_, ok := regionMap[fl.Field().String()]
	return len(fl.Field().String()) != 0 && ok
}
