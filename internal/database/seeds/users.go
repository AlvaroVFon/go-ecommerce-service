package seeds

import (
	"database/sql"
	"ecommerce-service/internal/users"
	"ecommerce-service/pkg/cryptox"
)

func SeedUsers(db *sql.DB) error {
	query := `INSERT INTO users (email, password, role_id) VALUES ($1, $2, $3) ON CONFLICT (email) DO NOTHING;`
	users := []users.User{
		{Email: "superadmin@email.com", Password: "superadmin123", RoleID: 1}, // password: superadmin123
		{Email: "admin@email.com", Password: "admin123", RoleID: 2},           // password: admin12345
		{Email: "user@email.com", Password: "user123", RoleID: 3},             // password: user123456
	}

	for _, u := range users {
		pass, err := cryptox.HashPassword(u.Password, 10)
		if err != nil {
			return err
		}

		if _, err := db.Exec(query, u.Email, pass, u.RoleID); err != nil {
			return err
		}
	}
	return nil
}
