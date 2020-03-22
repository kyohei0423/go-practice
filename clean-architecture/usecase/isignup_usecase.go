package usecase

type ISignupUsecase interface {
	Handle(InputData) OutputData
}
