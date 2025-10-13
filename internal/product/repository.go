package product

import (
	"context"
	"database/sql"
	"fmt"
	"strings"
	"time"
)

type ProductRepository struct {
	db *sql.DB
}

func NewProductRepository(db *sql.DB) *ProductRepository {
	return &ProductRepository{db: db}
}

func (pr *ProductRepository) Create(ctx context.Context, data CreateProductRequest) error {
	query := "INSERT INTO products (name, price, description, stock) VALUES ($1, $2, $3, $4) RETURNING name, price, description, stock"
	_, err := pr.db.ExecContext(ctx, query, data.Name, data.Price, data.Description, data.Stock)
	if err != nil {
		return err
	}

	return nil
}

func (pr *ProductRepository) FindByID(ctx context.Context, id int) (*Product, error) {
	query := "SELECT id, name, price, description, stock FROM products WHERE id = $1"
	row := pr.db.QueryRowContext(ctx, query, id)
	var product Product
	err := row.Scan(&product.ID, &product.Name, &product.Price, &product.Description, &product.Stock)
	if err != nil {
		return nil, err
	}
	return &product, nil
}

func (pr *ProductRepository) FindAll(ctx context.Context) ([]Product, error) {
	query := "SELECT id, name, price, description, stock FROM products"
	rows, err := pr.db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var products []Product

	for rows.Next() {
		var product Product
		if err := rows.Scan(&product.ID, &product.Name, &product.Price, &product.Description, &product.Stock); err != nil {
			return nil, err
		}
		products = append(products, product)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return products, nil
}

func (pr *ProductRepository) Update(ctx context.Context, id int, p UpdateProductRequest) error {
	fields := []string{}
	args := []any{}
	i := 1

	if p.Name != "" {
		fields = append(fields, fmt.Sprintf("name = $%d", i))
		args = append(args, p.Name)
		i++
	}
	if p.Description != nil {
		fields = append(fields, fmt.Sprintf("description = $%d", i))
		args = append(args, p.Description)
		i++
	}
	if p.Price != 0 {
		fields = append(fields, fmt.Sprintf("price = $%d", i))
		args = append(args, p.Price)
		i++
	}
	if p.Stock != nil {
		fields = append(fields, fmt.Sprintf("stock = $%d", i))
		args = append(args, p.Stock)
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
	query := fmt.Sprintf("UPDATE products SET %s WHERE id = $%d",
		strings.Join(fields, ", "),
		i, // placeholder para id
	)

	// Ejecutar
	res, err := pr.db.ExecContext(ctx, query, args...)
	if err != nil {
		return err
	}

	rows, err := res.RowsAffected()
	if err != nil {
		return err
	}
	if rows == 0 {
		return fmt.Errorf("no product updated with id %d", id)
	}

	return nil
}

func (pr *ProductRepository) Delete(ctx context.Context, id int) error {
	query := "DELETE FROM products where id=$1"
	_, err := pr.db.ExecContext(ctx, query, id)
	if err != nil {
		return err
	}

	return nil
}
