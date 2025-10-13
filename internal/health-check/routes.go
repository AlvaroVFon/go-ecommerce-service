package healthcheck

import "github.com/go-chi/chi/v5"

func RegisterRoutes(r chi.Router, h *HealthCheckHandler) {
	r.Get("/health", h.Check)
}
