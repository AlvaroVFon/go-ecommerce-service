// Package mocks contains mock implementations for testing purposes.
package mocks

import (
	"context"
	"database/sql"
	"time"

	"ecommerce-service/internal/config"
	"ecommerce-service/internal/users"
	"ecommerce-service/pkg/cryptox"
)

type UserMockRepository struct {
	db     []users.User
	config *config.Config
}

func NewUserMockRepository(c *config.Config) *UserMockRepository {
	return &UserMockRepository{
		db:     []users.User{},
		config: c,
	}
}

func (r *UserMockRepository) Create(ctx context.Context, u *users.CreateUserRequest) error {
	var pass string
	if u.Password != "" {
		hashPass, err := cryptox.HashPassword(u.Password, 10)
		if err != nil {
			return err
		}
		pass = hashPass
	}
	user := users.User{
		ID:        len(r.db),
		Email:     u.Email,
		Password:  pass,
		RoleID:    u.RoleID,
		CreatedAt: &time.Time{},
		UpdatedAt: &time.Time{},
	}

	r.db = append(r.db, user)

	return nil
}

func (r *UserMockRepository) FindByID(ctx context.Context, id int) (*users.User, error) {
	for _, u := range r.db {
		if u.ID == id {
			return &u, nil
		}
	}
	return nil, sql.ErrNoRows
}

func (r *UserMockRepository) FindByEmail(ctx context.Context, email string) (*users.User, error) {
	for _, u := range r.db {
		if u.Email == email {
			return &u, nil
		}
	}
	return nil, sql.ErrNoRows
}

func (r *UserMockRepository) FindAll(ctx context.Context, limit, offset int) ([]users.User, error) {
	if offset > len(r.db) {
		return []users.User{}, nil
	}

	end := offset + limit
	end = min(end, len(r.db))

	return r.db[offset:end], nil
}

// TODO: implement dynamic update (patch)
func (r *UserMockRepository) Update(ctx context.Context, id int, u users.UpdateUserRequest) {}

func (r *UserMockRepository) Delete(ctx context.Context, id int) error {
	for i, u := range r.db {
		if u.ID == id {
			r.db = append(r.db[:i], r.db[i+1:]...)
		}
		return nil
	}
	return sql.ErrNoRows
}

func (r *UserMockRepository) Count(ctx context.Context) (int, error) {
	return len(r.db), nil
}
