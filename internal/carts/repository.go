package carts

import (
	"context"
	"database/sql"
)

type CartRepository struct {
	db *sql.DB
}

func NewCartRepository(db *sql.DB) *CartRepository {
	return &CartRepository{db: db}
}

func (r *CartRepository) Create(ctx context.Context, userID int64) (*Cart, error) {
	query := "INSERT INTO carts (user_id, subtotal, total) VALUES ($1) RETURNING id"
	_, err := r.db.ExecContext(ctx, query, userID, 0, 0)
	if err != nil {
		return nil, err
	}
	var cartID int64
	if err := r.db.QueryRowContext(ctx, query, userID).Scan(&cartID); err != nil {
		return nil, err
	}
	return &Cart{ID: cartID, UserID: userID}, nil
}

func (r *CartRepository) FindOrCreateActiveCart(ctx context.Context, userID int64) (*Cart, error) {
	query := "SELECT * from carts WHERE user_id = $1 and status = 'active' LIMIT 1"

	var cart Cart
	err := r.db.QueryRowContext(ctx, query, userID).Scan(&cart.ID, &cart.UserID, &cart.Status, &cart.Subtotal, &cart.Discount, &cart.Tax, &cart.Total, &cart.CreatedAt, &cart.UpdatedAt)
	if err != nil {
		return nil, err
	}

	if cart.ID == 0 {
		cart, err := r.Create(ctx, userID)
		if err != nil {
			return nil, err
		}
		return cart, nil
	}
	return &cart, nil
}

// UpsertItem adds an item to the cart or updates the quantity if it already exists, also deleting if quantity is zero
func (r *CartRepository) UpsertItem(ctx context.Context, cartID int64, productID int64, quantity int) error {
	query := "INSERT INTO cart_items (cart_id, product_id, quantity) VALUES ($1, $2, $3) ON CONFLICT (cart_id, product_id) DO UPDATE SET quantity = cart_items.quantity + EXCLUDED.quantity"
	_, err := r.db.ExecContext(ctx, query, cartID, productID, quantity)
	if err != nil {
		return err
	}
	return nil
}

// GetItems retrieves all items in the specified cart
func (r *CartRepository) GetItems(ctx context.Context, cartID int64) ([]CartItem, error) {
	query := "SELECT cart_id, product_id, name, description, quantity, snapshot_price, discount_rate, total_price, image_url, added_at, updated_at FROM cart_items WHERE cart_id = $1"
	rows, err := r.db.QueryContext(ctx, query, cartID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var items []CartItem
	for rows.Next() {
		var item CartItem
		if err := rows.Scan(
			&item.CartID,
			&item.ProductID,
			&item.Name,
			&item.Description,
			&item.Quantity,
			&item.SnapshotPrice,
			&item.DiscountRate,
			&item.TotalPrice,
			&item.ImageURL,
			&item.AddedAt,
			&item.UpdatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, item)
	}

	return items, nil
}

// ClearCart removes all items from the specified cart
func (r *CartRepository) ClearCart(ctx context.Context, cartID int64) error {
	query := "DELETE FROM cart_items WHERE cart_id = $1"
	_, err := r.db.ExecContext(ctx, query, cartID)
	if err != nil {
		return err
	}
	return nil
}

// SetCompleted marks the cart as completed
func (r *CartRepository) SetCompleted(ctx context.Context, cartID int64) error {
	query := "UPDATE carts SET status = 'completed' WHERE id = $1"
	_, err := r.db.ExecContext(ctx, query, cartID)
	if err != nil {
		return err
	}

	return nil
}
