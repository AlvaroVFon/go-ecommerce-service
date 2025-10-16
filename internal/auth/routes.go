package auth

import "github.com/go-chi/chi/v5"

func RegisterRoutes(r chi.Router, ah *AuthHandler) {
	r.Route("/auth", func(r chi.Router) {
		r.Post("/login", ah.Login)
	})
}
