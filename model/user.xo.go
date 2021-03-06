// Package model contains the types for schema 'go_echo_api_boilerplate_development'.
package model

// GENERATED BY XO. DO NOT EDIT.

import (
	"context"
	"database/sql"
	"errors"
	"time"

	"github.com/go-sql-driver/mysql"
)

// User represents a row from 'users'.
type User struct {
	ID                     uint64         `json:"id" db:"id"`                                             // id
	FirstName              string         `json:"first_name" db:"first_name"`                             // first_name
	LastName               string         `json:"last_name" db:"last_name"`                               // last_name
	Gender                 int64          `json:"gender" db:"gender"`                                     // gender
	DateOfBirth            mysql.NullTime `json:"date_of_birth" db:"date_of_birth"`                       // date_of_birth
	PhoneNumber            string         `json:"phone_number" db:"phone_number"`                         // phone_number
	UnconfirmedPhoneNumber sql.NullString `json:"unconfirmed_phone_number" db:"unconfirmed_phone_number"` // unconfirmed_phone_number
	Email                  string         `json:"email" db:"email"`                                       // email
	PasswordDigest         []byte         `json:"password_digest" db:"password_digest"`                   // password_digest
	CreatedAt              time.Time      `json:"created_at" db:"created_at"`                             // created_at
	UpdatedAt              time.Time      `json:"updated_at" db:"updated_at"`                             // updated_at

	// xo fields
	_exists, _deleted bool
}

// Exists determines if the User exists in the database.
func (u *User) Exists() bool {
	return u._exists
}

// GetUser gets a User by primary key
func GetUser(ctx context.Context, db Queryer, key uint64) (*User, error) {
	// sql query
	const sqlstr = `SELECT ` +
		`id, first_name, last_name, gender, date_of_birth, phone_number, unconfirmed_phone_number, email, password_digest, created_at, updated_at ` +
		`FROM users ` +
		`WHERE id = ?`

	var u User
	err := db.QueryRowxContext(ctx, sqlstr, key).Scan(&u.ID, &u.FirstName, &u.LastName, &u.Gender, &u.DateOfBirth, &u.PhoneNumber, &u.UnconfirmedPhoneNumber, &u.Email, &u.PasswordDigest, &u.CreatedAt, &u.UpdatedAt)
	if err != nil {
		return nil, err
	}
	return &u, nil
}

// Deleted provides information if the User has been deleted from the database.
func (u *User) Deleted() bool {
	return u._deleted
}

// Insert inserts the User to the database.
func (u *User) Insert(ctx context.Context, db Execer) error {
	// if already exist, bail
	if u._exists {
		return errors.New("insert failed: already exists")
	}

	// sql insert query, primary key provided by autoincrement
	const sqlstr = `INSERT INTO users (` +
		`first_name, last_name, gender, date_of_birth, phone_number, unconfirmed_phone_number, email, password_digest, created_at, updated_at` +
		`) VALUES (` +
		`?, ?, ?, ?, ?, ?, ?, ?, ?, ?` +
		`)`

	// run query
	XOLog(ctx, sqlstr, u.FirstName, u.LastName, u.Gender, u.DateOfBirth, u.PhoneNumber, u.UnconfirmedPhoneNumber, u.Email, u.PasswordDigest, u.CreatedAt, u.UpdatedAt)
	res, err := db.ExecContext(ctx, sqlstr, u.FirstName, u.LastName, u.Gender, u.DateOfBirth, u.PhoneNumber, u.UnconfirmedPhoneNumber, u.Email, u.PasswordDigest, u.CreatedAt, u.UpdatedAt)
	if err != nil {
		return err
	}

	// retrieve id
	id, err := res.LastInsertId()
	if err != nil {
		return err
	}

	// set primary key and existence
	u.ID = uint64(id)
	u._exists = true

	return nil
}

