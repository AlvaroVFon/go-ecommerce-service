// Package bootstrap initializes configuration and database connection for the application.
package bootstrap

import (
	"database/sql"
	"log"

	"ecommerce-service/internal/auth"
	"ecommerce-service/internal/auth/strategies"
	"ecommerce-service/internal/carts"
	"ecommerce-service/internal/categories"
	"ecommerce-service/internal/config"
	"ecommerce-service/internal/orders"
	"ecommerce-service/internal/products"
	"ecommerce-service/internal/roles"
	"ecommerce-service/internal/tokens"
	"ecommerce-service/internal/users"

	healthcheck "ecommerce-service/internal/health-check"

	"github.com/go-chi/chi/v5"
	"github.com/go-playground/validator/v10"
)

type Bootstrapper struct {
	Config *config.Config
	DB     *sql.DB
	Router chi.Router
}

func Bootstrap() (*Bootstrapper, error) {
	// Load environment variables
	c := config.LoadEnvVars()

	// Connect to the database
	db, err := config.ConnectDatabase(c)
	if err != nil {
		log.Println("Error connecting to the database:", err)
		return nil, err
	}

	// Initialize the router
	r := chi.NewRouter()

	// Create the bootstrapper instance
	b := Bootstrapper{
		Config: c,
		DB:     db,
		Router: r,
	}

	// validator
	validate := validator.New()

	// Initialize modules

	// health-check module
	healthCheckHandler := healthcheck.NewHealthCheckHandler()

	// roles module
	roleRepository := roles.NewRoleRepository(db)
	roleService := roles.NewRoleService(roleRepository)
	roleHandler := roles.NewRoleHandler(roleService)

	// user module
	userRepository := users.NewUserRepository(b.DB)
	userService := users.NewUserService(userRepository, b.Config)
	userHandler := users.NewUserHandler(userService, validate, b.Config)

	// token module
	tokenService := tokens.NewTokenService(b.Config)

	// strategies
	passwordStrategy := strategies.NewPasswordStrategy(userService)

	// strategies registry
	authStrategies := map[string]auth.AuthStrategy{
		"password": passwordStrategy,
	}

	// auth module
	authService := auth.NewAuthService(authStrategies)
	authHandler := auth.NewAuthHandler(authService, tokenService)

	// Initialize product module
	productRepository := products.NewProductRepository(b.DB)
	productService := products.NewProductService(productRepository, b.Config)
	productHandler := products.NewProductHandler(productService, validate, b.Config)

	// category module
	categoryRepository := categories.NewCategoryRepository(b.DB)
	categoryService := categories.NewCategoryService(categoryRepository)
	categoryHandler := categories.NewCategoryHandler(categoryService, b.Config)

	// cart module
	cartRepository := carts.NewCartRepository(b.DB)
	cartService := carts.NewCartService(cartRepository)
	cartHandler := carts.NewCartHandler(cartService, validate)

	// orders module
	orderRepository := orders.NewOrderRepository(b.DB)
	orderService := orders.NewOrderService(orderRepository, cartRepository)
	orderHandler := orders.NewOrderHandler(orderService, validate, b.Config)

	// Register routes
	healthcheck.RegisterRoutes(b.Router, healthCheckHandler)
	roles.RegisterRoutes(b.Router, roleHandler)
	users.RegisterRoutes(b.Router, userHandler)
	products.RegisterRoutes(b.Router, productHandler)
	auth.RegisterRoutes(b.Router, authHandler)
	categories.RegisterRoutes(b.Router, categoryHandler)
	carts.RegisterRoutes(b.Router, cartHandler)
	orders.RegisterRoutes(b.Router, orderHandler)

	return &b, nil
}
