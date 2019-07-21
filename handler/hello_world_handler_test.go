package handler

import (
	"github.com/k-yomo/go_echo_boilerplate/pkg/params_validator"
	"github.com/k-yomo/go_echo_boilerplate/usecase"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHelloWorld(t *testing.T) {
	// Setup
	e := echo.New()
	e.Validator = params_validator.NewValidator()

	req := httptest.NewRequest(http.MethodGet, "/?name=World", nil)
	rec := httptest.NewRecorder()

	testCases := []struct {
		input          usecase.HelloWorldUsecase
		ExpectedStatus int
	}{
		{usecase.NewHelloWorldUsecase(), http.StatusOK},
		// we should test the case that HelloWorldUsecase returns error, but we don't since this is example...
	}

	for _, tc := range testCases {
		h := HelloWorldHandler{usecase: tc.input}
		if assert.NoError(t, h.HelloWorld(e.NewContext(req, rec))) {
			assert.Equal(t, tc.ExpectedStatus, rec.Code)
		}
	}
}
