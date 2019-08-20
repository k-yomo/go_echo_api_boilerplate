package main

import (
	"github.com/k-yomo/go_echo_api_boilerplate/config"
	_ "github.com/k-yomo/go_echo_api_boilerplate/docs"
	"github.com/k-yomo/go_echo_api_boilerplate/handler"
	"github.com/k-yomo/go_echo_api_boilerplate/internal/error_handler"
	"github.com/k-yomo/go_echo_api_boilerplate/middleware/jwt_middleware"
	"github.com/k-yomo/go_echo_api_boilerplate/middleware/logger_middleware"
	"github.com/k-yomo/go_echo_api_boilerplate/pkg/logger"
	"github.com/k-yomo/go_echo_api_boilerplate/pkg/params_validator"
	"github.com/k-yomo/go_echo_api_boilerplate/pkg/sms"
	"github.com/k-yomo/go_echo_api_boilerplate/repository"
	"github.com/k-yomo/go_echo_api_boilerplate/usecase"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"go.uber.org/zap"
)

// @title Go API Boilerplate
// @version 0.0.1
// @description API server

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:1323
// @BasePath /v1

// @securityDefinitions.apikey JWTAuth
// @in header
// @name Authorization
func main() {
	e := echo.New()

	db, err := config.NewDB()
	if err != nil {
		panic(err)
	}
	defer db.Close()

	zapLogger, _ := zap.NewProduction()
	e.Logger = logger.NewLogger(zapLogger.Named("system_log"))
	e.Validator = params_validator.NewValidator()
	e.HTTPErrorHandler = error_handler.HTTPErrorHandler
	e.Use(logger_middleware.LoggerMiddleware(zapLogger.Named("access_log")))
	e.Use(middleware.Recover())

	repo := repository.NewRepository(db)
	uc := usecase.NewUsecase(repo, sms.NewSMSMessenger())
	handler.NewHandler(e, uc, jwt_middleware.NewJWTMiddleware())
	e.Logger.Fatal("", e.Start(":1323"))
}
