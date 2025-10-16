// Package bootstrap initializes configuration and database connection for the application.
package bootstrap

import (
	"database/sql"
	"ecommerce-service/internal/auth"
	"ecommerce-service/internal/config"
	"ecommerce-service/internal/products"
	"ecommerce-service/internal/tokens"
	"ecommerce-service/internal/users"
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
	// Load environment variables
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

	// Initialize user module
	userRepository := users.NewUserRepository(b.DB)
	userService := users.NewUserService(userRepository, b.Config)
	userHandler := users.NewUserHandler(userService)

	// Initialize token service
	tokenService := tokens.NewTokenService(b.Config)

	// Initialize auth module
	authService := auth.NewAuthService(userService)
	authHandler := auth.NewAuthHandler(authService, tokenService)

	// Initialize product module
	productRepository := products.NewProductRepository(b.DB)
	productService := products.NewProductService(productRepository, b.Config)
	productHandler := products.NewProductHandler(productService)

	// Register routes
	healthcheck.RegisterRoutes(b.Router, healthCheckHandler)
	users.RegisterRoutes(b.Router, userHandler)
	products.RegisterRoutes(b.Router, productHandler)
	auth.RegisterRoutes(b.Router, authHandler)

	return &b, nil
}
