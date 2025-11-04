package seeds

import (
	"database/sql"
	"ecommerce-service/internal/orders"
)

func SeedOrders(db *sql.DB) error {
	orders := []orders.Order{
		{
			UserID:          1,
			Status:          "pending",
			Total:           1443,
			ShippingAddress: "123 Main St, Anytown, USA",
			PaymentMethod:   "credit_card",
		},
		{
			UserID:          2,
			Status:          "shipped",
			Total:           944,
			ShippingAddress: "456 Oak Ave, Anytown, USA",
			PaymentMethod:   "paypal",
		},
	}
	query := "INSERT INTO orders (user_id, status, total, shipping_address, payment_method) VALUES ($1, $2, $3, $4, $5) ON CONFLICT DO NOTHING"

	for _, order := range orders {
		if _, err := db.Exec(query, order.UserID, order.Status, order.Total, order.ShippingAddress, order.PaymentMethod); err != nil {
			return err
		}
	}
	return nil
}

func SeedOrderItems(db *sql.DB) error {
	orderItems := []orders.OrderItem{
		{
			OrderID:  1,
			ProductID: 1,
			Quantity:  1,
			Price:     1200.00,
		},
		{
			OrderID:  1,
			ProductID: 3,
			Quantity:  1,
			Price:     150.00,
		},
		{
			OrderID:  2,
			ProductID: 2,
			Quantity:  1,
			Price:     800.00,
		},
	}
	query := "INSERT INTO order_items (order_id, product_id, quantity, price) VALUES ($1, $2, $3, $4) ON CONFLICT DO NOTHING"

	for _, item := range orderItems {
		if _, err := db.Exec(query, item.OrderID, item.ProductID, item.Quantity, item.Price); err != nil {
			return err
		}
	}
	return nil
}
