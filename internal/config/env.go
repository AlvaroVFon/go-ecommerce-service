package config

import (
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type Config struct {
	// app
	AppName string
	AppEnv  string
	AppHost string
	AppPort string

	// database
	DBHost     string
	DBPort     int
	DBUser     string
	DBPassword string
	DBName     string
	DBSSLMode  string

	// pagination
	Limit    int
	MaxLimit int
	Offset   int

	// Cryptox
	BcryptCost int

	// JWT
	JWTSecret        string
	JWTExp           int // in seconds
	JWTRefreshSecret string
	JWTRefreshExp    int // in seconds
}

func LoadEnvVars() *Config {
	err := godotenv.Load()
	if err != nil {
		log.Println("⚠️  No se pudo cargar .env, usando variables del sistema")
	}

	// Leer y parsear variables
	dbPort, err := strconv.Atoi(getEnv("DB_PORT", "5433"))
	if err != nil {
		log.Fatalf("DB_PORT debe ser un número válido: %v", err)
	}

	// pagination
	limit, err := strconv.Atoi(getEnv("PAGINATION_LIMIT", "10"))
	if err != nil {
		log.Fatalf("PAGINATION_LIMIT debe ser un número válido: %v", err)
	}

	offset, err := strconv.Atoi(getEnv("PAGINATION_OFFSET", "0"))
	if err != nil {
		log.Fatalf("PAGINATION_OFFSET debe ser un número válido: %v", err)
	}

	// cryptox
	bcryptCost, err := strconv.Atoi(getEnv("BCRYPT_COST", "10"))
	if err != nil {
		log.Fatalf("BCRYPT_COST debe ser un número válido: %v", err)
	}

	JWTExp, err := strconv.Atoi(getEnv("JWT_EXP", "3600"))
	if err != nil {
		log.Fatalf("JWT_EXP debe ser un número válido: %v", err)
	}

	JWTRefreshExp, err := strconv.Atoi(getEnv("JWT_REFRESH_EXP", "86400"))
	if err != nil {
		log.Fatalf("JWT_REFRESH_EXP debe ser un número válido: %v", err)
	}

	MaxLimit, err := strconv.Atoi(getEnv("PAGINATION_MAX_LIMIT", "30"))
	if err != nil {
		log.Fatalf("PAGINATION_MAX_LIMIT debe ser un número válido: %v", err)
	}

	cfg := &Config{
		AppName: os.Getenv("APP_NAME"),
		AppEnv:  getEnv("APP_ENV", "development"),
		AppHost: getEnv("APP_HOST", "localhost"),
		AppPort: getEnv("APP_PORT", "8080"),

		DBHost:     getEnv("DB_HOST", "localhost"),
		DBPort:     dbPort,
		DBUser:     getEnv("DB_USER", "postgres"),
		DBPassword: getEnv("DB_PASSWORD", "postgres"),
		DBName:     getEnv("DB_NAME", "ecommerce_db"),
		DBSSLMode:  getEnv("DB_SSLMODE", "disable"),

		Limit:    limit,
		MaxLimit: MaxLimit,
		Offset:   offset,

		BcryptCost: bcryptCost,

		JWTSecret:        getEnv("JWT_SECRET", "your-secret-key"),
		JWTExp:           JWTExp,
		JWTRefreshSecret: getEnv("JWT_REFRESH_SECRET", "your-refresh-secret-key"),
		JWTRefreshExp:    JWTRefreshExp,
	}

	return cfg
}

// getEnv lee la variable de entorno o devuelve un valor por defecto
func getEnv(key, defaultValue string) string {
	if val := os.Getenv(key); val != "" {
		return val
	}
	return defaultValue
}
