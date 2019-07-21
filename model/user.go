package model

import (
	"context"
	"database/sql"
	"github.com/go-sql-driver/mysql"
	"github.com/k-yomo/go_echo_boilerplate/pkg/clock"
	"github.com/pkg/errors"
	"golang.org/x/crypto/bcrypt"
	"gopkg.in/guregu/null.v3"
)

// Gender represents user's gender
type Gender int

const (
	GenderUnknown Gender = iota
	GenderMale
	GenderFemale
)

var genderStrMapping = map[string]Gender{
	"unknown": GenderUnknown,
	"male":    GenderMale,
	"female":  GenderFemale,
}

func (g Gender) String() string {
	switch g {
	case GenderMale:
		return "male"
	case GenderFemale:
		return "female"
	default:
		return "unknown"
	}
}

// StringToGender converts string to gender
func StringToGender(s string) Gender {
	g, ok := genderStrMapping[s]
	if ok {
		return g
	} else {
		return GenderUnknown
	}
}

// Int64ToGender converts int64 to gender
func Int64ToGender(i int64) Gender {
	switch i {
	case int64(GenderMale):
		return GenderMale
	case int64(GenderFemale):
		return GenderFemale
	default:
		return GenderUnknown
	}
}

// GenderString returns gender string
func (u *User) GenderString() string {
	return Int64ToGender(u.Gender).String()
}

// NewUser returns initialized user
func NewUser(phoneNumber, email, password, firstName, lastName string, gender Gender, dateOfBirth null.Time) *User {
	passwordDigest, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return &User{
		PhoneNumber:    phoneNumber,
		Email:          email,
		PasswordDigest: passwordDigest,
		FirstName:      firstName,
		LastName:       lastName,
		Gender:         int64(gender),
		DateOfBirth:    mysql.NullTime{Time: dateOfBirth.Time, Valid: dateOfBirth.Valid},
		CreatedAt:      clock.Now(),
		UpdatedAt:      clock.Now(),
	}
}

// SMSReconfirmation returns the SMSReconfirmation associated with the User's ID (id).
func (u *User) SMSReconfirmation(ctx context.Context, db Queryer) (*SmsReconfirmation, error) {
	return SmsReconfirmationByUserID(ctx, db, u.ID)
}

func (u *User) ConfirmPhoneNumber() error {
	if !u.UnconfirmedPhoneNumber.Valid {
		return errors.New("unconfirmed phone number is not set")
	}
	u.PhoneNumber = u.UnconfirmedPhoneNumber.String
	u.UnconfirmedPhoneNumber = sql.NullString{String: "", Valid: false}
	return nil
}
