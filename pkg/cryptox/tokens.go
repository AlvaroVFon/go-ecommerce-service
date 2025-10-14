package cryptox

import (
	"ecommerce-service/internal/config"
)

const (
	AccessToken  = "access"
	RefreshToken = "refresh"
	ResetToken   = "reset"
	VerifyToken  = "verify"
)

type TokenService struct {
	config *config.Config
}

func NewTokenService(c *config.Config) *TokenService {
	return &TokenService{config: c}
}

func (ts *TokenService) CreateToken(payload any, tokenType string) {
}
