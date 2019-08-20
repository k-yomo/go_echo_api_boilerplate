package output

import (
	"github.com/k-yomo/go_echo_api_boilerplate/model"
	"github.com/k-yomo/go_echo_api_boilerplate/pkg/jwt_generator"
	"github.com/ttacon/libphonenumber"
	"gopkg.in/guregu/null.v3"
	"time"
)

// ConfirmOutput represents response body for current user
type CurrentUserOutput struct {
	ID          uint64      `json:"id" example:"1"`
	PhoneNumber string      `json:"phoneNumber" example:"080-1111-2222"`
	FirstName   string      `json:"firstName" example:"Kanji"`
	LastName    string      `json:"lastName" example:"Yomoda"`
	DateOfBirth null.String `json:"dateOfBirth" example:"1995-07-05"`
	Gender      string      `json:"gender" example:"unknown | male | female"`
	Email       string      `json:"email" example:"test@example.com"`
	CreatedAt   string      `json:"createdAt" example:"2020-01-01T00:00:00+09:00"`
	UpdatedAt   string      `json:"updatedAt" example:"2020-01-01T00:00:00+09:00"`
	AuthToken   string      `json:"authToken" example:"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE1NjI4NDY5MTUsImlkIjo1fQ.2VZWDiWu3KDULr8p0MlPwxbTKJGnHLhcg3L_Ishx9e4"`
}

// NewCurrentUserOutput returns initialized current user
func NewCurrentUserOutput(u *model.User) *CurrentUserOutput {
	phoneNumber, _ := libphonenumber.Parse(u.PhoneNumber, "")
	token, _ := jwt_generator.GenerateJwt(u.ID)
	return &CurrentUserOutput{
		ID:          u.ID,
		PhoneNumber: libphonenumber.Format(phoneNumber, libphonenumber.NATIONAL),
		Email:       u.Email,
		FirstName:   u.FirstName,
		LastName:    u.LastName,
		Gender:      u.GenderString(),
		DateOfBirth: null.NewString(u.DateOfBirth.Time.Format("2006-01-02"), u.DateOfBirth.Valid),
		CreatedAt:   u.CreatedAt.Format(time.RFC3339),
		UpdatedAt:   u.CreatedAt.Format(time.RFC3339),
		AuthToken:   token,
	}
}
