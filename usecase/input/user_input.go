package input

import "github.com/k-yomo/go_echo_api_boilerplate/pkg/clock"

// UpdateUserInput represents parameters for UpdateProfile usecase
type UpdateProfileInput struct {
	FirstName   string             `json:"firstName" validate:"required,min=1" example:"Taro | validation: min=1"`
	LastName    string             `json:"lastName" validate:"required,min=1" example:"Tanaka | validation: min=1"`
	Gender      string             `json:"gender" validate:"oneof=unknown male female" example:"unknown | male | female"`
	DateOfBirth clock.NullableDate `json:"dateOfBirth" validate:"omitempty,date" example:"1995-07-05 | validation: iso8601date"`
}

// UpdateEmailInput represents parameters for UpdateEmail usecase
type UpdateEmailInput struct {
	Email string `json:"email" validate:"required,email" example:"test@example.com | validation: email_format"`
}
