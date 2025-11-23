package orders

import (
	"context"
	"errors"
	"fmt"
	"log"

	"ecommerce-service/internal/carts"
)

var ErrEmptyCart = errors.New("cannot create order from an empty cart")

type (
	// Repository is the interface for the order repository.
	Repository interface {
		Create(ctx context.Context, o *Order) (*Order, error)
		FindByID(ctx context.Context, id int) (*Order, error)
		ListByUserID(ctx context.Context, userID, limit, offset int) ([]*Order, error)
		Update(ctx context.Context, id int, o *UpdateOrderRequest) error
		Delete(ctx context.Context, id int) error
		CountByUserID(ctx context.Context, userID int) (int, error)
	}

	// CartRepository defines the dependency on the cart repository.
	CartRepository interface {
		GetItems(ctx context.Context, cartID int64) ([]carts.CartItem, error)
		SetCompleted(ctx context.Context, cartID int64) error
		ClearCart(ctx context.Context, cartID int64) error
	}

	// OrderService is the service for managing orders.
	OrderService struct {
		orderRepo Repository
		cartRepo  CartRepository
	}
)

// NewOrderService creates a new OrderService.
func NewOrderService(orderRepo Repository, cartRepo CartRepository) *OrderService {
	return &OrderService{
		orderRepo: orderRepo,
		cartRepo:  cartRepo,
	}
}

// CreateOrderFromCart creates a new order from a shopping cart.
func (s *OrderService) CreateOrderFromCart(ctx context.Context, req *CreateOrderRequest) (*Order, error) {
	// 1. Get items from the cart
	cartItems, err := s.cartRepo.GetItems(ctx, req.CartID)
	if err != nil {
		return nil, fmt.Errorf("failed to get cart items: %w", err)
	}
	if len(cartItems) == 0 {
		return nil, ErrEmptyCart
	}

	// 2. Calculate total and prepare order items
	var total float64
	orderItems := make([]OrderItem, 0, len(cartItems))
	for _, item := range cartItems {
		price := float64(item.TotalPrice) / 100.0 // Convert cents to dollars, includes discount
		total += price
		unitPrice := price / float64(item.Quantity)
		orderItems = append(orderItems, OrderItem{
			ProductID: item.ProductID,
			Quantity:  int(item.Quantity),
			Price:     unitPrice, // Unit price after discount
		})
	}

	// 3. Create the order object
	order := &Order{
		UserID:          req.UserID,
		Items:           orderItems,
		Total:           total,
		Status:          "pending", // Initial status
		ShippingAddress: req.ShippingAddress,
		PaymentMethod:   req.PaymentMethod,
	}

	// 4. Use the repository to create the order transactionally
	createdOrder, err := s.orderRepo.Create(ctx, order)
	if err != nil {
		return nil, fmt.Errorf("failed to create order in repository: %w", err)
	}

	// 5. Mark cart as completed and clear it.
	// Log errors but don't fail the order creation, as these are cleanup steps.
	if err := s.cartRepo.SetCompleted(ctx, req.CartID); err != nil {
		log.Printf("warning: failed to mark cart %d as completed: %v\n", req.CartID, err)
	}
	if err := s.cartRepo.ClearCart(ctx, req.CartID); err != nil {
		log.Printf("warning: failed to clear cart %d: %v\n", req.CartID, err)
	}

	return createdOrder, nil
}

// FindByID is a pass-through to the repository.
func (s *OrderService) FindByID(ctx context.Context, id int) (*Order, error) {
	return s.orderRepo.FindByID(ctx, id)
}

// ListByUserID is a pass-through to the repository.
func (s *OrderService) ListByUserID(ctx context.Context, userID, page, limit int) ([]*Order, error) {
	offset := (page - 1) * limit
	return s.orderRepo.ListByUserID(ctx, userID, limit, offset)
}

// Update is a pass-through to the repository.
func (s *OrderService) Update(ctx context.Context, id int, o *UpdateOrderRequest) error {
	return s.orderRepo.Update(ctx, id, o)
}

// Delete is a pass-through to the repository.
func (s *OrderService) Delete(ctx context.Context, id int) error {
	return s.orderRepo.Delete(ctx, id)
}

// CountByUserID is a pass-through to the repository.
func (s *OrderService) CountByUserID(ctx context.Context, userID int) (int, error) {
	return s.orderRepo.CountByUserID(ctx, userID)
}
