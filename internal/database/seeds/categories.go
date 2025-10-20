package seeds

import (
	"database/sql"
	"ecommerce-service/internal/categories"
)

func SeedCategories(db *sql.DB) error {
	categories := []categories.Category{
		{Name: "Electronics", Description: "Devices and gadgets"},
		{Name: "Books", Description: "Printed and digital books"},
		{Name: "Clothing", Description: "Apparel and accessories"},
		{Name: "Home & Kitchen", Description: "Household items and kitchenware"},
		{Name: "Sports & Outdoors", Description: "Sporting goods and outdoor equipment"},
	}
	query := "INSERT INTO categories (name, description) VALUES ($1, $2) ON CONFLICT (name) DO NOTHING"

	for _, category := range categories {
		if _, err := db.Exec(query, category.Name, category.Description); err != nil {
			return err
		}
	}
	return nil
}
