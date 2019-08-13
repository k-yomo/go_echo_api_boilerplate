package usecase

import (
	"github.com/k-yomo/go_echo_boilerplate/internal/custom_context"
	"github.com/k-yomo/go_echo_boilerplate/repository"
)

type HealthcheckUsecase interface {
	CheckReadiness(ctx *custom_context.Context) error
}

type healthcheckUsecase struct {
	HealthcheckRepo repository.HealthcheckRepository
}

// NewHealthcheckUsecase returns healthcheckUsecase
func NewHealthcheckUsecase(healthcheckRepo repository.HealthcheckRepository) *healthcheckUsecase {
	return &healthcheckUsecase{HealthcheckRepo: healthcheckRepo}
}

// CheckReadiness checks db connection
func (hu *healthcheckUsecase) CheckReadiness(ctx *custom_context.Context) error {
	err := hu.HealthcheckRepo.PingDB()
	return err
}
