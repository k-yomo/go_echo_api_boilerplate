package fixture

import (
	"context"
	"github.com/k-yomo/go_echo_api_boilerplate/model"
	"github.com/pkg/errors"
	"testing"
)

// CreateTempUser creates temporary user fixture for test. phone number can be given to overwrite values.
func CreateTempUser(t *testing.T, db model.Execer, phoneNumber string) *model.TempUser {
	tu := model.NewTempUser("+818011112222")
	if phoneNumber != "" {
		tu.PhoneNumber = phoneNumber
	}
	if err := tu.Insert(context.Background(), db); err != nil {
		t.Error(errors.Wrap(err, "failed to create temp user"))
	}
	return tu
}
