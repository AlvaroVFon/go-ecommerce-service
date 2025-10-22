package seeds

import (
	"database/sql"
)

type ProductCategory struct {
	ProductID  int
	CategoryID int
}

func SeedProductCategory(db *sql.DB) error {
	productCategories := []ProductCategory{
		{ProductID: 1, CategoryID: 1}, // Laptop -> Electronics
		{ProductID: 2, CategoryID: 1}, // Smartphone -> Electronics
		{ProductID: 3, CategoryID: 4}, // Coffee Maker -> Home & Kitchen
		{ProductID: 4, CategoryID: 3}, // Running Shoes -> Clothing
		{ProductID: 4, CategoryID: 5}, // Running Shoes -> Sports & Outdoors
		{ProductID: 5, CategoryID: 2}, // Novel -> Books
	}

	query := "INSERT INTO product_category (product_id, category_id) VALUES ($1, $2) ON CONFLICT (product_id, category_id) DO NOTHING"

	for _, pc := range productCategories {
		if _, err := db.Exec(query, pc.ProductID, pc.CategoryID); err != nil {
			return err
		}
	}

	return nil
}
