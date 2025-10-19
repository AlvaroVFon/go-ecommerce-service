package main

import (
	"ecommerce-service/internal/config"
	"ecommerce-service/internal/database/seeding"
)

func main() {
	db, err := config.ConnectDatabase()
	if err != nil {
		return
	}
	seeding.SeedRoles(db)
}
