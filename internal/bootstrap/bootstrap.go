// Package bootstrap initializes configuration and database connection for the application.
package bootstrap

import (
	"database/sql"
	"ecommerce-service/internal/config"
	"ecommerce-service/internal/product"
	"log"

	healthcheck "ecommerce-service/internal/health-check"

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

	// Initialize health check module
	healthCheckHandler := healthcheck.NewHealthCheckHandler()

	// Initialize product module
	productRepository := product.NewProductRepository(b.DB)
	productService := product.NewProductService(productRepository)
	productHandler := product.NewProductHandler(productService)

	// Register routes
	healthcheck.RegisterRoutes(b.Router, healthCheckHandler)
	product.RegisterRoutes(b.Router, productHandler)

	return &b, nil
}
