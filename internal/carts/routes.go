package carts

import "github.com/go-chi/chi/v5"

func RegisterRoutes(r chi.Router, h *CartHandler) {
	r.Route("/carts", func(r chi.Router) {
		r.Get("/{id}", h.GetCart)
		r.Post("/{id}/items", h.AddItemToCart)
		r.Delete("/{id}/clear", h.ClearCart)
		r.Post("/{id}/complete", h.CompleteCart)
	})
}
