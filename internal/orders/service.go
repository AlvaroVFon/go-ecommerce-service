package orders

import "context"

type (
	Repository interface {
		Create(ctx context.Context, o *CreateOrderRequest) error
		FindByID(ctx context.Context, id int) (*Order, error)
		ListByUserID(ctx context.Context, userID, limit, offset int) ([]*Order, error)
		Update(ctx context.Context, id int, o *UpdateOrderRequest) error
		Delete(ctx context.Context, id int) error
		CountByUserID(ctx context.Context, userID int) (int, error)
	}

	OrderService struct {
		orderRepo Repository
	}
)

func NewOrderService(orderRepo Repository) *OrderService {
	return &OrderService{orderRepo: orderRepo}
}

func (r *OrderService) Create(ctx context.Context, o *CreateOrderRequest) error {
	return r.orderRepo.Create(ctx, o)
}

func (r *OrderService) FindByID(ctx context.Context, id int) (*Order, error) {
	return r.orderRepo.FindByID(ctx, id)
}

func (r *OrderService) ListByUserID(ctx context.Context, userID, page, limit int) ([]*Order, error) {
	offset := (page - 1) * limit
	return r.orderRepo.ListByUserID(ctx, userID, limit, offset)
}

func (r *OrderService) Update(ctx context.Context, id int, o *UpdateOrderRequest) error {
	return r.orderRepo.Update(ctx, id, o)
}

func (r *OrderService) Delete(ctx context.Context, id int) error {
	return r.orderRepo.Delete(ctx, id)
}

func (r *OrderService) CountByUserID(ctx context.Context, userID int) (int, error) {
	return r.orderRepo.CountByUserID(ctx, userID)
}
