package error_code

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestErrorCodeString(t *testing.T) {
	testCases := []struct {
		errorCode ErrorCode
		expected  string
	}{
		{BadRequest, "BadRequest"},
		{InvalidHeader, "InvalidHeader"},
		{Unauthenticated, "Unauthenticated"},
		{Expired, "Expired"},
		{InvalidParams, "InvalidParams"},
		{NotFound, "NotFound"},
		{MethodNotAllowed, "MethodNotAllowed"},
		{AlreadyTaken, "AlreadyTaken"},
		{InternalError, "InternalError"},
		{UnhandledError, "UnhandledError"},
	}
	for _, tc := range testCases {
		assert.IsType(t, tc.expected, tc.errorCode.String())
	}
}

func TestGetHTTPStatusByErrorCode(t *testing.T) {
	testCases := []struct {
		input    ErrorCode
		expected int
	}{
		{BadRequest, 400},
		{InvalidHeader, 400},
		{Unauthenticated, 401},
		{Expired, 401},
		{InvalidParams, 422},
		{NotFound, 404},
		{MethodNotAllowed, 405},
		{AlreadyTaken, 409},
		{InternalError, 500},
		{UnhandledError, 500},
	}
	for _, tc := range testCases {
		result := GetHTTPStatusByErrorCode(tc.input)
		assert.IsType(t, tc.expected, result)
	}
}
