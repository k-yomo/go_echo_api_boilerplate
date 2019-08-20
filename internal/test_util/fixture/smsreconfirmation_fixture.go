package fixture

import (
	"context"
	"database/sql"
	"github.com/k-yomo/go_echo_api_boilerplate/model"
	"github.com/pkg/errors"
	"testing"
)

// CreateSMSReconfirmation creates sms reconfirmation fixture for test.
// SmsReconfirmation and user struct can be given to overwrite values.
func CreateSMSReconfirmation(t *testing.T, db model.Execer, smsReconfirmation *model.SmsReconfirmation, user *model.User) *model.SmsReconfirmation {
	if user == nil || !user.Exists() {
		if smsReconfirmation != nil && smsReconfirmation.PhoneNumber != "" {
			user.UnconfirmedPhoneNumber = sql.NullString{String: smsReconfirmation.PhoneNumber, Valid: true}
		}
		user = CreateUser(t, db, user)
	}

	sr := model.NewSMSReconfirmation(user.ID, user.UnconfirmedPhoneNumber.String)
	if err := sr.Insert(context.Background(), db); err != nil {
		t.Error(errors.Wrap(err, "failed to create sms reconfirmation"))
	}
	return sr
}
