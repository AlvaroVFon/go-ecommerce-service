package products

import (
	"context"
	"encoding/json"
	"net/http"
	"strconv"

	"ecommerce-service/internal/config"
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
		httpx.Error(w, http.StatusBadRequest, "Invalid request payload")
		return
	}

	err = ph.validate.Struct(product)
	if err != nil {
		validationErrors := httpx.FormatValidatorErrors(err)
		httpx.Errors(w, http.StatusBadRequest, validationErrors)
		return
	}

	err = ph.productService.Create(ctx, product)
	if err != nil {
		httpx.Error(w, http.StatusInternalServerError, "Error creating product")
		return
	}

	httpx.JSON(w, http.StatusCreated, map[string]string{"message": "Product created successfully"})
}

func (ph *ProductHandler) FindAll(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	limitStr := r.URL.Query().Get("limit")
	offsetStr := r.URL.Query().Get("offset")

	limit, err := strconv.Atoi(limitStr)
	if limitStr != "" && err != nil {
		httpx.Error(w, http.StatusBadRequest, "Invalid limit parameter")
		return
	}
	offset, err := strconv.Atoi(offsetStr)
	if offsetStr != "" && err != nil {
		httpx.Error(w, http.StatusBadRequest, "Invalid offset parameter")
		return
	}

	products, err := ph.productService.FindAll(ctx, limit, offset)
	if err != nil {
		httpx.Error(w, http.StatusInternalServerError, "Error fetching products")
		return
	}

	if products == nil {
		products = []Product{}
	}

	httpx.JSON(w, http.StatusOK, products)
}

func (ph *ProductHandler) FindByID(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	idStr := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		httpx.Error(w, http.StatusBadRequest, "Invalid product ID")
		return
	}

	product, err := ph.productService.FindByID(ctx, id)
	if err != nil && product == nil {
		httpx.Error(w, http.StatusNotFound, "Product not found")
		return
	} else if err != nil {
		httpx.Error(w, http.StatusInternalServerError, "Error fetching product")
	}

	httpx.JSON(w, http.StatusOK, &product)
}

func (ph *ProductHandler) Update(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	idStr := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		httpx.Error(w, http.StatusBadRequest, "Invalid product ID")
		return
	}

	product := UpdateProductRequest{}
	if err = json.NewDecoder(r.Body).Decode(&product); err != nil || product == (UpdateProductRequest{}) {
		httpx.Error(w, http.StatusBadRequest, "Invalid request payload")
		return
	}

	err = ph.validate.Struct(product)
	if err != nil {
		validationErrors := httpx.FormatValidatorErrors(err)
		httpx.Errors(w, http.StatusBadRequest, validationErrors)
		return
	}

	err = ph.productService.Update(ctx, id, product)
	if err != nil {
		httpx.Error(w, http.StatusInternalServerError, err.Error())
		return
	}

	httpx.JSON(w, http.StatusOK, map[string]string{"message": "Product updated successfully"})
}

func (ph *ProductHandler) Delete(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	idStr := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		httpx.Error(w, http.StatusBadRequest, "Invalid product ID")
		return
	}

	err = ph.productService.Delete(ctx, id)
	if err != nil {
		httpx.Error(w, http.StatusInternalServerError, "Error deleting product")
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
