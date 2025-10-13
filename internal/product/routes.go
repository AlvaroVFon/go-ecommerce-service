package product

import "github.com/go-chi/chi/v5"

func RegisterRoutes(r chi.Router, ph *ProductHandler) {
	r.Route("/products", func(r chi.Router) {
		r.Get("/", ph.FindAll)
		r.Post("/", ph.Create)

		r.Route("/{id}", func(r chi.Router) {
			r.Get("/", ph.FindByID)
			r.Put("/", ph.Update)
			r.Delete("/", ph.Delete)
		})
	})
}
