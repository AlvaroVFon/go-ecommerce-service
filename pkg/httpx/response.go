// Package httpx provides extended HTTP functionalities.
package httpx

import (
	"encoding/json"
	"log"
	"net/http"
)

func JSON(w http.ResponseWriter, status int, data any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	if err := json.NewEncoder(w).Encode(data); err != nil {
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

func Error(w http.ResponseWriter, status int, message string) {
	JSON(w, status, map[string]string{"error": message})
}
