// Package auth implements authentication and authorization services.
package auth

import (
	"context"
	"ecommerce-service/internal/users"
	"ecommerce-service/pkg/cryptox"
)

type UserService interface {
	FindByEmail(ctx context.Context, email string) (*users.User, error)
}

type TokenService interface {
	GenerateTokens(userID int) (accessToken, refreshToken string, err error)
}

type AuthService struct {
	userService UserService
}

func NewAuthService(u UserService) *AuthService {
	return &AuthService{userService: u}
}

func (as *AuthService) Authenticate(ctx context.Context, email, password string) (*users.User, error) {
	u, err := as.userService.FindByEmail(ctx, email)
	if err != nil {
		return nil, err
	}

	if err := cryptox.VerifyPassword(u.Password, password); err != nil {
		return nil, err
	}

	return u, nil
}
