package carts

import (
	"context"
	"net/http"
	"strconv"

	"ecommerce-service/pkg/httpx"

	"github.com/go-chi/chi/v5"
	"github.com/go-playground/validator/v10"
)

type (
	Service interface {
		GetCart(ctx context.Context, userID int64) (*Cart, error)
		AddItemToCart(ctx context.Context, userID int64, productID int64, quantity int) (*Cart, error)
		ClearCart(ctx context.Context, userID int64) error
		CompleteCart(ctx context.Context, userID int64) error
	}
	CartHandler struct {
		cartService Service
		validate    *validator.Validate
	}
)

func NewCartHandler(cartService Service, validate *validator.Validate) *CartHandler {
	return &CartHandler{cartService: cartService, validate: validate}
}

func (h *CartHandler) GetCart(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	cartIDStr := chi.URLParam(r, "id")
	cartID, err := strconv.ParseInt(cartIDStr, 10, 64)
	if err != nil {
		httpx.HTTPError(w, http.StatusBadRequest, httpx.InvalidIDError)
		return
	}

	cart, err := h.cartService.GetCart(ctx, cartID)
	if err != nil {
		httpx.HTTPError(w, http.StatusNotFound, httpx.NotFoundError)
		return
	}

	httpx.HTTPResponse(w, http.StatusOK, &cart)
}

func (h *CartHandler) AddItemToCart(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	userIDStr := chi.URLParam(r, "id")

	userID, err := strconv.ParseInt(userIDStr, 10, 64)
	if err != nil {
		httpx.HTTPError(w, http.StatusBadRequest, httpx.InvalidIDError)
		return
	}

	var req struct {
		ProductID int64 `json:"product_id"`
		Quantity  int   `json:"quantity"`
	}

	if err := httpx.ParseJSON(r, &req); err != nil {
		httpx.HTTPError(w, http.StatusBadRequest, httpx.BadRequestError)
		return
	}

	if err := h.validate.Struct(req); err != nil {
		httpx.HTTPResponse(w, http.StatusBadRequest, httpx.FormatValidatorErrors(err))
		return
	}

	cart, err := h.cartService.AddItemToCart(ctx, userID, req.ProductID, req.Quantity)
	if err != nil {
		httpx.HTTPError(w, http.StatusInternalServerError, httpx.InternalServerError)
		return
	}

	httpx.HTTPResponse(w, http.StatusOK, &cart)
}

func (h *CartHandler) ClearCart(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	userIDStr := chi.URLParam(r, "id")
	userID, err := strconv.ParseInt(userIDStr, 10, 64)
	if err != nil {
		httpx.HTTPError(w, http.StatusBadRequest, httpx.InvalidIDError)
		return
	}
	err = h.cartService.ClearCart(ctx, userID)
	if err != nil {
		httpx.HTTPError(w, http.StatusInternalServerError, httpx.InternalServerError)
		return
	}

	httpx.HTTPResponse(w, http.StatusOK, map[string]string{"message": httpx.DeletedResponse})
}

func (h *CartHandler) CompleteCart(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	userIDStr := chi.URLParam(r, "id")
	userID, err := strconv.ParseInt(userIDStr, 10, 64)
	if err != nil {
		httpx.HTTPError(w, http.StatusBadRequest, httpx.InvalidIDError)
		return
	}
	err = h.cartService.CompleteCart(ctx, userID)
	if err != nil {
		httpx.HTTPError(w, http.StatusInternalServerError, httpx.InternalServerError)
		return
	}

	httpx.HTTPResponse(w, http.StatusOK, map[string]string{"message": httpx.OkResponse})
}
