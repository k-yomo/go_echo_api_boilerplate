package usecase

import (
	"fmt"
	"github.com/k-yomo/go_echo_boilerplate/internal/error_code"
)

type alreadyTakenError struct {
	column string
	value  interface{}
}

func newAlreadyTakenError(column string, value interface{}) *alreadyTakenError {
	return &alreadyTakenError{column, value}
}

func (ae *alreadyTakenError) Error() string {
	return fmt.Sprintf("%s: %s is already taken", ae.column, ae.value)
}

func (ae *alreadyTakenError) Code() error_code.ErrorCode {
	return error_code.AlreadyTaken
}

func (ae *alreadyTakenError) Messages() []string {
	return []string{ae.Error()}
}

type notFoundError struct {
	target string
	column string
	value  interface{}
}

func newNotFoundError(target string, column string, value interface{}) *notFoundError {
	return &notFoundError{target, column, value}
}

func (ne *notFoundError) Error() string {
	return fmt.Sprintf("%s with %s = %v is not found", ne.target, ne.column, ne.value)
}

func (ne *notFoundError) Code() error_code.ErrorCode {
	return error_code.NotFound
}

func (ne *notFoundError) Messages() []string {
	return []string{ne.Error()}
}

type unauthenticatedError struct{}

func newUnauthenticatedError() *unauthenticatedError {
	return &unauthenticatedError{}
}

func (*unauthenticatedError) Error() string {
	return "Unauthenticated"
}

func (*unauthenticatedError) Code() error_code.ErrorCode {
	return error_code.Unauthenticated
}

func (ue *unauthenticatedError) Messages() []string {
	return []string{ue.Error()}
}

type expiredError struct {
	err error
}

func newExpiredError(err error) *expiredError {
	return &expiredError{err: err}
}

func (ee *expiredError) Error() string {
	return ee.err.Error()
}

func (ee *expiredError) Code() error_code.ErrorCode {
	return error_code.Expired
}

func (ee *expiredError) Messages() []string {
	return []string{ee.Error()}
}

type badRequestError struct {
	err error
}

func newBadRequestError(err error) *badRequestError {
	return &badRequestError{err: err}
}

func (ee *badRequestError) Error() string {
	return ee.err.Error()
}

func (ee *badRequestError) Code() error_code.ErrorCode {
	return error_code.BadRequest
}

func (ee *badRequestError) Messages() []string {
	return []string{ee.Error()}
}
