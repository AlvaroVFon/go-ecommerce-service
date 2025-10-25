package orders

import (
	"context"
	"encoding/json"
	"net/http"
	"strconv"

	"ecommerce-service/pkg/httpx"

	"github.com/go-chi/chi/v5"
	"github.com/go-playground/validator/v10"
)

type (
	Service interface {
		Create(ctx context.Context, o *CreateOrderRequest) error
		FindByID(ctx context.Context, id int) (*Order, error)
		ListByUserID(ctx context.Context, userID int) ([]*Order, error)
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
		httpx.HTTPError(w, http.StatusBadRequest, httpx.BadRequestError)
		return
	}

	if err := h.validate.Struct(req); err != nil {
		validationErrors := httpx.FormatValidatorErrors(err)
		httpx.HTTPErrors(w, http.StatusBadRequest, validationErrors)
		return
	}

	if err := h.orderService.Create(ctx, &req); err != nil {
		httpx.HTTPError(w, http.StatusInternalServerError, httpx.InternalServerError)
		return
	}

	httpx.HTTPResponse(w, http.StatusCreated, map[string]string{"message": httpx.CreatedResponse})
}

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

// TODO: Paginate results
func (h *OrdersHandler) ListByUserID(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	idStr := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		httpx.HTTPError(w, http.StatusBadRequest, httpx.InvalidIDError)
		return
	}
	orders, err := h.orderService.ListByUserID(ctx, id)
	if err != nil {
		httpx.HTTPError(w, http.StatusInternalServerError, httpx.InternalServerError)
		return
	}

	httpx.HTTPResponse(w, http.StatusOK, orders)
}

func (h *OrdersHandler) Update(w http.ResponseWriter, r *http.Request) {}

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
