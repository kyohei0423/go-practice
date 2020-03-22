package usecase

import "github.com/kyohei0423/go-practice/clean-architecture/domain"

type IUserRepository interface {
	Create(domain.ISignup) (domain.ISignup, error)
}
