package custom_context

import (
	"fmt"
	"github.com/k-yomo/go_echo_api_boilerplate/internal/error_code"
	"strings"
)

type bindError struct {
	info    string
	message string
}

func (be *bindError) Error() string {
	return fmt.Sprintf("%s: %s", be.info, be.message)
}

func (be *bindError) Code() error_code.ErrorCode {
	return error_code.BadRequest
}

func (be *bindError) Messages() []string {
	return []string{be.message}
}

type validationError struct {
	info     string
	messages []string
}

func (ve *validationError) Error() string {
	return fmt.Sprintf("%s: %s", ve.info, strings.Join(ve.messages, ", "))
}

func (ve *validationError) Code() error_code.ErrorCode {
	return error_code.InvalidParams
}

func (ve *validationError) Messages() []string {
	return ve.messages
}
