package domain

type signup struct {
	email                string
	password             string
	passwordConfirmation string
}

type ISignup interface {
	Email() string
	Password() string
	PasswordConfirmation() string
}

func NewSignup(email, password, passwordConfirmation string) ISignup {
	return &signup{
		email: email,
		password: password,
		passwordConfirmation: passwordConfirmation,
	}
}

func (s *signup) Email() string {
	return s.email
}

func (s *signup) Password() string {
	return s.password
}

func (s *signup) PasswordConfirmation() string {
	return s.passwordConfirmation
}
