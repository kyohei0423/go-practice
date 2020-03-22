package infra

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/kyohei0423/go-practice/clean-architecture/interface/signup"
)

func Router() {
	r := chi.NewRouter()
	r.Route("/users", func(r chi.Router) {
		r.Post("/signup", signup.Controller)
	})
	http.ListenAndServe(":3000", r)
}
