package categories

import (
	"context"
	"net/http"
	"strconv"

	"ecommerce-service/internal/config"
	"ecommerce-service/internal/utils"
	"ecommerce-service/pkg/httpx"

	"github.com/go-chi/chi/v5"
)

type (
	Service interface {
		FindAll(ctx context.Context, page, limit int) ([]Category, error)
		FindByID(ctx context.Context, id int) (*Category, error)
		Count(ctx context.Context) (int, error)
	}

	CategoryHandler struct {
		categoryService Service
		config          *config.Config
	}
)

func NewCategoryHandler(categoryService Service, config *config.Config) *CategoryHandler {
	return &CategoryHandler{categoryService: categoryService, config: config}
}

func (h *CategoryHandler) FindAll(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	pageStr := r.URL.Query().Get("page")
	limitStr := r.URL.Query().Get("limit")

	page, limit := utils.ParsePaginationParams(pageStr, limitStr, h.config.Limit, h.config.MaxLimit)
	total, err := h.categoryService.Count(ctx)
	if err != nil {
		httpx.HTTPError(w, http.StatusInternalServerError, httpx.InternalServerError)
		return
	}

	categories, err := h.categoryService.FindAll(ctx, page, limit)
	if err != nil {
		httpx.HTTPError(w, http.StatusInternalServerError, httpx.InternalServerError)
		return
	}

	httpx.HTTPPaginatedResponse(w, http.StatusOK, categories, page, limit, total)
}

func (h *CategoryHandler) FindByID(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	idStr := chi.URLParam(r, "id")
	if idStr == "" {
		httpx.HTTPError(w, http.StatusBadRequest, httpx.InvalidIDError)
		return
	}

	id, err := strconv.Atoi(idStr)
	if err != nil {
		httpx.HTTPError(w, http.StatusBadRequest, httpx.InvalidIDError)
		return
	}

	category, err := h.categoryService.FindByID(ctx, id)
	if err != nil {
		httpx.HTTPError(w, http.StatusNotFound, httpx.NotFoundError)
		return
	}

	httpx.HTTPResponse(w, http.StatusOK, category)
}
