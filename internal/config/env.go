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
		log.Println("⚠️ No se pudo cargar .env, usando variables del sistema")
	}

	// Leer y parsear variables
	dbPort, err := getIntEnv("DB_PORT", 5432)
	if err != nil {
		log.Printf("⚠️ Error al leer DB_PORT: %v", err)
	}

	// pagination
	limit, err := getIntEnv("PAGINATION_LIMIT", 10)
	if err != nil {
		log.Printf("⚠️ Error al leer PAGINATION_LIMIT: %v", err)
	}
	MaxLimit, err := getIntEnv("PAGINATION_MAX_LIMIT", 100)
	if err != nil {
		log.Printf("⚠️ Error al leer PAGINATION_MAX_LIMIT: %v", err)
	}

	offset, err := getIntEnv("PAGINATION_OFFSET", 0)
	if err != nil {
		log.Printf("⚠️ Error al leer PAGINATION_OFFSET: %v", err)
	}

	// cryptox
	bcryptCost, err := getIntEnv("BCRYPT_COST", 10)
	if err != nil {
		log.Printf("⚠️ Error al leer BCRYPT_COST: %v", err)
	}

	// jwt
	JWTExp, err := getIntEnv("JWT_EXP", 3600)
	if err != nil {
		log.Printf("⚠️ Error al leer JWT_EXP: %v", err)
	}
	JWTRefreshExp, err := getIntEnv("JWT_REFRESH_EXP", 86400)
	if err != nil {
		log.Printf("⚠️ Error al leer JWT_REFRESH_EXP: %v", err)
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

func getIntEnv(key string, defaultValue int) (int, error) {
	if val := os.Getenv(key); val != "" {
		if intValue, err := strconv.Atoi(val); err != nil {
			return 0, err
		} else {
			return intValue, nil
		}
	}
	return defaultValue, nil
}
