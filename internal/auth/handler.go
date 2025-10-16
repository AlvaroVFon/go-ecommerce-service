package auth

import (
	"context"
	"ecommerce-service/internal/users"
	"ecommerce-service/pkg/httpx"
	"net/http"
)

type (
	Service interface {
		Authenticate(ctx context.Context, email, password string) (*users.User, error)
	}

	TokensService interface {
		GenerateTokens(userID int) (accessToken, refreshToken string, err error)
	}

	AuthHandler struct {
		authService   Service
		tokensService TokensService
	}
)

func NewAuthHandler(a Service, t TokenService) *AuthHandler {
	return &AuthHandler{authService: a, tokensService: t}
}

func (ah *AuthHandler) Login(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var req struct {
		Email    string
		Password string
	}

	err := httpx.ParseJSON(r, &req)
	if err != nil {
		httpx.Error(w, http.StatusBadRequest, "invalid request payload")
		return
	}

	u, err := ah.authService.Authenticate(ctx, req.Email, req.Password)
	if err != nil && u == nil {
		httpx.Error(w, http.StatusUnauthorized, "authentication failed")
		return
	}

	accessToken, refreshToken, err := ah.tokensService.GenerateTokens(u.ID)
	if err != nil {
		httpx.Error(w, http.StatusInternalServerError, "failed to generate tokens")
		return
	}

	httpx.JSON(w, http.StatusOK, map[string]string{
		"access_token":  accessToken,
		"refresh_token": refreshToken,
	})
}
