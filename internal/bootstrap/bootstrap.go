// Package bootstrap initializes configuration and database connection for the application.
package bootstrap

import (
	"database/sql"
	"log"

	"ecommerce-service/internal/config"
	healthcheck "ecommerce-service/internal/health-check"
	"ecommerce-service/internal/product"

	"github.com/go-chi/chi/v5"
)

type Bootstrapper struct {
	Config *config.Config
	DB     *sql.DB
	Router chi.Router
}

func Bootstrap() (*Bootstrapper, error) {
	// Load environment variables and connect to the database
	envVars := config.LoadEnvVars()

	// Connect to the database
	db, err := config.ConnectDatabase()
	if err != nil {
		log.Println("Error connecting to the database:", err)
		return nil, err
	}

	// Initialize the router
	r := chi.NewRouter()

	// Create the bootstrapper instance
	b := Bootstrapper{
		Config: envVars,
		DB:     db,
		Router: r,
	}

	// Initialize modules
	healthcheck.Wire(b.Router, b.DB, b.Config)

	// Initialize product module
	productRepository := product.NewProductRepository(b.DB)
	productService := product.NewProductService(productRepository)
	productHandler := product.NewProductHandler(productService)
	product.RegisterRoutes(b.Router, productHandler)

	return &b, nil
}
