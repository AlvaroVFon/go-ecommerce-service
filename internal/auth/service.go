// Package auth implements authentication and authorization services.
package auth

import (
	"context"

	"ecommerce-service/internal/users"
)

type UserService interface {
	FindByID(ctx context.Context, id int) (*users.User, error)
}

type AuthService struct {
	userService UserService
}
