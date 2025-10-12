// Package product defines the data structures and types related to products in the system.
package product

import "time"

type Product struct {
	ID          int32      `json:"id"`
	Name        string     `json:"name"`
	Description *string    `json:"description,omitempty"`
	Price       float64    `json:"price"`
	Stock       *int32     `json:"stock,omitempty"`
	CreatedAt   *time.Time `json:"created_at,omitempty"`
	UpdatedAt   *time.Time `json:"updated_at,omitempty"`
}
