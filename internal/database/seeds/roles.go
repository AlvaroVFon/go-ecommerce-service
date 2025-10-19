// Package seeds
package seeds

import (
	"database/sql"
)

func SeedRoles(db *sql.DB) error {
	query := `INSERT INTO roles (name) VALUES ($1) ON CONFLICT (name) DO NOTHING;`
	roles := []string{"superadmin", "admin", "user"}

	for _, n := range roles {
		_, err := db.Exec(query, n)
		if err != nil {
			return err
		}
	}
	return nil
}
