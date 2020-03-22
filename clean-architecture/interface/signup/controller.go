package signup

import (
	"encoding/json"
	"net/http"

	"github.com/kyohei0423/go-practice/clean-architecture/usecase"
)

func Controller(w http.ResponseWriter, r *http.Request) {
	input := &usecase.InputData{}
	if err := json.NewDecoder(r.Body).Decode(input); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
	interactor := InitializeInteractor()
	output := interactor.Handle(*input)
	res, err := Output(output)
	if err != nil {
	}
	w.Write(res)
}
