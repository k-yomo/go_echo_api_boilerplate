package jwt_middleware

import (
	"github.com/k-yomo/go_echo_boilerplate/config"
	"github.com/k-yomo/go_echo_boilerplate/internal/error_code"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

// NewJWTMiddleware returns a middleware that handle jwt_generator.
func NewJWTMiddleware() echo.MiddlewareFunc {
	return middleware.JWTWithConfig(middleware.JWTConfig{SigningKey: []byte(config.JwtSigningKey), ErrorHandler: jwtErrorHandler})
}

func jwtErrorHandler(_ error) error {
	return &jwtError{}
}

type jwtError struct{}

func (je *jwtError) Error() string {
	return "Unauthenticated"
}

func (je *jwtError) Code() error_code.ErrorCode {
	return error_code.Unauthenticated
}

func (je *jwtError) Messages() []string {
	return []string{je.Error()}
}