// Update updates the User in the database.
func (u *User) Update(ctx context.Context, db Execer) error {
	// if doesn't exist, bail
	if !u._exists {
		return errors.New("update failed: does not exist")
	}
	// if deleted, bail
	if u._deleted {
		return errors.New("update failed: marked for deletion")
	}

	// sql query
	const sqlstr = `UPDATE users SET ` +
		`first_name = ?, last_name = ?, gender = ?, date_of_birth = ?, phone_number = ?, unconfirmed_phone_number = ?, email = ?, password_digest = ?, created_at = ?, updated_at = ?` +
		` WHERE id = ?`
	// run query
	XOLog(ctx, sqlstr, u.FirstName, u.LastName, u.Gender, u.DateOfBirth, u.PhoneNumber, u.UnconfirmedPhoneNumber, u.Email, u.PasswordDigest, u.CreatedAt, u.UpdatedAt, u.ID)
	_, err := db.ExecContext(ctx, sqlstr, u.FirstName, u.LastName, u.Gender, u.DateOfBirth, u.PhoneNumber, u.UnconfirmedPhoneNumber, u.Email, u.PasswordDigest, u.CreatedAt, u.UpdatedAt, u.ID)
	return err
}

// Save saves the User to the database.
func (u *User) Save(ctx context.Context, db Execer) error {
	if u.Exists() {
		return u.Update(ctx, db)
	}
	return u.Insert(ctx, db)
}

// Delete deletes the User from the database.
func (u *User) Delete(ctx context.Context, db Execer) error {
	// if doesn't exist, bail
	if !u._exists {
		return nil
	}

	// if deleted, bail
	if u._deleted {
		return nil
	}

	// sql query
	const sqlstr = `DELETE FROM users WHERE id = ?`

	// run query
	XOLog(ctx, sqlstr, u.ID)
	_, err := db.ExecContext(ctx, sqlstr, u.ID)
	if err != nil {
		return err
	}

	// set deleted
	u._deleted = true

	return nil
}

// UserByEmail retrieves a row from 'users' as a User.
//
// Generated from index 'email'.
func UserByEmail(ctx context.Context, db Queryer, email string) (*User, error) {
	var err error

	// sql query
	const sqlstr = `SELECT ` +
		`id, first_name, last_name, gender, date_of_birth, phone_number, unconfirmed_phone_number, email, password_digest, created_at, updated_at ` +
		`FROM users ` +
		`WHERE email = ?`

	// run query
	XOLog(ctx, sqlstr, email)
	u := User{
		_exists: true,
	}

	err = db.QueryRowxContext(ctx, sqlstr, email).Scan(&u.ID, &u.FirstName, &u.LastName, &u.Gender, &u.DateOfBirth, &u.PhoneNumber, &u.UnconfirmedPhoneNumber, &u.Email, &u.PasswordDigest, &u.CreatedAt, &u.UpdatedAt)
	if err != nil {
		return nil, err
	}

	return &u, nil
}

// UserByPhoneNumber retrieves a row from 'users' as a User.
//
// Generated from index 'phone_number'.
func UserByPhoneNumber(ctx context.Context, db Queryer, phoneNumber string) (*User, error) {
	var err error

	// sql query
	const sqlstr = `SELECT ` +
		`id, first_name, last_name, gender, date_of_birth, phone_number, unconfirmed_phone_number, email, password_digest, created_at, updated_at ` +
		`FROM users ` +
		`WHERE phone_number = ?`

	// run query
	XOLog(ctx, sqlstr, phoneNumber)
	u := User{
		_exists: true,
	}

	err = db.QueryRowxContext(ctx, sqlstr, phoneNumber).Scan(&u.ID, &u.FirstName, &u.LastName, &u.Gender, &u.DateOfBirth, &u.PhoneNumber, &u.UnconfirmedPhoneNumber, &u.Email, &u.PasswordDigest, &u.CreatedAt, &u.UpdatedAt)
	if err != nil {
		return nil, err
	}

	return &u, nil
}

// UserByID retrieves a row from 'users' as a User.
//
// Generated from index 'users_id_pkey'.
func UserByID(ctx context.Context, db Queryer, id uint64) (*User, error) {
	var err error

	// sql query
	const sqlstr = `SELECT ` +
		`id, first_name, last_name, gender, date_of_birth, phone_number, unconfirmed_phone_number, email, password_digest, created_at, updated_at ` +
		`FROM users ` +
		`WHERE id = ?`

	// run query
	XOLog(ctx, sqlstr, id)
	u := User{
		_exists: true,
	}

	err = db.QueryRowxContext(ctx, sqlstr, id).Scan(&u.ID, &u.FirstName, &u.LastName, &u.Gender, &u.DateOfBirth, &u.PhoneNumber, &u.UnconfirmedPhoneNumber, &u.Email, &u.PasswordDigest, &u.CreatedAt, &u.UpdatedAt)
	if err != nil {
		return nil, err
	}

	return &u, nil
}
