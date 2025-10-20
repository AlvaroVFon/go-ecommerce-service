package categories

import (
	"context"
	"database/sql"
)

type CategoryRepository struct {
	db *sql.DB
}

func NewCategoryRepository(db *sql.DB) *CategoryRepository {
	return &CategoryRepository{db: db}
}

func (r *CategoryRepository) FindAll(ctx context.Context) ([]Category, error) {
	query := "SELECT id, name, description FROM categories"

	rows, err := r.db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}

	var categories []Category

	for rows.Next() {
		var category Category
		if err := rows.Scan(&category.ID, &category.Name, &category.Description); err != nil {
			return nil, err
		}
		categories = append(categories, category)
	}

	return categories, nil
}

func (r *CategoryRepository) FindByID(ctx context.Context, id int) (*Category, error) {
	query := "SELECT id, name, description FROM categories WHERE id = $1"
	row := r.db.QueryRowContext(ctx, query, id)
	var category Category

	if err := row.Scan(&category.ID, &category.Name, &category.Description); err != nil {
		return nil, err
	}

	return &category, nil
}

func (r *CategoryRepository) FindByName(ctx context.Context, name string) (*Category, error) {
	query := "SELECT id, name, description FROM categories WHERE name = $1"
	row := r.db.QueryRowContext(ctx, query, name)
	var category Category
	if err := row.Scan(&category.ID, &category.Name, &category.Description); err != nil {
		return nil, err
	}

	return &category, nil
}
