package product

import (
	"encoding/json"
	"net/http"

	"ecommerce-service/pkg/httpx"
)

type ProductHandler struct {
	productService *ProductService
}

func NewProductHandler(productService *ProductService) *ProductHandler {
	return &ProductHandler{productService: productService}
}

func (ph *ProductHandler) Create(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	product := &CreateProductDTO{}
	err := json.NewDecoder(r.Body).Decode(product)
	if err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}
	err = ph.productService.Create(ctx, product)
	if err != nil {
		httpx.Error(w, http.StatusInternalServerError, "Error creating product")
		return
	}

	httpx.JSON(w, http.StatusCreated, map[string]string{"message": "Product created successfully"})
}
