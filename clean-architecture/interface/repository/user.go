package repository

import (
	"github.com/kyohei0423/go-practice/clean-architecture/domain"
	"github.com/kyohei0423/go-practice/clean-architecture/usecase"
)

type user struct {
	db []domain.ISignup
}

// NewUser is constutor
func NewUser() usecase.IUserRepository {
	return &user{
		db: []domain.ISignup{},
	}
}

func (u *user) Create(s domain.ISignup) (domain.ISignup, error) {
	u.db = append(u.db, s)
	return s, nil
}
