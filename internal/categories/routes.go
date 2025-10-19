package categories

import "github.com/go-chi/chi/v5"

func RegisterRoutes(r chi.Router, h *CategoryHandler) {
	r.Route("/categories", func(r chi.Router) {
		r.Get("/", h.FindAll)
		r.Get("/name/{name}", h.FindByName)

		r.Route("/{id}", func(r chi.Router) {
			r.Get("/", h.FindByID)
		})
	})
}
