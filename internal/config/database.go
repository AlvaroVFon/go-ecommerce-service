// Package config handles database connection setup and configuration loading.
package config

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

func ConnectDatabase() (*sql.DB, error) {
	connectionString := generateConnectionString(LoadEnvVars())
	db, err := sql.Open("postgres", connectionString)
	if err != nil {
		log.Println("Error connecting to the database:", err)
		return nil, err
	}

	if err := db.Ping(); err != nil {
		log.Println("Database is unreachable:", err)
		return nil, err
	}
	return db, nil
}

func generateConnectionString(cfg *Config) string {
	return fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=%s connect_timeout=10",
		cfg.DBHost, cfg.DBPort, cfg.DBUser, cfg.DBPassword, cfg.DBName, cfg.DBSSLMode)
}
