// Package carts defines the data models for the shopping cart feature.
package carts

import "time"

type CartItem struct {
	CartID        int64     `json:"cart_id"`
	ProductID     int64     `json:"product_id"`
	Name          string    `json:"name"`
	Description   string    `json:"description"`
	Quantity      int64     `json:"quantity"`
	SnapshotPrice int64     `json:"snapshot_price"`
	DiscountRate  float64   `json:"discount_rate"` // Percentage discount applied
	TotalPrice    int64     `json:"total_price"`   // Quantity * SnapshotPrice - Discount
	ImageURL      string    `json:"image_url"`
	AddedAt       time.Time `json:"added_at"`
	UpdatedAt     time.Time `json:"updated_at"`
}

type Cart struct {
	ID        int64      `json:"id"`
	UserID    int64      `json:"user_id"`
	CartItems []CartItem `json:"cart_items"`

	// Calculated fields
	Subtotal int64 `json:"subtotal"` // Added price of all items before discounts and taxes
	Discount int64 `json:"discount"` // Total discount applied
	Tax      int64 `json:"tax"`
	Total    int64 `json:"total"` // Subtotal - Discount + Tax

	// Metadata
	Status    string     `json:"status"` // e.g., "active", "abandoned", "completed"
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	ExpiresAt *time.Time `json:"expires_at,omitempty"` // To clear abandoned carts
}
