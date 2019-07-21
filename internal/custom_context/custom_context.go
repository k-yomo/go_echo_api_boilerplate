package custom_context

import (
	"github.com/labstack/echo/v4"
	"gopkg.in/go-playground/validator.v9"
	"strings"
)

// CustomContext is a wrapper of echo.Context
type Context struct {
	echo.Context
}

// BindWithValidation binds request params into given argument then validates
func (c *Context) BindWithValidation(i interface{}) error {
	if err := c.Bind(i); err != nil {
		return &bindError{
			info:    "request failed",
			message: err.Error(),
		}
	}

	if err := c.Validate(i); err != nil {
		errors := err.(validator.ValidationErrors)
		return &validationError{
			info:     "validate failed",
			messages: strings.Split(errors.Error(), "\n"),
		}
	}
	return nil
}
