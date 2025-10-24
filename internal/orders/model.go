// Package orders defines the data models for the orders module.
package orders

import "ecommerce-service/internal/carts"

type Order struct {
	ID              int64            `json:"id"`
	UserID          int64            `json:"user_id"`
	CartID          int64            `json:"cart_id"`
	CartItems       []carts.CartItem `json:"cart_items"`
	Status          string           `json:"status"`
	ShippingAddress string           `json:"shipping_address"`
	PaymentMethod   string           `json:"payment_method"`
	CreatedAt       int64            `json:"created_at"`
	UpdatedAt       int64            `json:"updated_at"`
}

type CreateOrderRequest struct {
	UserID          int64  `json:"user_id" validate:"required"`
	CartID          int64  `json:"cart_id" validate:"required"`
	ShippingAddress string `json:"shipping_address" validate:"required"`
	PaymentMethod   string `json:"payment_method" validate:"required"`
}

type UpdateOrderRequest struct {
	Status          *string `json:"status" validate:"required,oneof=pending processing shipped delivered cancelled"`
	ShippingAddress *string `json:"shipping_address,omitempty"`
}
