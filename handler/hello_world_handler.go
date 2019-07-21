package handler

import (
	"fmt"
	"github.com/k-yomo/go_echo_boilerplate/internal/custom_context"
	"github.com/k-yomo/go_echo_boilerplate/usecase"
	"github.com/labstack/echo/v4"
	"net/http"
)

type HelloWorldHandler struct {
	usecase usecase.HelloWorldUsecase
}

func NewHelloWorldHandler(e *echo.Group, us usecase.HelloWorldUsecase) {
	handler := &HelloWorldHandler{
		usecase: us,
	}
	e.GET("/", handler.HelloWorld)
}

// HelloWorld godoc
// @Summary Greeting
// @Description Example endpoint that return greeting
// @Tags Example
// @Produce json
// @Param name query string false "Name for greeting"
// @Success 200 {string} string	"ok"
// @Router / [get]
func (h *HelloWorldHandler) HelloWorld(ce echo.Context) error {
	ctx := &custom_context.Context{ce}
	hw, err := h.usecase.Greet(ctx)
	if err != nil {
		return err
	}
	return ctx.String(http.StatusOK, fmt.Sprintf("Hello, %s!", hw.Name))
}
