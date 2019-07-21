package repository

import (
	"fmt"
)

type repositoryError struct {
	err error
}

func newRepositoryError(err error) *repositoryError {
	return &repositoryError{err: err}
}

func (re *repositoryError) IsInternal() bool {
	return true
}
func (re *repositoryError) Error() string {
	return fmt.Sprint(re.err.Error())
}
