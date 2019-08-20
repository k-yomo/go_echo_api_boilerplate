package handler

import (
	"github.com/k-yomo/go_echo_api_boilerplate/internal/custom_context"
	"github.com/k-yomo/go_echo_api_boilerplate/usecase"
	"github.com/labstack/echo/v4"
	"net/http"
)

type authHandler struct {
	usecase usecase.AuthUsecase
}

func NewAuthHandler(e *echo.Group, us usecase.AuthUsecase, authMiddleWare echo.MiddlewareFunc) {
	handler := &authHandler{us}
	g := e.Group("/auth")
	g.POST("/temp_sign_up", handler.TempSignUp)
	g.POST("/confirm", handler.ConfirmTempUser)
	g.POST("/sign_in", handler.SignIn)
	g.POST("/sign_up", handler.SignUp, authMiddleWare)
	g.PATCH("/phone_number", handler.UpdateUnconfirmedPhoneNumber, authMiddleWare)
	g.POST("/phone_number/confirm", handler.ConfirmPhoneNumber, authMiddleWare)
}

// TempSignUp
// @Summary Temporary Sign Up
// @Description SMS認証用の仮ユーザー登録
// @Tags Auth
// @Accept json
// @Produce json
// @Param body body input.TempSignUpInput true "SMS auth info"
// @Success 200 {object} output.TempUserOutput
// @Success 400 {object} error_handler.ErrorResponse "Bad request"
// @Success 409 {object} error_handler.ErrorResponse "phoneNumber is already taken"
// @Success 422 {object} error_handler.ErrorResponse "Invalid phone number / region format"
// @Router /auth/temp_sign_up [post]
func (h *authHandler) TempSignUp(ce echo.Context) error {
	ctx := &custom_context.Context{ce}
	output, err := h.usecase.TempSignUp(ctx)
	if err != nil {
		return err
	}
	return ctx.JSON(http.StatusCreated, output)
}

// ConfirmTempUser
// @Summary Confirm temporary SMS Auth
// @Description SMS認証チェック
// @Tags Auth
// @Accept json
// @Produce json
// @Param body body input.ConfirmTempUserInput true "body"
// @Success 200 {object} output.AuthTokenOutput
// @Success 400 {object} error_handler.ErrorResponse "Bad request"
// @Success 409 {object} error_handler.ErrorResponse "Phone number is already registered"
// @Success 422 {object} error_handler.ErrorResponse "Invalid phone number / region format"
// @Router /auth/confirm [post]
func (h *authHandler) ConfirmTempUser(ce echo.Context) error {
	ctx := &custom_context.Context{ce}
	output, err := h.usecase.ConfirmTempUser(ctx)
	if err != nil {
		return err
	}
	return ctx.JSON(http.StatusOK, output)
}

// SignUp godoc
// @Summary Sign Up
// @Description ユーザー登録
// @Tags Auth
// @Accept json
// @Produce json
// @Security JWTAuth
// @Param body body input.SignUpInput true "body"
// @Success 200 {object} output.CurrentUserOutput
// @Success 400 {object} error_handler.ErrorResponse "Bad request"
// @Success 409 {object} error_handler.ErrorResponse "Email is already registered"
// @Success 422 {object} error_handler.ErrorResponse "Invalid email format"
// @Router /auth/sign_up [post]
func (h *authHandler) SignUp(ce echo.Context) error {
	ctx := &custom_context.Context{ce}
	user, err := h.usecase.SignUp(ctx)
	if err != nil {
		return err
	}
	return ctx.JSON(http.StatusCreated, user)
}

// SignIn godoc
// @Summary Sign In
// @Description サインイン
// @Tags Auth
// @Accept json
// @Produce json
// @Param body body input.SignInInput true "body"
// @Success 200 {object} output.CurrentUserOutput
// @Success 401 {object} error_handler.ErrorResponse "Unauthenticated"
// @Success 422 {object} error_handler.ErrorResponse "Invalid email format"
// @Router /auth/sign_in [post]
func (h *authHandler) SignIn(ce echo.Context) error {
	ctx := &custom_context.Context{ce}
	user, err := h.usecase.SignIn(ctx)
	if err != nil {
		return err
	}
	return ctx.JSON(http.StatusOK, user)
}

// UpdateUnconfirmedPhoneNumber godoc
// @Summary Update Current User's Phone Number
// @Description ログインユーザーの電話番号変更
// @Tags Auth
// @Accept json
// @Produce json
// @Security JWTAuth
// @Param body body input.UpdateUnconfirmedPhoneNumberInput true "body"
// @Success 200 {object} output.CurrentUserOutput
// @Success 400 {object} error_handler.ErrorResponse "Bad request"
// @Success 409 {object} error_handler.ErrorResponse "Phone number is already registered"
// @Success 422 {object} error_handler.ErrorResponse "Invalid phone number format"
// @Router /auth/phone_number [patch]
func (h *authHandler) UpdateUnconfirmedPhoneNumber(ce echo.Context) error {
	ctx := &custom_context.Context{ce}
	user, err := h.usecase.UpdateUnconfirmedPhoneNumber(ctx)
	if err != nil {
		return err
	}
	return ctx.JSON(http.StatusOK, user)
}

// ConfirmPhoneNumber godoc
// @Summary Confirm Current User's Phone Number
// @Description ログインユーザーの電話番号を確認 & 更新
// @Tags Auth
// @Accept json
// @Produce json
// @Security JWTAuth
// @Param body body input.ConfirmPhoneNumberInput true "body"
// @Success 200 {object} output.CurrentUserOutput
// @Success 400 {object} error_handler.ErrorResponse "Bad request"
// @Success 409 {object} error_handler.ErrorResponse "Phone number is already taken"
// @Success 422 {object} error_handler.ErrorResponse "Invalid phone number format"
// @Router /auth/phone_number/confirm [post]
func (h *authHandler) ConfirmPhoneNumber(ce echo.Context) error {
	ctx := &custom_context.Context{ce}
	user, err := h.usecase.ConfirmPhoneNumber(ctx)
	if err != nil {
		return err
	}
	return ctx.JSON(http.StatusOK, user)
}
