package repository

import (
	"context"
	"database/sql"
	"github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"github.com/k-yomo/go_echo_api_boilerplate/model"
	"github.com/pkg/errors"
	"gopkg.in/guregu/null.v3"
)

// UserRepository represent the user's repository contract
type UserRepository interface {
	FindByID(ctx context.Context, id uint64) (*model.User, error)
	FindByPhoneNumber(ctx context.Context, phoneNumber string) (*model.User, error)
	FindByEmail(ctx context.Context, email string) (*model.User, error)
	Create(ctx context.Context, user *model.User) error
	Update(ctx context.Context, user *model.User) error
	UpdateProfile(ctx context.Context, user *model.User, firstName, lastName string, gender model.Gender, dateOfBirth null.Time) error
	FindSMSReconfirmationByUser(ctx context.Context, user *model.User) (*model.SmsReconfirmation, error)
	FindSMSReconfirmationByPhoneNumber(ctx context.Context, phoneNumber string) (*model.SmsReconfirmation, error)
	SaveSMSReconfirmation(ctx context.Context, smsReconfirmation *model.SmsReconfirmation) error
	DestroySMSReconfirmation(ctx context.Context, smsReconfirmation *model.SmsReconfirmation) error
}

type userRepository struct {
	DB *sqlx.DB
}

// NewUserRepository returns userRepository
func NewUserRepository(db *sqlx.DB) *userRepository {
	return &userRepository{db}
}

// FindByID finds user by id
func (ur *userRepository) FindByID(ctx context.Context, id uint64) (*model.User, error) {
	u, err := model.UserByID(ctx, ur.DB, id)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, &repositoryError{errors.Wrap(err, "find user by id failed")}
	}
	return u, nil
}

// FindByEmail finds user by email
func (ur *userRepository) FindByEmail(ctx context.Context, email string) (*model.User, error) {
	u, err := model.UserByEmail(ctx, ur.DB, email)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, &repositoryError{errors.Wrap(err, "find user by email failed")}
	}
	return u, nil
}

// FindByEmail finds user by phone number
func (ur *userRepository) FindByPhoneNumber(ctx context.Context, phoneNumber string) (*model.User, error) {
	u, err := model.UserByPhoneNumber(ctx, ur.DB, phoneNumber)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, &repositoryError{errors.Wrap(err, "find user by phone number failed")}
	}
	return u, nil
}

// Create creates a user
func (ur *userRepository) Create(ctx context.Context, user *model.User) error {
	err := user.Insert(ctx, ur.DB)
	if err != nil {
		return &repositoryError{errors.Wrap(err, "create user failed")}
	}
	return nil
}

// Update updates user
func (ur *userRepository) Update(ctx context.Context, user *model.User) error {
	err := user.Update(ctx, ur.DB)
	if err != nil {
		return &repositoryError{errors.Wrap(err, "update user failed")}
	}
	return nil
}

// UpdateProfile updates user with given arguments
func (ur *userRepository) UpdateProfile(ctx context.Context, user *model.User, firstName, lastName string, gender model.Gender, dateOfBirth null.Time) error {
	user.FirstName = firstName
	user.LastName = lastName
	user.Gender = int64(gender)
	user.DateOfBirth = mysql.NullTime{Time: dateOfBirth.Time, Valid: dateOfBirth.Valid}
	err := user.Update(ctx, ur.DB)
	if err != nil {
		return &repositoryError{errors.Wrap(err, "update user failed")}
	}
	return nil
}

// FindSMSReconfirmationByUser finds sms reconfirmation by associated user
func (ur *userRepository) FindSMSReconfirmationByUser(ctx context.Context, user *model.User) (*model.SmsReconfirmation, error) {
	smsReconfirmation, err := user.SMSReconfirmation(ctx, ur.DB)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, &repositoryError{errors.Wrap(err, "find sms reconfirmation by user failed")}
	}
	return smsReconfirmation, nil
}

// FindSMSReconfirmationByPhoneNumber finds sms reconfirmation by phone number
func (ur *userRepository) FindSMSReconfirmationByPhoneNumber(ctx context.Context, phoneNumber string) (*model.SmsReconfirmation, error) {
	smsReconfirmation, err := model.SmsReconfirmationByPhoneNumber(ctx, ur.DB, phoneNumber)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, &repositoryError{errors.Wrap(err, "find sms reconfirmation by phone number failed")}
	}
	return smsReconfirmation, nil
}

// SaveSMSReconfirmation saves sms reconfirmation
func (ur *userRepository) SaveSMSReconfirmation(ctx context.Context, smsReconfirmation *model.SmsReconfirmation) error {
	err := smsReconfirmation.Save(ctx, ur.DB)
	if err != nil {
		return &repositoryError{errors.Wrap(err, "save sms reconfirmation failed")}
	}
	return nil
}

// DestroySMSReconfirmation deletes sms reconfirmation
func (ur *userRepository) DestroySMSReconfirmation(ctx context.Context, smsReconfirmation *model.SmsReconfirmation) error {
	err := smsReconfirmation.Delete(ctx, ur.DB)
	if err != nil {
		return newRepositoryError(errors.Wrap(err, "destroy sms reconfirmation failed"))
	}
	return nil
}
