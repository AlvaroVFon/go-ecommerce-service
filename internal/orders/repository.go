package orders

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"strings"
	"time"
)

type OrderRepository struct {
	db *sql.DB
}

func NewOrderRepository(db *sql.DB) *OrderRepository {
	return &OrderRepository{db: db}
}

func (r *OrderRepository) Create(ctx context.Context, order *Order) (*Order, error) {
	// Start a transaction
	tx, err := r.db.BeginTx(ctx, nil)
	if err != nil {
		return nil, fmt.Errorf("error starting transaction: %w", err)
	}
	// Defer a rollback only if an error occurs
	defer func() {
		if err != nil {
			tx.Rollback()
		}
	}()

	// 1. Insert into orders table and get the new order ID
	orderQuery := "INSERT INTO orders (user_id, total, status, shipping_address, payment_method) VALUES ($1, $2, $3, $4, $5) RETURNING id, created_at"
	err = tx.QueryRowContext(ctx, orderQuery, order.UserID, order.Total, order.Status, order.ShippingAddress, order.PaymentMethod).Scan(&order.ID, &order.CreatedAt)
	if err != nil {
		return nil, fmt.Errorf("error inserting order: %w", err)
	}

	// 2. Prepare statement for inserting order items
	itemStmt, err := tx.PrepareContext(ctx, "INSERT INTO order_items (order_id, product_id, quantity, price) VALUES ($1, $2, $3, $4)")
	if err != nil {
		return nil, fmt.Errorf("error preparing order item statement: %w", err)
	}
	defer itemStmt.Close()

	// 3. Insert all order items
	for i, item := range order.Items {
		_, err := itemStmt.ExecContext(ctx, order.ID, item.ProductID, item.Quantity, item.Price)
		if err != nil {
			return nil, fmt.Errorf("error inserting order item #%d: %w", i+1, err)
		}
	}

	// 4. Commit the transaction
	if err := tx.Commit(); err != nil {
		return nil, fmt.Errorf("error committing transaction: %w", err)
	}

	return order, nil
}

func (r *OrderRepository) FindByID(ctx context.Context, id int) (*Order, error) {
	query := "SELECT id, user_id, shipping_address, payment_method, created_at FROM orders WHERE id = $1"
	row := r.db.QueryRowContext(ctx, query, id)

	var o Order
	if err := row.Scan(&o.ID, &o.UserID, &o.ShippingAddress, &o.PaymentMethod, &o.CreatedAt); err != nil {
		if err == sql.ErrNoRows {
			return nil, err
		}
		return nil, err
	}
	return &o, nil
}

func (r *OrderRepository) ListByUserID(ctx context.Context, userID, limit, offset int) ([]*Order, error) {
	query := "SELECT id, user_id, shipping_address, payment_method, created_at FROM orders WHERE user_id = $1 LIMIT $2 OFFSET $3"
	rows, err := r.db.QueryContext(ctx, query, userID, limit, offset)
	if err != nil {
		return nil, err
	}
	defer func() {
		if err := rows.Close(); err != nil {
			log.Fatal("Eror closing rows")
		}
	}()

	var orders []*Order
	for rows.Next() {
		var o Order
		if err := rows.Scan(&o.ID, &o.UserID, &o.ShippingAddress, &o.PaymentMethod, &o.CreatedAt); err != nil {
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

func (r *OrderRepository) CountByUserID(ctx context.Context, userID int) (int, error) {
	query := "SELECT COUNT(*) FROM orders WHERE user_id = $1"
	row := r.db.QueryRowContext(ctx, query, userID)

	var count int
	if err := row.Scan(&count); err != nil {
		return 0, err
	}
	return count, nil
}
