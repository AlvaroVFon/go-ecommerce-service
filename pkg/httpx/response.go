// Package httpx provides extended HTTP functionalities.
package httpx

import (
	"encoding/json"
	"log"
	"net/http"
)

const (
	OkResponse        = "Request processed successfully."
	NoContentResponse = "No content to display."
	CreatedResponse   = "Resource created successfully."
	UpdatedResponse   = "Resource updated successfully."
	DeletedResponse   = "Resource deleted successfully."
)

func HTTPResponse(w http.ResponseWriter, status int, data any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	if err := json.NewEncoder(w).Encode(data); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func HTTPPaginatedResponse(w http.ResponseWriter, status int, data any, page, limit, total int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	if err := json.NewEncoder(w).Encode(map[string]any{
		"data":        data,
		"total":       total,
		"page":        page,
		"total_pages": (total + limit - 1) / limit,
	}); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func ParseJSON(r *http.Request, dst any) error {
	defer func() {
		if err := r.Body.Close(); err != nil {
			log.Println("Error closing request body:", err)
		}
	}()
	return json.NewDecoder(r.Body).Decode(&dst)
}
