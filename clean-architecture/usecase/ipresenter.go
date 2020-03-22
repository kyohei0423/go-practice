package usecase

type IPresenter interface {
	Output(OutputData) ([]byte, error)
}
