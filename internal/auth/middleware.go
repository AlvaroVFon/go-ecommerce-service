package auth

import (
	"context"
	"net/http"
	"strings"

	"ecommerce-service/internal/config"
	"ecommerce-service/pkg/httpx"

	"github.com/golang-jwt/jwt/v5"
)

type (
	TokenService interface {
		VerifyToken(tokenStr string, secret string) (*jwt.Token, error)
		ExtractClaims(token *jwt.Token) (jwt.MapClaims, error)
	}
	AuthMiddleware struct {
		tokenService TokenService
		config       *config.Config
	}

	contextKey string
)

const userClaimsKey contextKey = "user_id"

func NewAuthMiddleware(ts TokenService, c *config.Config) *AuthMiddleware {
	return &AuthMiddleware{tokenService: ts, config: c}
}

func (am *AuthMiddleware) VerifyToken(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		headerToken := r.Header.Get("Authorization")
		if headerToken == "" {
			httpx.HTTPError(w, http.StatusUnauthorized, "Missing Authorization header")
			return
		}

		tokenStr := strings.TrimPrefix(headerToken, "Bearer ")

		token, err := am.tokenService.VerifyToken(tokenStr, am.config.JWTSecret)
		if err != nil || !token.Valid {
			httpx.HTTPError(w, http.StatusUnauthorized, "Invalid token")
			return
		}

		claims, err := am.tokenService.ExtractClaims(token)
		if err != nil {
			httpx.HTTPError(w, http.StatusUnauthorized, "Failed to extract claims")
			return
		}

		ctx := context.WithValue(r.Context(), userClaimsKey, claims)

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

// TODO: Add role-based authorization middleware
