package orders

import (
	"context"
	"encoding/json"
	"net/http"

	"ecommerce-service/pkg/httpx"

	"github.com/go-playground/validator/v10"
)

type (
	Service interface {
		Create(ctx context.Context, o *CreateOrderRequest) error
		FindByID(ctx context.Context, id string) (*Order, error)
		ListByUserID(ctx context.Context, userID string) ([]*Order, error)
		Update(ctx context.Context, id int, o *UpdateOrderRequest) error
		Delete(ctx context.Context, id int) error
	}

	OrdersHandler struct {
		orderService Service
		validate     validator.Validate
	}
)

func NewOrderHandler(orderService Service) *OrdersHandler {
	return &OrdersHandler{orderService: orderService, validate: *validator.New()}
}

func (h *OrdersHandler) Create(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	var req CreateOrderRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		httpx.Error(w, http.StatusBadRequest, "invalid request body")
		return
	}

	if err := h.validate.Struct(req); err != nil {
		validationErrors := httpx.FormatValidatorErrors(err)
		httpx.Errors(w, http.StatusBadRequest, validationErrors)
		return
	}

	if err := h.orderService.Create(ctx, &req); err != nil {
		httpx.Error(w, http.StatusInternalServerError, err.Error())
		return
	}

	if err := h.orderService.Create(ctx, &req); err != nil {
		httpx.Error(w, http.StatusInternalServerError, err.Error())
		return
	}

	httpx.JSON(w, http.StatusCreated, map[string]string{"message": "order created successfully"})
}
