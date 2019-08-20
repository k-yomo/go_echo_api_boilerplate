package custom_context

import (
	"github.com/bmizerany/assert"
	"github.com/k-yomo/go_echo_api_boilerplate/internal/error_code"
	"testing"
)

var be = &bindError{"request failed", "invalid json format"}
var ve = &validationError{"request failed", []string{"name is required", "invalid email format"}}

func TestBindErrorError(t *testing.T) {
	assert.Equal(t, "request failed: invalid json format", be.Error())
}

func TestBindErrorCode(t *testing.T) {
	assert.Equal(t, error_code.BadRequest, be.Code())
}

func TestBindErrorMessages(t *testing.T) {
	assert.Equal(t, []string{"invalid json format"}, be.Messages())
}

func TestValidationErrorError(t *testing.T) {
	assert.Equal(t, "request failed: name is required, invalid email format", ve.Error())
}

func TestValidationErrorCode(t *testing.T) {
	assert.Equal(t, error_code.InvalidParams, ve.Code())
}

func TestValidationErrorMessages(t *testing.T) {
	assert.Equal(t, []string{"name is required", "invalid email format"}, ve.Messages())
}
