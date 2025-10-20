package categories

import (
	"context"
	"ecommerce-service/pkg/httpx"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
)

type (
	Service interface {
		FindAll(ctx context.Context) ([]Category, error)
		FindByID(ctx context.Context, id int) (*Category, error)
		FindByName(ctx context.Context, name string) (*Category, error)
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
		httpx.Error(w, http.StatusInternalServerError, "Failed to fetch categories")
		return
	}

	httpx.JSON(w, http.StatusOK, categories)
}

func (h *CategoryHandler) FindByID(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	idStr := chi.URLParam(r, "id")
	if idStr == "" {
		httpx.Error(w, http.StatusBadRequest, "Category ID is required")
		return
	}

	id, err := strconv.Atoi(idStr)
	if err != nil {
		httpx.Error(w, http.StatusBadRequest, "Invalid category ID")
		return
	}

	category, err := h.categoryService.FindByID(ctx, id)
	if err != nil {
		httpx.Error(w, http.StatusNotFound, "Category not found")
		return
	}

	httpx.JSON(w, http.StatusOK, category)
}

func (h *CategoryHandler) FindByName(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	name := chi.URLParam(r, "name")
	if name == "" {
		httpx.Error(w, http.StatusBadRequest, "Category name is required")
		return
	}

	category, err := h.categoryService.FindByName(ctx, name)
	if err != nil {
		httpx.Error(w, http.StatusNotFound, "Category not found")
		return
	}

	httpx.JSON(w, http.StatusOK, category)
}
