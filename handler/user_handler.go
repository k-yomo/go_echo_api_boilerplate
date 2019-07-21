package handler

import (
	"github.com/k-yomo/go_echo_boilerplate/internal/custom_context"
	"github.com/k-yomo/go_echo_boilerplate/usecase"
	"github.com/labstack/echo/v4"
	"net/http"
)

type userHandler struct {
	usecase usecase.UserUsecase
}

func NewUserHandler(e *echo.Group, us usecase.UserUsecase, authMiddleWare echo.MiddlewareFunc) {
	handler := &userHandler{us}
	e.GET("/users/self", handler.GetProfile, authMiddleWare)
	e.PATCH("/users/self", handler.UpdateProfile, authMiddleWare)
	e.PATCH("/users/self/email", handler.UpdateEmail, authMiddleWare)
}

// GetProfile godoc
// @Summary Get Current User Profile
// @Description ログインユーザーのプロフィール取得
// @Tags User
// @Produce json
// @Security JWTAuth
// @Success 200 {object} output.CurrentUserOutput
// @Success 401 {object} error_handler.ErrorResponse "Unauthenticated"
// @Router /users/self [get]
func (h *userHandler) GetProfile(ce echo.Context) error {
	ctx := &custom_context.Context{ce}
	user, err := h.usecase.GetProfile(ctx)
	if err != nil {
		return err
	}
	return ctx.JSON(http.StatusOK, user)
}

// UpdateProfile godoc
// @Summary Update Current User Profile
// @Description ログインユーザーのプロフィール更新
// @Tags User
// @Accept json
// @Produce json
// @Security JWTAuth
// @Param body body input.UpdateProfileInput true "body"
// @Success 200 {object} output.CurrentUserOutput
// @Success 400 {object} error_handler.ErrorResponse "Bad request"
// @Success 422 {object} error_handler.ErrorResponse "Invalid date format"
// @Router /users/self [put]
func (h *userHandler) UpdateProfile(ce echo.Context) error {
	ctx := &custom_context.Context{ce}
	user, err := h.usecase.UpdateProfile(ctx)
	if err != nil {
		return err
	}
	return ctx.JSON(http.StatusOK, user)
}

// UpdateEmail godoc
// @Summary Update Current User Email
// @Description ログインユーザーのEメール更新
// @Tags User
// @Accept json
// @Produce json
// @Security JWTAuth
// @Param body body input.UpdateEmailInput true "body"
// @Success 200 {object} output.CurrentUserOutput
// @Success 400 {object} error_handler.ErrorResponse "Bad request"
// @Success 409 {object} error_handler.ErrorResponse "Email is already taken"
// @Success 422 {object} error_handler.ErrorResponse "Invalid email format"
// @Router /users/self/email [patch]
func (h *userHandler) UpdateEmail(ce echo.Context) error {
	ctx := &custom_context.Context{ce}
	user, err := h.usecase.UpdateEmail(ctx)
	if err != nil {
		return err
	}
	return ctx.JSON(http.StatusOK, user)
}
