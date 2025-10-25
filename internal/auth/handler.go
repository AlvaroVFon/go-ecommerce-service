package auth

import (
	"context"
	"net/http"

	"ecommerce-service/internal/auth/strategies"
	"ecommerce-service/internal/users"
	"ecommerce-service/pkg/httpx"
)

type (
	Service interface {
		Authenticate(ctx context.Context, strategyName string, credentials any) (*users.User, error)
	}

	TokensService interface {
		GenerateTokens(userID int) (accessToken, refreshToken string, err error)
	}

	AuthHandler struct {
		authService   Service
		tokensService TokensService
	}
)

func NewAuthHandler(a Service, t TokensService) *AuthHandler {
	return &AuthHandler{authService: a, tokensService: t}
}

func (ah *AuthHandler) Login(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var req strategies.PasswordCredentials
	err := httpx.ParseJSON(r, &req)
	if err != nil {
		httpx.HTTPError(w, http.StatusBadRequest, httpx.BadRequestError)
		return
	}

	u, err := ah.authService.Authenticate(ctx, "password", req)
	if err != nil {
		httpx.HTTPError(w, http.StatusUnauthorized, httpx.UnauthorizedError)
		return
	}

	accessToken, refreshToken, err := ah.tokensService.GenerateTokens(u.ID)
	if err != nil {
		httpx.HTTPError(w, http.StatusInternalServerError, httpx.InternalServerError)
		return
	}

	httpx.HTTPResponse(w, http.StatusOK, map[string]string{
		"access_token":  accessToken,
		"refresh_token": refreshToken,
	})
}
