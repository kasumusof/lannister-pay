package pkg

import (
	"net/http"

	"github.com/go-chi/render"

	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/chi/v5"
)

func Router() *chi.Mux {
	r := chi.NewRouter()
	r.Use(middleware.Logger)

	r.Get("/", hello)
	r.Post("/split-payments/computes", split)

	return r
}

func hello(w http.ResponseWriter, r *http.Request) {
	renderResponse(w,
		map[string]string{
			"status":  "ok",
			"message": "Welcome to Lannister Pay!",
		},
		http.StatusOK,
	)
}

func split(w http.ResponseWriter, r *http.Request) {

	var p Payload
	if err := render.Bind(r, &p); err != nil {
		renderResponse(w, err, http.StatusBadRequest)
		return
	}

	var breakDown []splitBreakdown
	compute(&p, &breakDown)

	renderResponse(w,
		response{
			ID:             p.ID,
			Balance:        p.Amount,
			SplitBreakdown: breakDown,
		},
		http.StatusOK)
}
