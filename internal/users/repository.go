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
	query := "INSERT INTO users (email, password, role_id) VALUES ($1, $2)"
	_, err := r.db.ExecContext(ctx, query, u.Email, u.Password, u.RoleID)
	if err != nil {
		return err
	}
	return nil
}

func (r *UserRepository) FindByID(ctx context.Context, id int) (*User, error) {
	query := "SELECT id, email, password, role_id, created_at, updated_at FROM users WHERE id = $1"
	row := r.db.QueryRowContext(ctx, query, id)
	var u User
	err := row.Scan(&u.ID, &u.Email, &u.Password, &u.RoleID, &u.CreatedAt, &u.UpdatedAt)
	if err != nil {
		return nil, err
	}
	return &u, nil
}

func (r *UserRepository) FindByEmail(ctx context.Context, email string) (*User, error) {
	query := "SELECT id, email, password, role_id, created_at, updated_at FROM users WHERE email = $1"
	row := r.db.QueryRowContext(ctx, query, email)
	var u User
	err := row.Scan(&u.ID, &u.Email, &u.Password, &u.RoleID, &u.CreatedAt, &u.UpdatedAt)
	if err != nil {
		return nil, err
	}

	return &u, nil
}

func (r *UserRepository) FindAll(ctx context.Context, limit, offset int) ([]User, error) {
	query := "SELECT id, email, password, role_id, created_at, updated_at FROM users LIMIT $1 OFFSET $2"
	rows, err := r.db.QueryContext(ctx, query, limit, offset)
	if err != nil {
		return nil, err
	}

	defer func() {
		if err := rows.Close(); err != nil {
			log.Printf("error closing rows: %v\n", err)
		}
	}()

	var users []User

	for rows.Next() {
		var u User
		if err := rows.Scan(&u.ID, &u.Email, &u.Password, &u.RoleID, &u.CreatedAt, &u.UpdatedAt); err != nil {
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

	if u.RoleID != nil {
		fields = append(fields, fmt.Sprintf("role_id = $%d", i))
		args = append(args, u.RoleID)
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

func (r *UserRepository) Delete(ctx context.Context, id int) error {
	query := "DELETE FROM users WHERE id = $1"

	result, err := r.db.ExecContext(ctx, query, id)
	if err != nil {
		log.Printf("error deleting user with id %d: %v\n", id, err)
		return err
	}

	if rows, err := result.RowsAffected(); err != nil || rows == 0 {
		return err
	}

	return nil
}

func (r *UserRepository) Count(ctx context.Context) (int, error) {
	query := "SELECT COUNT(*) FROM users"
	row := r.db.QueryRowContext(ctx, query)
	var count int

	if err := row.Scan(&count); err != nil {
		return 0, fmt.Errorf("error scanning user count: %v", err)
	}
	return count, nil
}
