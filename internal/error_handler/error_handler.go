package error_handler

import (
	"fmt"
	"github.com/k-yomo/go_echo_api_boilerplate/internal/custom_context"
	"github.com/k-yomo/go_echo_api_boilerplate/internal/error_code"
	"github.com/labstack/echo/v4"
	"github.com/pkg/errors"
	"net/http"
)

// generalError interface should be implemented if the errors are suppose to be handled by client
type generalError interface {
	// Code returns ErrorCode to let client know what kind of error
	Code() error_code.ErrorCode
	// Messages returns error details
	Messages() []string
}

// internalError interface should be implemented by errors that should be handled by service provider.
type internalError interface {
	// IsInternal returns bool that tells if the error is internal error or not
	IsInternal() bool
}

// ErrorResponse represents api error response
type ErrorResponse struct {
	Code   error_code.ErrorCode `json:"code"`
	Errors []string             `json:"errors"`
}

func newErrorResponse(code error_code.ErrorCode, errs []string) *ErrorResponse {
	return &ErrorResponse{Code: code, Errors: errs}
}

// HTTPErrorHandler handles error depending on the error type(general or internal)
func HTTPErrorHandler(err error, ec echo.Context) {
	c := &custom_context.Context{ec}
	// is general error
	if ge, ok := errors.Cause(err).(generalError); ok {
		c.Logger().Info("error_type", "general error", "error", err.Error(), "error_code", ge.Code().String())
		httpStatus := error_code.GetHTTPStatusByErrorCode(ge.Code())
		errorRes := newErrorResponse(ge.Code(), ge.Messages())
		c.JSON(httpStatus, errorRes)
		return
	}

	// is not found
	ee, ok := errors.Cause(err).(*echo.HTTPError)
	if ok {
		if ee.Code == http.StatusNotFound {
			c.Logger().Info(fmt.Sprintf("Not found: %s", err.Error()))
			errorRes := newErrorResponse(error_code.NotFound, []string{"Not Found"})
			c.JSON(ee.Code, errorRes)
			return
		}
		if ee.Code == http.StatusMethodNotAllowed {
			c.Logger().Info(fmt.Sprintf("Method not allowed: %s", err.Error()))
			errorRes := newErrorResponse(error_code.MethodNotAllowed, []string{"Method not allowed"})
			c.JSON(ee.Code, errorRes)
			return
		}
	}

	// is internal error
	ie, ok := errors.Cause(err).(internalError)
	if ok && ie.IsInternal() {
		c.Logger().Error("error_type", "internal server error", "error", errors.Cause(err), "error_code", error_code.InternalError.String())
	} else {
		c.Logger().Error("error_type", "unexpected server error", "error", errors.WithStack(err), "error_code", error_code.UnhandledError.String())
	}

	errorRes := newErrorResponse(error_code.InternalError, []string{"Internal server error"})
	c.JSON(http.StatusInternalServerError, errorRes)
}
