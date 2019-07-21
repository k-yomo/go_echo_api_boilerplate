package handler

import (
	"github.com/k-yomo/go_echo_boilerplate/usecase"
	"github.com/labstack/echo/v4"
	"github.com/swaggo/echo-swagger"
)

// Handler struct
type Handler struct {
	HelloWorldUsecase usecase.HelloWorldUsecase
}

// NewHandler returns initialized Handler
func NewHandler(e *echo.Echo, us *usecase.Usecase, authMiddleWare echo.MiddlewareFunc) {
	e.GET("/swagger/*", echoSwagger.WrapHandler)
	api := e.Group("/v1")
	// this is example
	NewHelloWorldHandler(api, us.HelloWorldUsecase)
	NewAuthHandler(api, us.AuthUsecase, authMiddleWare)
	NewUserHandler(api, us.UserUsecase, authMiddleWare)
}
