package healthcheck

import (
	"database/sql"

	"ecommerce-service/internal/config"

	"github.com/go-chi/chi/v5"
)

func Wire(r chi.Router, db *sql.DB, c *config.Config) chi.Router {
	hc := NewHealthCheckHandler()
	r.Get("/health", hc.Check)
	return r
}
