package orders

import (
	"context"
	"database/sql"
	"fmt"
	"strings"
	"time"
)

type OrderRepository struct {
	db *sql.DB
}

var ErrOrderNotFound error = fmt.Errorf("Order not found")

func NewOrderRepository(db *sql.DB) *OrderRepository {
	return &OrderRepository{db: db}
}

func (r *OrderRepository) Create(ctx context.Context, o *CreateOrderRequest) error {
	query := "INSERT INTO orders (user_id, cart_id, shipping_address, payment_method) VALUES ($1, $2, $3, $4)"
	_, err := r.db.ExecContext(ctx, query, o.UserID, o.CartID, o.ShippingAddress, o.PaymentMethod)
	return err
}

func (r *OrderRepository) FindByID(ctx context.Context, id int) (*Order, error) {
	query := "SELECT id, user_id, cart_id, shipping_address, payment_method, created_at FROM orders WHERE id = $1"
	row := r.db.QueryRowContext(ctx, query, id)

	var o Order
	if err := row.Scan(&o.ID, &o.UserID, &o.CartID, &o.ShippingAddress, &o.PaymentMethod, &o.CreatedAt); err != nil {
		if err == sql.ErrNoRows {
			return nil, ErrOrderNotFound
		}
		return nil, err
	}
	return &o, nil
}

func (r *OrderRepository) ListByUserID(ctx context.Context, userID int) ([]*Order, error) {
	query := "SELECT id, user_id, cart_id, shipping_address, payment_method, created_at FROM orders WHERE user_id = $1"
	rows, err := r.db.QueryContext(ctx, query, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var orders []*Order
	for rows.Next() {
		var o Order
		if err := rows.Scan(&o.ID, &o.UserID, &o.CartID, &o.ShippingAddress, &o.PaymentMethod, &o.CreatedAt); err != nil {
			return nil, err
		}
		orders = append(orders, &o)
	}
	return orders, nil
}

func (r *OrderRepository) Update(ctx context.Context, id int, o *UpdateOrderRequest) error {
	fields := []string{}
	args := []any{}
	i := 1

	if o.ShippingAddress != nil {
		fields = append(fields, fmt.Sprintf("shipping_address = $%d", i))
		args = append(args, o.ShippingAddress)
		i++
	}

	if o.Status != nil {
		fields = append(fields, fmt.Sprintf("status = $%d", i))
		args = append(args, o.Status)
		i++
	}

	// Siempre actualizamos updated_at
	now := time.Now()
	fields = append(fields, fmt.Sprintf("updated_at = $%d", i))
	args = append(args, now)
	i++

	if len(fields) == 0 {
		return fmt.Errorf("no fields to update")
	}

	// Query final
	query := fmt.Sprintf("UPDATE orders SET %s WHERE id = $%d",
		strings.Join(fields, ", "),
		i, // placeholder para id
	)

	// Ejecutar
	args = append(args, id)
	res, err := r.db.ExecContext(ctx, query, args...)
	if err != nil {
		return err
	}

	rows, err := res.RowsAffected()
	if err != nil {
		return err
	}

	if rows == 0 {
		return fmt.Errorf("no order updated with id %d", id)
	}

	return nil
}

func (r *OrderRepository) Delete(ctx context.Context, id int) error {
	query := "DELETE FROM orders WHERE id = $1"
	_, err := r.db.ExecContext(ctx, query, id)
	return err
}
