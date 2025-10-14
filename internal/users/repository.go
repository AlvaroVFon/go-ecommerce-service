package users

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"strings"
	"time"
)

type UserRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (r *UserRepository) Create(ctx context.Context, u *CreateUserRequest) error {
	query := "INSERT INTO users (email, password) VALUES ($1, $2)"
	_, err := r.db.ExecContext(ctx, query, u.Email, u.Password)
	if err != nil {
		return err
	}
	return nil
}

func (r *UserRepository) FindByID(ctx context.Context, id int) (*PublicUser, error) {
	query := "SELECT id, email,  created_at, updated_at FROM users WHERE id = $1"
	row := r.db.QueryRowContext(ctx, query, id)
	var u PublicUser
	err := row.Scan(&u.ID, &u.Email, &u.CreatedAt, &u.UpdatedAt)
	if err != nil {
		return nil, err
	}
	return &u, nil
}

func (r *UserRepository) FindAll(ctx context.Context) ([]PublicUser, error) {
	query := "SELECT id, email, created_at, updated_at FROM users"
	rows, err := r.db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}

	defer func() {
		if err := rows.Close(); err != nil {
			log.Printf("error closing rows: %v\n", err)
		}
	}()

	var users []PublicUser

	for rows.Next() {
		var u PublicUser
		if err := rows.Scan(&u.ID, &u.Email, &u.CreatedAt, &u.UpdatedAt); err != nil {
			log.Printf("error scanning user: %v\n", err)
			return nil, err
		}
		users = append(users, u)
	}

	return users, nil
}

func (r *UserRepository) Update(ctx context.Context, id int, u UpdateUserRequest) error {
	fields := []string{}
	args := []any{}
	i := 1

	if u.Email != nil {
		fields = append(fields, fmt.Sprintf("email = $%d", i))
		args = append(args, u.Email)
		i++
	}
	if u.Password != nil {
		fields = append(fields, fmt.Sprintf("password = $%d", i))
		args = append(args, u.Password)
		i++
	}

	// Siempre actualizamos updated_at
	now := time.Now()
	fields = append(fields, fmt.Sprintf("updated_at = $%d", i))
	args = append(args, now)
	i++

	if len(fields) == 0 {
		return fmt.Errorf("no fields to update")
	}

	// Query final
	query := fmt.Sprintf("UPDATE users SET %s WHERE id = $%d",
		strings.Join(fields, ", "),
		i, // placeholder para id
	)

	// Ejecutar
	args = append(args, id)
	res, err := r.db.ExecContext(ctx, query, args...)
	if err != nil {
		return err
	}

	rows, err := res.RowsAffected()
	if err != nil {
		return err
	}

	if rows == 0 {
		return fmt.Errorf("no user updated with id %d", id)
	}

	return nil
}
