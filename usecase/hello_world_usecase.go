package usecase

import (
	"github.com/k-yomo/go_echo_boilerplate/internal/custom_context"
)

type HelloWorldUsecase interface {
	Greet(ctx *custom_context.Context) (*helloWorld, error)
}

type helloWorldUsecase struct{}

type helloWorld struct {
	Name string `validate:"required,max=20"`
}

// NewHelloWorldUsecase returns helloWorldUsecase
func NewHelloWorldUsecase() *helloWorldUsecase {
	return new(helloWorldUsecase)
}

func (*helloWorldUsecase) Greet(ctx *custom_context.Context) (*helloWorld, error) {
	hw := new(helloWorld)
	if err := ctx.BindWithValidation(hw); err != nil {
		return nil, err
	}
	return hw, nil
}
