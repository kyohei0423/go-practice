//+build wireinject

package signup

import (
	"github.com/google/wire"
	"github.com/kyohei0423/go-practice/clean-architecture/interface/repository"
	"github.com/kyohei0423/go-practice/clean-architecture/usecase"
)

func InitializeInteractor() usecase.ISignupUsecase {
	wire.Build(repository.NewUser, usecase.NewSignupInteractor)

	return &usecase.SignupInteractor{}
}
