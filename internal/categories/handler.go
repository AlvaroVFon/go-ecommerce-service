package categories

import (
	"context"
	"net/http"
	"strconv"
	"strings"

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
	name := r.URL.Query().Get("name")

	categories, err := h.categoryService.FindAll(ctx)
	if err != nil {
		httpx.Error(w, http.StatusInternalServerError, "Failed to fetch categories")
		return
	}

	if name != "" {
		for _, category := range categories {
			if strings.EqualFold(category.Name, name) {
				httpx.JSON(w, http.StatusOK, []Category{category})
				return
			}
		}
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
