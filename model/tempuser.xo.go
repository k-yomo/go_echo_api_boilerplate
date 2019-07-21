// Package model contains the types for schema 'go_echo_boilerplate_development'.
package model

// GENERATED BY XO. DO NOT EDIT.

import (
	"context"
	"errors"
	"time"
)

// TempUser represents a row from 'temp_users'.
type TempUser struct {
	ID          uint64    `json:"id" db:"id"`                     // id
	PhoneNumber string    `json:"phone_number" db:"phone_number"` // phone_number
	AuthCode    string    `json:"auth_code" db:"auth_code"`       // auth_code
	AuthKey     string    `json:"auth_key" db:"auth_key"`         // auth_key
	CreatedAt   time.Time `json:"created_at" db:"created_at"`     // created_at
	UpdatedAt   time.Time `json:"updated_at" db:"updated_at"`     // updated_at

	// xo fields
	_exists, _deleted bool
}

// Exists determines if the TempUser exists in the database.
func (tu *TempUser) Exists() bool {
	return tu._exists
}

// GetTempUser gets a TempUser by primary key
func GetTempUser(ctx context.Context, db Queryer, key uint64) (*TempUser, error) {
	// sql query
	const sqlstr = `SELECT ` +
		`id, phone_number, auth_code, auth_key, created_at, updated_at ` +
		`FROM temp_users ` +
		`WHERE id = ?`

	var tu TempUser
	err := db.QueryRowxContext(ctx, sqlstr, key).Scan(&tu.ID, &tu.PhoneNumber, &tu.AuthCode, &tu.AuthKey, &tu.CreatedAt, &tu.UpdatedAt)
	if err != nil {
		return nil, err
	}
	return &tu, nil
}

// Deleted provides information if the TempUser has been deleted from the database.
func (tu *TempUser) Deleted() bool {
	return tu._deleted
}

// Insert inserts the TempUser to the database.
func (tu *TempUser) Insert(ctx context.Context, db Execer) error {
	// if already exist, bail
	if tu._exists {
		return errors.New("insert failed: already exists")
	}

	// sql insert query, primary key provided by autoincrement
	const sqlstr = `INSERT INTO temp_users (` +
		`phone_number, auth_code, auth_key, created_at, updated_at` +
		`) VALUES (` +
		`?, ?, ?, ?, ?` +
		`)`

	// run query
	XOLog(ctx, sqlstr, tu.PhoneNumber, tu.AuthCode, tu.AuthKey, tu.CreatedAt, tu.UpdatedAt)
	res, err := db.ExecContext(ctx, sqlstr, tu.PhoneNumber, tu.AuthCode, tu.AuthKey, tu.CreatedAt, tu.UpdatedAt)
	if err != nil {
		return err
	}

	// retrieve id
	id, err := res.LastInsertId()
	if err != nil {
		return err
	}

	// set primary key and existence
	tu.ID = uint64(id)
	tu._exists = true

	return nil
}

// Update updates the TempUser in the database.
func (tu *TempUser) Update(ctx context.Context, db Execer) error {
	// if doesn't exist, bail
	if !tu._exists {
		return errors.New("update failed: does not exist")
	}
	// if deleted, bail
	if tu._deleted {
		return errors.New("update failed: marked for deletion")
	}

	// sql query
	const sqlstr = `UPDATE temp_users SET ` +
		`phone_number = ?, auth_code = ?, auth_key = ?, created_at = ?, updated_at = ?` +
		` WHERE id = ?`
	// run query
	XOLog(ctx, sqlstr, tu.PhoneNumber, tu.AuthCode, tu.AuthKey, tu.CreatedAt, tu.UpdatedAt, tu.ID)
	_, err := db.ExecContext(ctx, sqlstr, tu.PhoneNumber, tu.AuthCode, tu.AuthKey, tu.CreatedAt, tu.UpdatedAt, tu.ID)
	return err
}

// Save saves the TempUser to the database.
func (tu *TempUser) Save(ctx context.Context, db Execer) error {
	if tu.Exists() {
		return tu.Update(ctx, db)
	}
	return tu.Insert(ctx, db)
}

// Delete deletes the TempUser from the database.
func (tu *TempUser) Delete(ctx context.Context, db Execer) error {
	// if doesn't exist, bail
	if !tu._exists {
		return nil
	}

	// if deleted, bail
	if tu._deleted {
		return nil
	}

	// sql query
	const sqlstr = `DELETE FROM temp_users WHERE id = ?`

	// run query
	XOLog(ctx, sqlstr, tu.ID)
	_, err := db.ExecContext(ctx, sqlstr, tu.ID)
	if err != nil {
		return err
	}

	// set deleted
	tu._deleted = true

	return nil
}

// TempUserByAuthCodeAuthKey retrieves a row from 'temp_users' as a TempUser.
//
// Generated from index 'auth_code_auth_key_idx'.
func TempUserByAuthCodeAuthKey(ctx context.Context, db Queryer, authCode string, authKey string) (*TempUser, error) {
	var err error

	// sql query
	const sqlstr = `SELECT ` +
		`id, phone_number, auth_code, auth_key, created_at, updated_at ` +
		`FROM temp_users ` +
		`WHERE auth_code = ? AND auth_key = ?`

	// run query
	XOLog(ctx, sqlstr, authCode, authKey)
	tu := TempUser{
		_exists: true,
	}

	err = db.QueryRowxContext(ctx, sqlstr, authCode, authKey).Scan(&tu.ID, &tu.PhoneNumber, &tu.AuthCode, &tu.AuthKey, &tu.CreatedAt, &tu.UpdatedAt)
	if err != nil {
		return nil, err
	}

	return &tu, nil
}

// TempUserByPhoneNumber retrieves a row from 'temp_users' as a TempUser.
//
// Generated from index 'phone_number'.
func TempUserByPhoneNumber(ctx context.Context, db Queryer, phoneNumber string) (*TempUser, error) {
	var err error

	// sql query
	const sqlstr = `SELECT ` +
		`id, phone_number, auth_code, auth_key, created_at, updated_at ` +
		`FROM temp_users ` +
		`WHERE phone_number = ?`

	// run query
	XOLog(ctx, sqlstr, phoneNumber)
	tu := TempUser{
		_exists: true,
	}

	err = db.QueryRowxContext(ctx, sqlstr, phoneNumber).Scan(&tu.ID, &tu.PhoneNumber, &tu.AuthCode, &tu.AuthKey, &tu.CreatedAt, &tu.UpdatedAt)
	if err != nil {
		return nil, err
	}

	return &tu, nil
}

// TempUserByID retrieves a row from 'temp_users' as a TempUser.
//
// Generated from index 'temp_users_id_pkey'.
func TempUserByID(ctx context.Context, db Queryer, id uint64) (*TempUser, error) {
	var err error

	// sql query
	const sqlstr = `SELECT ` +
		`id, phone_number, auth_code, auth_key, created_at, updated_at ` +
		`FROM temp_users ` +
		`WHERE id = ?`

	// run query
	XOLog(ctx, sqlstr, id)
	tu := TempUser{
		_exists: true,
	}

	err = db.QueryRowxContext(ctx, sqlstr, id).Scan(&tu.ID, &tu.PhoneNumber, &tu.AuthCode, &tu.AuthKey, &tu.CreatedAt, &tu.UpdatedAt)
	if err != nil {
		return nil, err
	}

	return &tu, nil
}