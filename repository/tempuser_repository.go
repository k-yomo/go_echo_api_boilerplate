package repository

import (
	"context"
	"database/sql"
	"github.com/jmoiron/sqlx"
	"github.com/k-yomo/go_echo_boilerplate/model"
	"github.com/pkg/errors"
)

// TempUserRepository represent the user's repository contract
type TempUserRepository interface {
	FindByID(ctx context.Context, id uint64) (*model.TempUser, error)
	FindByPhoneNumber(ctx context.Context, phoneNumber string) (*model.TempUser, error)
	FindByCodeAndKey(ctx context.Context, authCode string, authKey string) (*model.TempUser, error)
	Create(ctx context.Context, user *model.TempUser) error
	Update(ctx context.Context, user *model.TempUser) error
	UpsertByPhoneNumber(ctx context.Context, user *model.TempUser) (*model.TempUser, error)
	Destroy(ctx context.Context, user *model.TempUser) error
}

type tempUserRepository struct {
	DB *sqlx.DB
}

// NewUserRepository returns tempUserRepository
func NewTempUserRepository(db *sqlx.DB) *tempUserRepository {
	return &tempUserRepository{db}
}

// FindByID finds temporary user by id
func (tr *tempUserRepository) FindByID(ctx context.Context, id uint64) (*model.TempUser, error) {
	tempUser, err := model.TempUserByID(ctx, tr.DB, id)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, newRepositoryError(errors.Wrap(err, "find temp user by id failed"))
	}
	return tempUser, nil
}

// FindByPhoneNumber finds temporary user by phone number
func (tr *tempUserRepository) FindByPhoneNumber(ctx context.Context, phoneNumber string) (*model.TempUser, error) {
	tempUser, err := model.TempUserByPhoneNumber(ctx, tr.DB, phoneNumber)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, newRepositoryError(errors.Wrap(err, "find temp user by phone number failed"))
	}
	return tempUser, nil
}

// FindByCodeAndKey finds temporary user by auth code and auth key
func (tr *tempUserRepository) FindByCodeAndKey(ctx context.Context, authCode string, authKey string) (*model.TempUser, error) {
	tempUser, err := model.TempUserByAuthCodeAuthKey(ctx, tr.DB, authCode, authKey)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, newRepositoryError(errors.Wrap(err, "find temp user by auth code and auth key failed"))
	}
	return tempUser, nil
}

// Create creates a temporary user
func (tr *tempUserRepository) Create(ctx context.Context, tempUser *model.TempUser) error {
	err := tempUser.Insert(ctx, tr.DB)
	if err != nil {
		return newRepositoryError(errors.Wrap(err, "create temp user failed"))
	}
	return nil
}

// Update updates temporary user
func (tr *tempUserRepository) Update(ctx context.Context, tempUser *model.TempUser) error {
	err := tempUser.Update(ctx, tr.DB)
	if err != nil {
		return newRepositoryError(errors.Wrap(err, "update temp user failed"))
	}
	return nil
}

// Save creates or updates temporary user depending on phone number existence
func (tr *tempUserRepository) UpsertByPhoneNumber(ctx context.Context, tempUser *model.TempUser) (*model.TempUser, error) {
	tu, err := tr.FindByPhoneNumber(ctx, tempUser.PhoneNumber)
	if err != nil && err != sql.ErrNoRows {
		return nil, err
	}

	if tu != nil {
		tu.AuthKey = tempUser.AuthKey
		tu.AuthCode = tempUser.AuthCode
		tempUser = tu
	}

	if err := tempUser.Save(ctx, tr.DB); err != nil {
		return nil, newRepositoryError(errors.Wrap(err, "save temp user failed"))
	}
	return tempUser, nil
}

// Destroy deletes temporary user
func (tr *tempUserRepository) Destroy(ctx context.Context, tempUser *model.TempUser) error {
	err := tempUser.Delete(ctx, tr.DB)
	if err != nil {
		return newRepositoryError(errors.Wrap(err, "destroy temp user failed"))
	}
	return nil
}
