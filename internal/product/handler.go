package product

import (
	"context"
	"encoding/json"
	"net/http"
	"strconv"

	"ecommerce-service/pkg/httpx"

	"github.com/go-chi/chi/v5"
	"github.com/go-playground/validator/v10"
)

type Service interface {
	Create(ctx context.Context, data *CreateProductRequest) error
	FindById(ctx context.Context, id int) (*Product, error)
	FindAll(ctx context.Context) ([]Product, error)
	Update(ctx context.Context, id int, data *UpdateProductRequest) error
	Delete(ctx context.Context, id int) error
}

type ProductHandler struct {
	productService Service
	validate       *validator.Validate
}

func NewProductHandler(productService Service) *ProductHandler {
	return &ProductHandler{
		productService: productService,
		validate:       validator.New(),
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
		validationErrors := err.(validator.ValidationErrors)
		httpx.Error(w, http.StatusBadRequest, validationErrors.Error())
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
	products, err := ph.productService.FindAll(ctx)
	if err != nil {
		httpx.Error(w, http.StatusInternalServerError, "Error fetching products")
		return
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

	product, err := ph.productService.FindById(ctx, id)
	if err != nil {
		httpx.Error(w, http.StatusNotFound, "Product not found")
		return
	}

	httpx.JSON(w, http.StatusOK, product)
}

func (ph *ProductHandler) Update(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	idStr := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		httpx.Error(w, http.StatusBadRequest, "Invalid product ID")
		return
	}

	product := &UpdateProductRequest{}
	err = json.NewDecoder(r.Body).Decode(product)
	if err != nil {
		httpx.Error(w, http.StatusBadRequest, "Invalid request payload")
		return
	}

	err = ph.validate.Struct(product)
	if err != nil {
		validationErrors := err.(validator.ValidationErrors)
		httpx.Error(w, http.StatusBadRequest, validationErrors.Error())
		return
	}

	err = ph.productService.Update(ctx, id, product)
	if err != nil {
		httpx.Error(w, http.StatusInternalServerError, "Error updating product")
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
