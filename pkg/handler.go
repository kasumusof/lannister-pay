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

	r.Post("/split-payments/computes", split)

	return r
}

func split(w http.ResponseWriter, r *http.Request) {

	var p Payload
	if err := render.Bind(r, &p); err != nil {
		renderResponse(w, err, http.StatusBadRequest)
		return
	}

	renderResponse(w,
		response{
			ID:      p.ID,
			Balance: p.Amount,
			SplitBreakdown: []splitBreakdown{
				{
					SplitEntityID: "1",
					Amount:        p.Amount / 2,
				},
			},
		},
		http.StatusOK)
}
