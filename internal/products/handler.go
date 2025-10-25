package products

import (
	"context"
	"encoding/json"
	"net/http"
	"strconv"
	"strings"

	"ecommerce-service/internal/config"
	"ecommerce-service/internal/utils"
	"ecommerce-service/pkg/httpx"

	"github.com/go-chi/chi/v5"
	"github.com/go-playground/validator/v10"
)

type Service interface {
	Create(ctx context.Context, data *CreateProductRequest) error
	FindByID(ctx context.Context, id int) (*Product, error)
	FindAll(ctx context.Context, limit, offset int) ([]Product, error)
	Update(ctx context.Context, id int, data UpdateProductRequest) error
	Delete(ctx context.Context, id int) error
	Count(ctx context.Context) (int, error)
}

type ProductHandler struct {
	productService Service
	validate       *validator.Validate
	config         *config.Config
}

func NewProductHandler(productService Service) *ProductHandler {
	return &ProductHandler{
		productService: productService,
		validate:       validator.New(),
		config:         config.LoadEnvVars(),
	}
}

func (ph *ProductHandler) Create(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	product := &CreateProductRequest{}
	err := json.NewDecoder(r.Body).Decode(product)
	if err != nil {
		httpx.HTTPError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}

	err = ph.validate.Struct(product)
	if err != nil {
		validationErrors := httpx.FormatValidatorErrors(err)
		httpx.HTTPErrors(w, http.StatusBadRequest, validationErrors)
		return
	}

	err = ph.productService.Create(ctx, product)
	if err != nil {
		httpx.HTTPError(w, http.StatusInternalServerError, "Error creating product")
		return
	}

	httpx.HTTPResponse(w, http.StatusCreated, map[string]string{"message": "Product created successfully"})
}

func (ph *ProductHandler) FindAll(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	name := r.URL.Query().Get("name")
	limitStr := r.URL.Query().Get("limit")
	pageStr := r.URL.Query().Get("page")

	page, limit := utils.ParsePaginationParams(pageStr, limitStr, ph.config.Limit, ph.config.MaxLimit)
	total, err := ph.productService.Count(ctx)
	if err != nil {
		httpx.HTTPError(w, http.StatusInternalServerError, "Error counting products")
		return
	}

	products, err := ph.productService.FindAll(ctx, limit, page)
	if err != nil {
		httpx.HTTPError(w, http.StatusInternalServerError, "Error fetching products")
		return
	}

	if name != "" {
		for _, product := range products {
			if strings.EqualFold(product.Name, name) {
				httpx.HTTPResponse(w, http.StatusOK, []Product{product})
				return
			}
		}
	}

	if products == nil {
		products = []Product{}
	}

	httpx.HTTPPaginatedResponse(w, http.StatusOK, products, page, limit, total)
}

func (ph *ProductHandler) FindByID(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	idStr := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		httpx.HTTPError(w, http.StatusBadRequest, "Invalid product ID")
		return
	}

	product, err := ph.productService.FindByID(ctx, id)
	if err != nil && product == nil {
		httpx.HTTPError(w, http.StatusNotFound, "Product not found")
		return
	} else if err != nil {
		httpx.HTTPError(w, http.StatusInternalServerError, "Error fetching product")
	}

	httpx.HTTPResponse(w, http.StatusOK, &product)
}

func (ph *ProductHandler) Update(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	idStr := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		httpx.HTTPError(w, http.StatusBadRequest, "Invalid product ID")
		return
	}

	product := UpdateProductRequest{}
	if err = json.NewDecoder(r.Body).Decode(&product); err != nil || product == (UpdateProductRequest{}) {
		httpx.HTTPError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}

	err = ph.validate.Struct(product)
	if err != nil {
		validationErrors := httpx.FormatValidatorErrors(err)
		httpx.HTTPErrors(w, http.StatusBadRequest, validationErrors)
		return
	}

	err = ph.productService.Update(ctx, id, product)
	if err != nil {
		httpx.HTTPError(w, http.StatusInternalServerError, err.Error())
		return
	}

	httpx.HTTPResponse(w, http.StatusOK, map[string]string{"message": "Product updated successfully"})
}

func (ph *ProductHandler) Delete(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	idStr := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		httpx.HTTPError(w, http.StatusBadRequest, "Invalid product ID")
		return
	}

	err = ph.productService.Delete(ctx, id)
	if err != nil {
		httpx.HTTPError(w, http.StatusInternalServerError, "Error deleting product")
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
