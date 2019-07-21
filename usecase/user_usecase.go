package usecase

import (
	"context"
	"github.com/k-yomo/go_echo_boilerplate/internal/custom_context"
	"github.com/k-yomo/go_echo_boilerplate/model"
	"github.com/k-yomo/go_echo_boilerplate/repository"
	"github.com/k-yomo/go_echo_boilerplate/usecase/input"
	"github.com/k-yomo/go_echo_boilerplate/usecase/output"
)

// UserUsecase is an interface for User related usecase
type UserUsecase interface {
	GetProfile(ctx *custom_context.Context) (*output.CurrentUserOutput, error)
	UpdateProfile(ctx *custom_context.Context) (*output.CurrentUserOutput, error)
	UpdateEmail(ctx *custom_context.Context) (*output.CurrentUserOutput, error)
}

type userUsecase struct {
	UserRepo repository.UserRepository
}

// NewUserUsecase returns userUsecase
func NewUserUsecase(userRepo repository.UserRepository) *userUsecase {
	return &userUsecase{UserRepo: userRepo}
}

// GetProfile returns current user profile
func (uu *userUsecase) GetProfile(ctx *custom_context.Context) (*output.CurrentUserOutput, error) {
	userId := getUserIdFromContext(ctx)
	u, err := uu.UserRepo.FindByID(context.Background(), userId)
	if err != nil {
		return nil, err
	} else if u == nil {
		return nil, newNotFoundError("User", "id", userId)
	}

	return output.NewCurrentUserOutput(u), nil
}

// UpdateProfile updates current user profile
func (uu *userUsecase) UpdateProfile(ctx *custom_context.Context) (*output.CurrentUserOutput, error) {
	userId := getUserIdFromContext(ctx)

	params := new(input.UpdateProfileInput)
	if err := ctx.BindWithValidation(params); err != nil {
		return nil, err
	}

	user, err := uu.UserRepo.FindByID(context.Background(), uint64(userId))
	if err != nil {
		return nil, err
	} else if user == nil {
		return nil, newNotFoundError("User", "id", userId)
	}

	err = uu.UserRepo.UpdateProfile(context.Background(), user, params.FirstName, params.LastName, model.StringToGender(params.Gender), params.DateOfBirth.ToNullTime())
	if err != nil {
		return nil, err
	}
	return output.NewCurrentUserOutput(user), nil
}

// UpdateEmail updates current user email
func (uu *userUsecase) UpdateEmail(ctx *custom_context.Context) (*output.CurrentUserOutput, error) {
	userId := getUserIdFromContext(ctx)

	params := new(input.UpdateEmailInput)
	if err := ctx.BindWithValidation(params); err != nil {
		return nil, err
	}

	user, err := uu.UserRepo.FindByID(context.Background(), uint64(userId))
	if err != nil {
		return nil, err
	} else if user == nil {
		return nil, newNotFoundError("User", "id", userId)
	}

	if u, err := uu.UserRepo.FindByEmail(context.Background(), params.Email); err != nil {
		return nil, err
	} else if u != nil {
		return nil, newAlreadyTakenError("email", params.Email)
	}

	user.Email = params.Email
	err = uu.UserRepo.Update(context.Background(), user)
	if err != nil {
		return nil, err
	}
	return output.NewCurrentUserOutput(user), nil
}
