package handler

import (
	"github.com/k-yomo/go_echo_api_boilerplate/internal/custom_context"
	"github.com/k-yomo/go_echo_api_boilerplate/usecase"
	"github.com/labstack/echo/v4"
	"net/http"
)

type healthcheckHandler struct {
	usecase usecase.HealthcheckUsecase
}

func NewHealthcheckHandler(e *echo.Group, us usecase.HealthcheckUsecase) {
	handler := &healthcheckHandler{
		usecase: us,
	}
	e.GET("/healthz", handler.CheckLiveness)
	e.GET("/readyz", handler.CheckReadiness)
}

// CheckLiveness godoc
// @Summary Check Liveness
// @Description check if application is living
// @Tags Health Check
// @Produce plain
// @Success 200 {string} string	"ok"
// @Router /healthz [get]
func (h *healthcheckHandler) CheckLiveness(ce echo.Context) error {
	return ce.String(http.StatusOK, "ok")
}

// CheckReadiness godoc
// @Summary Check Readiness
// @Description check if application and the depending services are functioning
// @Tags Health Check
// @Produce plain
// @Success 200 {string} string	"ok"
// @Failure 500 {string} string "ping db failed: invalid connection"
// @Router /readyz [get]
func (h *healthcheckHandler) CheckReadiness(ce echo.Context) error {
	ctx := &custom_context.Context{ce}
	err := h.usecase.CheckReadiness(ctx)
	if err != nil {
		return ce.String(http.StatusInternalServerError, err.Error())
	}
	return ce.String(http.StatusOK, "ok")
}
