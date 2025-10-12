package product

import (
	"database/sql"

	"ecommerce-service/internal/config"

	"github.com/go-chi/chi/v5"
)

func Wire(r chi.Router, db *sql.DB, c *config.Config) {
	ph := NewProductHandler(NewProductService(NewProductRepository(db)))
	r.Post("/products", ph.Create)
}
