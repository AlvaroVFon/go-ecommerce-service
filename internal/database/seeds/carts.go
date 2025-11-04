package seeds

import (
	"database/sql"

	"ecommerce-service/internal/carts"
)

func SeedCarts(db *sql.DB) error {
	carts := []carts.Cart{
		{
			UserID:   1,
			Subtotal: 1350,
			Discount: 150,
			Tax:      243,
			Total:    1443,
		},
		{
			UserID:   2,
			Subtotal: 800,
			Discount: 0,
			Tax:      144,
			Total:    944,
		},
	}
	query := "INSERT INTO carts (user_id, subtotal, discount, tax, total) VALUES ($1, $2, $3, $4, $5)"

	for _, cart := range carts {
		if _, err := db.Exec(query, cart.UserID, cart.Subtotal, cart.Discount, cart.Tax, cart.Total); err != nil {
			return err
		}
	}
	return nil
}

func SeedCartItems(db *sql.DB) error {
	cartItems := []carts.CartItem{
		{
			CartID:        1,
			ProductID:     1,
			Name:          "Laptop",
			Description:   "A high-performance laptop",
			Quantity:      1,
			SnapshotPrice: 1200.00,
			DiscountRate:  10,
			TotalPrice:    1080,
			ImageURL:      "https://example.com/laptop.jpg",
		},
		{
			CartID:        1,
			ProductID:     3,
			Name:          "Coffee Maker",
			Description:   "A programmable coffee maker",
			Quantity:      1,
			SnapshotPrice: 150.00,
			DiscountRate:  0,
			TotalPrice:    150.00,
			ImageURL:      "https://example.com/coffe-maker.jpg",
		},
		{
			CartID:        2,
			ProductID:     2,
			Name:          "Smartphone",
			Description:   "A latest model smartphone",
			Quantity:      1,
			SnapshotPrice: 800.00,
			DiscountRate:  0,
			TotalPrice:    800.00,
			ImageURL:      "https://example.com/smartphone.jpg",
		},
	}
	query := "INSERT INTO cart_items (cart_id, product_id, name, description, quantity, snapshot_price, discount_rate, total_price, image_url) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9) ON CONFLICT (cart_id, product_id) DO NOTHING"

	for _, item := range cartItems {
		if _, err := db.Exec(query, item.CartID, item.ProductID, item.Name, item.Description, item.Quantity, item.SnapshotPrice, item.DiscountRate, item.TotalPrice, item.ImageURL); err != nil {
			return err
		}
	}
	return nil
}
