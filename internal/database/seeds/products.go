package seeds

import (
	"database/sql"

	"ecommerce-service/internal/products"
)

func SeedProducts(db *sql.DB) error {
	description1 := "A high-performance laptop"
	description2 := "A latest model smartphone"
	description3 := "A programmable coffee maker"
	description4 := "A comfortable pair of running shoes"
	description5 := "A best-selling novel"

	products := []products.Product{
		{Name: "Laptop", Description: description1, Price: 1200.00, Stock: 12},
		{Name: "Smartphone", Description: description2, Price: 800.00, Stock: 22},
		{Name: "Coffee Maker", Description: description3, Price: 150.00, Stock: 32},
		{Name: "Running Shoes", Description: description4, Price: 120.00, Stock: 55},
		{Name: "Novel", Description: description5, Price: 20.00, Stock: 65},
	}

	// Initialize stock values
	products[0].Stock = 50
	products[1].Stock = 200
	products[2].Stock = 75
	products[3].Stock = 150
	products[4].Stock = 300

	query := "INSERT INTO products (name, description, price, stock) VALUES ($1, $2, $3, $4) ON CONFLICT (name) DO NOTHING"

	for _, product := range products {
		if _, err := db.Exec(query, product.Name, product.Description, product.Price, product.Stock); err != nil {
			return err
		}
	}

	return nil
}
