// Package strategies implements various authentication strategies.
package strategies

import (
	"context"
	"ecommerce-service/internal/users"
	"ecommerce-service/pkg/cryptox"
	"errors"
)

type UserService interface {
	FindByEmail(ctx context.Context, email string) (*users.User, error)
}

type PasswordCredentials struct {
	Email    string
	Password string
}

type PasswordStrategy struct {
	userService UserService
}

func NewPasswordStrategy(u UserService) *PasswordStrategy {
	return &PasswordStrategy{userService: u}
}

func (ps *PasswordStrategy) Authenticate(ctx context.Context, credentials any) (*users.User, error) {
	creds, ok := credentials.(PasswordCredentials)
	if !ok {
		return nil, errors.New("invalid credentials type for password strategy")
	}

	u, err := ps.userService.FindByEmail(ctx, creds.Email)
	if err != nil {
		return nil, err
	}

	if err := cryptox.VerifyPassword(u.Password, creds.Password); err != nil {
		return nil, err
	}

	return u, nil
}
