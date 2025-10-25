package categories

import (
	"context"
	"net/http"
	"strconv"

	"ecommerce-service/pkg/httpx"

	"github.com/go-chi/chi/v5"
)

type (
	Service interface {
		FindAll(ctx context.Context) ([]Category, error)
		FindByID(ctx context.Context, id int) (*Category, error)
	}

	CategoryHandler struct {
		categoryService Service
	}
)

func NewCategoryHandler(categoryService Service) *CategoryHandler {
	return &CategoryHandler{categoryService: categoryService}
}

func (h *CategoryHandler) FindAll(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	categories, err := h.categoryService.FindAll(ctx)
	if err != nil {
		httpx.HTTPError(w, http.StatusInternalServerError, "Failed to fetch categories")
		return
	}

	httpx.HTTPResponse(w, http.StatusOK, categories)
}

func (h *CategoryHandler) FindByID(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	idStr := chi.URLParam(r, "id")
	if idStr == "" {
		httpx.HTTPError(w, http.StatusBadRequest, "Category ID is required")
		return
	}

	id, err := strconv.Atoi(idStr)
	if err != nil {
		httpx.HTTPError(w, http.StatusBadRequest, "Invalid category ID")
		return
	}

	category, err := h.categoryService.FindByID(ctx, id)
	if err != nil {
		httpx.HTTPError(w, http.StatusNotFound, "Category not found")
		return
	}

	httpx.HTTPResponse(w, http.StatusOK, category)
}
