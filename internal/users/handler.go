package users

import (
	"context"
	"ecommerce-service/pkg/httpx"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/go-playground/validator/v10"
)

type Service interface {
	Create(ctx context.Context, u *CreateUserRequest) error
	FindByID(ctx context.Context, id int) (*PublicUser, error)
	FindAll(ctx context.Context) ([]PublicUser, error)
	Update(ctx context.Context, id int, u UpdateUserRequest) error
}

type UserHandler struct {
	userService Service
	validate    *validator.Validate
}

func NewUserHandler(userService Service) *UserHandler {
	return &UserHandler{userService: userService, validate: validator.New()}
}

func (uh *UserHandler) Create(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	req := CreateUserRequest{}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		httpx.Error(w, http.StatusBadRequest, "Invalid request payload")
	}

	err := uh.validate.Struct(req)
	if err != nil {
		validationErrors := err.(validator.ValidationErrors)
		httpx.Error(w, http.StatusBadRequest, validationErrors.Error())
		return
	}

	if err := uh.userService.Create(ctx, &req); err != nil {
		httpx.Error(w, http.StatusInternalServerError, "Internal server error")
		return
	}

	httpx.JSON(w, http.StatusCreated, map[string]string{"message": "user created successfully"})
}

func (uh *UserHandler) FindByID(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	idStr := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		httpx.Error(w, http.StatusBadRequest, "Invalid user ID")
		return
	}

	u, err := uh.userService.FindByID(ctx, id)
	if err != nil && u == nil {
		httpx.Error(w, http.StatusNotFound, "User not found")
		return
	} else if err != nil {
		httpx.Error(w, http.StatusInternalServerError, "Internal server error")
	}

	httpx.JSON(w, http.StatusOK, &u)
}

func (uh *UserHandler) FindAll(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	users, err := uh.userService.FindAll(ctx)
	if err != nil {
		httpx.Error(w, http.StatusInternalServerError, "Internal server error")
	}

	httpx.JSON(w, http.StatusOK, &users)
}

func (uh *UserHandler) Update(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	idStr := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		httpx.Error(w, http.StatusBadRequest, "Invalid user ID")
		return
	}

	req := UpdateUserRequest{}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil || req == (UpdateUserRequest{}) {
		httpx.Error(w, http.StatusBadRequest, "Invalid request payload")
		return
	}

	if err = uh.validate.Struct(req); err != nil {
		validationErrors := err.(validator.ValidationErrors)
		httpx.Error(w, http.StatusBadRequest, validationErrors.Error())
		return
	}

	if err := uh.userService.Update(ctx, id, req); err != nil {
		httpx.Error(w, http.StatusInternalServerError, "Internal server error")
		return
	}

	httpx.JSON(w, http.StatusOK, map[string]string{"message": "user updated successfully"})
}
