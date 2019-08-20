package custom_context

import (
	"github.com/k-yomo/go_echo_api_boilerplate/pkg/params_validator"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

type mockStruct struct {
	Name  string `json:"name" validate:"required"`
	Email string `json:"email" validate:"required,email"`
}

func TestCustomContextBindWithValidation(t *testing.T) {
	testCases := map[string]struct {
		context  *Context
		expected error
	}{
		"valid json":                 {newMockContext(`{"name": "Bob", "email": "test@example.com"}`), nil},
		"invalid json format":        {newMockContext(`{"aa"}`), &bindError{"request Failed", ""}},
		"invalid args(email format)": {newMockContext(`{"name":"Bob","email":"testexamplecom"}`), &validationError{"validate failed", []string{""}}},
	}

	for _, tc := range testCases {
		err := tc.context.BindWithValidation(new(mockStruct))
		assert.IsType(t, tc.expected, err)
	}
}

func newMockContext(jsonStr string) *Context {
	e := echo.New()
	e.Validator = params_validator.NewValidator()
	req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(jsonStr))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	return &Context{e.NewContext(req, httptest.NewRecorder())}
}
