package roles

import "github.com/go-chi/chi/v5"

func RegisterRoutes(r chi.Router, h *RoleHandler) {
	r.Route("/roles", func(r chi.Router) {
		r.Get("/", h.FindAll)
		r.Get("/{id}", h.FindByID)
	})
}
