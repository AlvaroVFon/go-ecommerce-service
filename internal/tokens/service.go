// Package tokens provides functionality for managing and validating authentication tokens.
package tokens

import (
	"ecommerce-service/internal/config"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type TokenService struct {
	config *config.Config
}

type TokenType int

const (
	Accesss TokenType = iota
	Refresh
)

func NewTokenService(c *config.Config) *TokenService {
	return &TokenService{config: c}
}

func (ts *TokenService) GenerateToken(userID int, tokenType TokenType, exp int) (string, error) {
	claims := jwt.MapClaims{
		"user_id": userID,
		"type":    tokenType,
		"exp":     time.Now().Add(time.Duration(exp) * time.Second).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	secret := []byte(ts.config.JWTSecret)
	return token.SignedString(secret)
}

func (ts *TokenService) GenerateAccessToken(userID int, exp int) (string, error) {
	return ts.GenerateToken(userID, Accesss, exp)
}

func (ts *TokenService) GenerateRefreshToken(userID int, exp int) (string, error) {
	return ts.GenerateToken(userID, Refresh, exp)
}

func (ts *TokenService) GenerateTokens(userID int) (accessToken, refreshToken string, err error) {
	accessToken, err = ts.GenerateAccessToken(userID, ts.config.JWTExp)
	if err != nil {
		return "", "", err
	}

	refreshToken, err = ts.GenerateRefreshToken(userID, ts.config.JWTRefreshExp)
	if err != nil {
		return "", "", err
	}

	return accessToken, refreshToken, nil
}

func VerifyToken(tokenStr string, secret string) (*jwt.Token, error) {
	return jwt.Parse(tokenStr, func(token *jwt.Token) (any, error) {
		return []byte(secret), nil
	}, jwt.WithValidMethods([]string{jwt.SigningMethodHS256.Alg()}))
}

func ExtractClaims(token *jwt.Token) (jwt.MapClaims, error) {
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return claims, nil
	}

	return nil, jwt.ErrTokenInvalidClaims
}
