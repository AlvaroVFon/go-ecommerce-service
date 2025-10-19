// Package seeding
package seeding

import (
	"database/sql"
	"fmt"
)

func SeedRoles(db *sql.DB) {
	query := `INSERT INTO roles (name) VALUES ($1) ON CONFLICT (name) DO NOTHING;`
	roles := []string{"superadmin", "admin", "user"}

	for _, n := range roles {
		_, err := db.Exec(query, n)
		if err != nil {
			fmt.Println("Error seeding role", n, ":", err)
		}
	}
}
