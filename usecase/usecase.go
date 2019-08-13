package usecase

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/k-yomo/go_echo_boilerplate/internal/custom_context"
	"github.com/k-yomo/go_echo_boilerplate/pkg/sms"
	"github.com/k-yomo/go_echo_boilerplate/repository"
)

// Usecase manage all usecases
type Usecase struct {
	HealthCheckUsecase HealthcheckUsecase
	AuthUsecase        AuthUsecase
	UserUsecase        UserUsecase
}

// NewUsecase returns Usecase
func NewUsecase(repo *repository.Repository, smsMessenger sms.SMSMessenger) *Usecase {
	return &Usecase{
		HealthCheckUsecase: NewHealthcheckUsecase(repo.HealthcheckRepository),
		AuthUsecase: NewAuthUsecase(repo.TempUserRepository, repo.UserRepository, smsMessenger),
		UserUsecase: NewUserUsecase(repo.UserRepository),
	}
}

func getUserIdFromContext(ctx *custom_context.Context) uint64 {
	user := ctx.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	userId := claims["id"].(float64)
	return uint64(userId)
}
