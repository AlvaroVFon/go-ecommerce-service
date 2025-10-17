// Package auth implements authentication and authorization services.
package auth

import (
	"context"
	"ecommerce-service/internal/users"
	"errors"
)

var ErrStrategyNotFound = errors.New("authentication strategy not found")

type AuthStrategy interface {
	Authenticate(ctx context.Context, credentials any) (*users.User, error)
}

type AuthService struct {
	strategies map[string]AuthStrategy
}

func NewAuthService(strategies map[string]AuthStrategy) *AuthService {
	return &AuthService{strategies: strategies}
}

func (as *AuthService) Authenticate(ctx context.Context, strategyName string, credentials any) (*users.User, error) {
	strategy, ok := as.strategies[strategyName]
	if !ok {
		return nil, ErrStrategyNotFound
	}

	return strategy.Authenticate(ctx, credentials)
}
