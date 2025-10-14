package users

import "github.com/go-chi/chi/v5"

func RegisterRoutes(r chi.Router, uh *UserHandler) {
	r.Route("/users", func(r chi.Router) {
		r.Post("/", uh.Create)
		r.Get("/", uh.FindAll)
		r.Get("/{id}", uh.FindByID)
		r.Patch("/{id}", uh.Update)
	})
}
