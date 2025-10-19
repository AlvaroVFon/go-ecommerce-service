// Package categories defines the data structures and types related to categories in the application.
package categories

type Category struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}
