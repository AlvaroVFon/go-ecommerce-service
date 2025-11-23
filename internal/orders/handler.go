package orders

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"strconv"

	"ecommerce-service/internal/config"
	"ecommerce-service/internal/utils"
	"ecommerce-service/pkg/httpx"

	"github.com/go-chi/chi/v5"
	"github.com/go-playground/validator/v10"
)

type (
	// Service interface defines the methods we expect from our OrderService.
	Service interface {
		CreateOrderFromCart(ctx context.Context, o *CreateOrderRequest) (*Order, error)
		FindByID(ctx context.Context, id int) (*Order, error)
		ListByUserID(ctx context.Context, userID, page, limit int) ([]*Order, error)
		Update(ctx context.Context, id int, o *UpdateOrderRequest) error
		Delete(ctx context.Context, id int) error
		CountByUserID(ctx context.Context, userID int) (int, error)
	}

	// OrdersHandler is the HTTP handler for orders.
	OrdersHandler struct {
		orderService Service
		validate     *validator.Validate
		config       *config.Config
	}
)

// NewOrderHandler creates a new OrdersHandler.
func NewOrderHandler(orderService Service, validate *validator.Validate, config *config.Config) *OrdersHandler {
	return &OrdersHandler{orderService: orderService, validate: validate, config: config}
}

// Create handles the HTTP request to create a new order from a cart.
func (h *OrdersHandler) Create(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	var req CreateOrderRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		httpx.HTTPError(w, http.StatusBadRequest, httpx.BadRequestError)
		return
	}

	if err := h.validate.Struct(req); err != nil {
		validationErrors := httpx.FormatValidatorErrors(err)
		httpx.HTTPErrors(w, http.StatusBadRequest, validationErrors)
		return
	}

	createdOrder, err := h.orderService.CreateOrderFromCart(ctx, &req)
	if err != nil {
		if errors.Is(err, ErrEmptyCart) {
			httpx.HTTPError(w, http.StatusBadRequest, err.Error())
			return
		}
		httpx.HTTPError(w, http.StatusInternalServerError, httpx.InternalServerError)
		return
	}

	httpx.HTTPResponse(w, http.StatusCreated, createdOrder)
}

// FindByID handles the HTTP request to find an order by its ID.
func (h *OrdersHandler) FindByID(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	idStr := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		httpx.HTTPError(w, http.StatusBadRequest, httpx.InvalidIDError)
		return
	}

	order, err := h.orderService.FindByID(ctx, id)
	if err != nil {
		httpx.HTTPError(w, http.StatusNotFound, httpx.NotFoundError)
		return
	}

	httpx.HTTPResponse(w, http.StatusOK, order)
}

// ListByUserID handles the HTTP request to list orders for a given user.
func (h *OrdersHandler) ListByUserID(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	idStr := chi.URLParam(r, "id")
	limitStr := r.URL.Query().Get("limit")
	pageStr := r.URL.Query().Get("page")

	page, limit := utils.ParsePaginationParams(pageStr, limitStr, h.config.Limit, h.config.MaxLimit)

	id, err := strconv.Atoi(idStr)
	if err != nil {
		httpx.HTTPError(w, http.StatusBadRequest, httpx.InvalidIDError)
		return
	}

	total, err := h.orderService.CountByUserID(ctx, id)
	if err != nil {
		httpx.HTTPError(w, http.StatusInternalServerError, httpx.InternalServerError)
		return
	}
	orders, err := h.orderService.ListByUserID(ctx, id, page, limit)
	if err != nil {
		httpx.HTTPError(w, http.StatusInternalServerError, httpx.InternalServerError)
		return
	}

	httpx.HTTPPaginatedResponse(w, http.StatusOK, orders, page, limit, total)
}

// Update handles the HTTP request to update an order.
func (h *OrdersHandler) Update(w http.ResponseWriter, r *http.Request) {}

// Delete handles the HTTP request to delete an order.
func (h *OrdersHandler) Delete(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	idStr := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		httpx.HTTPError(w, http.StatusBadRequest, httpx.InvalidIDError)
		return
	}

	if err := h.orderService.Delete(ctx, id); err != nil {
		httpx.HTTPError(w, http.StatusInternalServerError, httpx.InternalServerError)
		return
	}

	httpx.HTTPResponse(w, http.StatusNoContent, map[string]string{})
}
