package fixture

import (
	"context"
	"database/sql"
	"github.com/k-yomo/go_echo_api_boilerplate/model"
	"github.com/k-yomo/go_echo_api_boilerplate/pkg/clock"
	"github.com/pkg/errors"
	"gopkg.in/guregu/null.v3"
	"testing"
)

// CreateUser creates user fixture for test. User struct can be given to overwrite values.
func CreateUser(t *testing.T, db model.Execer, user *model.User) *model.User {
	u := model.NewUser("+818011112222", "test@example.com", "password", "Kanji", "Yomoda", model.GenderUnknown, null.NewTime(clock.Now(), true))
	u.UnconfirmedPhoneNumber = sql.NullString{String: "+818099998888", Valid: true}
	if user != nil {
		if user.PhoneNumber != "" {
			u.PhoneNumber = user.PhoneNumber
		}
		if user.UnconfirmedPhoneNumber.Valid {
			u.UnconfirmedPhoneNumber = user.UnconfirmedPhoneNumber
		}
		if user.Email != "" {
			u.Email = user.Email
		}
		if len(user.PasswordDigest) != 0 {
			u.PasswordDigest = user.PasswordDigest
		}
		if user.FirstName != "" {
			u.FirstName = user.FirstName
		}
		if user.LastName != "" {
			u.LastName = user.LastName
		}
		if user.Gender != 0 {
			u.Gender = user.Gender
		}
		if user.DateOfBirth.Valid {
			u.DateOfBirth = user.DateOfBirth
		}

	}
	if err := u.Insert(context.Background(), db); err != nil {
		t.Error(errors.Wrap(err, "failed to create user"))
	}
	return u
}
