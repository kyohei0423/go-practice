package signup

import (
	"encoding/json"

	"github.com/kyohei0423/go-practice/clean-architecture/usecase"
)

func Output(o usecase.OutputData) ([]byte, error) {
	return json.Marshal(o)
}
