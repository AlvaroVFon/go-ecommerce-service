package main

import (
	"database/sql"
	"fmt"

	"ecommerce-service/internal/config"
	"ecommerce-service/internal/database/seeds"
)

func main() {
	db, err := config.ConnectDatabase()
	if err != nil {
		return
	}
	seeders := []func(db *sql.DB) error{
		seeds.SeedRoles,
		seeds.SeedUsers,
		seeds.SeedCategories,
		seeds.SeedProducts,
		seeds.SeedProductCategory,
	}

	for _, seeder := range seeders {
		err := seeder(db)
		if err != nil {
			fmt.Println("Seeding error:", err)
		}
	}
}
