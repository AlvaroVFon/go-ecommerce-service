package users

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"ecommerce-service/pkg/httpx"

	"github.com/go-chi/chi/v5"
	"github.com/go-playground/validator/v10"
)

type Service interface {
	Create(ctx context.Context, u *CreateUserRequest) error
	FindByID(ctx context.Context, id int) (*User, error)
	FindAll(ctx context.Context) ([]User, error)
	Update(ctx context.Context, id int, u UpdateUserRequest) error
	Delete(ctx context.Context, id int) error
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
		httpx.HTTPError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}

	err := uh.validate.Struct(req)
	if err != nil {
		validationErrors := httpx.FormatValidatorErrors(err)
		httpx.HTTPErrors(w, http.StatusBadRequest, validationErrors)
		return
	}

	if err := uh.userService.Create(ctx, &req); err != nil {
		httpx.HTTPError(w, http.StatusConflict, "User with this email already exists")
		return
	}

	httpx.HTTPResponse(w, http.StatusCreated, map[string]string{"message": "user created successfully"})
}

func (uh *UserHandler) FindByID(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	idStr := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		httpx.HTTPError(w, http.StatusBadRequest, "Invalid user ID")
		return
	}

	u, err := uh.userService.FindByID(ctx, id)
	if err != nil && u == nil {
		httpx.HTTPError(w, http.StatusNotFound, "User not found")
		return
	} else if err != nil {
		httpx.HTTPError(w, http.StatusInternalServerError, "Internal server error")
		return
	}

	pu := PublicUser{u.ID, u.Email, u.CreatedAt, u.UpdatedAt}

	httpx.HTTPResponse(w, http.StatusOK, &pu)
}

func (uh *UserHandler) FindAll(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	users, err := uh.userService.FindAll(ctx)
	publicUsers := []PublicUser{}
	for _, u := range users {
		publicUsers = append(publicUsers, PublicUser{u.ID, u.Email, u.CreatedAt, u.UpdatedAt})
	}
	if err != nil {
		log.Println("error fetching users:", err)
		httpx.HTTPError(w, http.StatusInternalServerError, "Internal server error")
		return
	}

	httpx.HTTPResponse(w, http.StatusOK, &publicUsers)
}

func (uh *UserHandler) Update(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	idStr := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		httpx.HTTPError(w, http.StatusBadRequest, "Invalid user ID")
		return
	}

	req := UpdateUserRequest{}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil || req == (UpdateUserRequest{}) {
		httpx.HTTPError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}

	if err = uh.validate.Struct(req); err != nil {
		validationErrors := httpx.FormatValidatorErrors(err)
		httpx.HTTPErrors(w, http.StatusBadRequest, validationErrors)
		return
	}

	if err := uh.userService.Update(ctx, id, req); err != nil {
		httpx.HTTPError(w, http.StatusInternalServerError, "Internal server error")
		return
	}

	httpx.HTTPResponse(w, http.StatusOK, map[string]string{"message": "user updated successfully"})
}

func (uh *UserHandler) Delete(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	idStr := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		httpx.HTTPError(w, http.StatusBadRequest, "Invalid user ID")
		return
	}

	if err := uh.userService.Delete(ctx, id); err != nil {
		httpx.HTTPError(w, http.StatusInternalServerError, "User not found")
		return
	}

	httpx.HTTPResponse(w, http.StatusOK, map[string]string{"message": "user deleted successfully"})
}
