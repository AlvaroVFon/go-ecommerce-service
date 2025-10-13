// Package product defines the data structures and types related to products in the system.
package product

import (
	"time"
)

type Product struct {
	ID          int32      `json:"id"`
	Name        string     `json:"name"`
	Description *string    `json:"description,omitempty"`
	Price       float64    `json:"price"`
	Stock       *int32     `json:"stock,omitempty"`
	CreatedAt   *time.Time `json:"created_at,omitempty"`
	UpdatedAt   *time.Time `json:"updated_at,omitempty"`
}

type CreateProductRequest struct {
	Name        string  `json:"name" validate:"required,min=3"`
	Description string  `json:"description" validate:"required"`
	Price       float64 `json:"price" validate:"required,gt=0"`
	Stock       int     `json:"stock" validate:"required,gte=0"`
}

type UpdateProductRequest struct {
	Name        *string  `json:"name"`
	Description *string  `json:"description,omitempty"`
	Price       *float64 `json:"price"`
	Stock       *int     `json:"stock,omitempty"`
}
