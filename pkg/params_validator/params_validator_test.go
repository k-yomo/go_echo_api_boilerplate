package params_validator

import (
	"github.com/stretchr/testify/assert"
	"gopkg.in/go-playground/validator.v9"
	"testing"
)

func TestNewValidator(t *testing.T) {
	assert.IsType(t, &validator.Validate{}, NewValidator().validator)
}

func TestValidate(t *testing.T) {
	type user struct {
		FirstName string `validate:"required"`
		Email     string `validate:"required,email"`
	}

	testCases := map[string]struct {
		input    interface{}
		expected error
	}{
		"valid struct":   {&user{"Kanji", "test@example.com"}, nil},
		"invalid struct": {&user{"Kanji", "testexample.com"}, validator.ValidationErrors{}},
	}

	v := NewValidator()
	for _, tc := range testCases {
		result := v.Validate(tc.input)
		assert.IsType(t, tc.expected, result)
	}
}
