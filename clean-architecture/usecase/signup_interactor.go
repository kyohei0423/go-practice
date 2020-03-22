package usecase

import "github.com/kyohei0423/go-practice/clean-architecture/domain"

type SignupInteractor struct {
	repository IUserRepository
}

func NewSignupInteractor(r IUserRepository) ISignupUsecase {
	return &SignupInteractor{
		repository: r,
	}
}

func (s *SignupInteractor) Handle(i InputData) OutputData {
	signup := domain.NewSignup(i.Email, i.Password, i.PasswordConfirmation)
	user, err := s.repository.Create(signup)
	if err != nil {

	}
	return OutputData{Email: user.Email()}
}
