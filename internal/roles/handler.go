package roles

import (
	"context"
	"net/http"
	"strconv"

	"ecommerce-service/pkg/httpx"

	"github.com/go-chi/chi/v5"
)

type Service interface {
	FindAll(ctx context.Context) ([]Role, error)
	FindByID(ctx context.Context, id int) (*Role, error)
}

type RoleHandler struct {
	roleService Service
}

func NewRoleHandler(roleService Service) *RoleHandler {
	return &RoleHandler{roleService: roleService}
}

func (h *RoleHandler) FindAll(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	roles, err := h.roleService.FindAll(ctx)
	if err != nil {
		httpx.HTTPError(w, http.StatusBadRequest, httpx.BadRequestError)
		return
	}

	httpx.HTTPResponse(w, http.StatusOK, roles)
}

func (h *RoleHandler) FindByID(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	idStr := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		httpx.HTTPError(w, http.StatusBadRequest, httpx.BadRequestError)
		return
	}

	role, err := h.roleService.FindByID(ctx, id)
	if err != nil {
		httpx.HTTPError(w, http.StatusNotFound, httpx.NotFoundError)
		return
	}

	httpx.HTTPResponse(w, http.StatusOK, role)
}
