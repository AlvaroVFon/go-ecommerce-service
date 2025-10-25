// Package httpx provides extended HTTP functionalities.
package httpx

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/go-playground/validator/v10"
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

func HTTPError(w http.ResponseWriter, status int, message string) {
	HTTPResponse(w, status, map[string]string{"error": message})
}

func HTTPErrors(w http.ResponseWriter, status int, messages map[string]string) {
	HTTPResponse(w, status, map[string]map[string]string{"errors": messages})
}

func FormatValidatorErrors(err error) map[string]string {
	errors := make(map[string]string)
	for _, err := range err.(validator.ValidationErrors) {
		field := strings.ToLower(err.Field())
		var message string
		switch err.Tag() {
		case "required":
			message = fmt.Sprintf("the %s field is required", field)
		case "email":
			message = fmt.Sprintf("the %s field must be a valid email address", field)
		case "min":
			message = fmt.Sprintf("the %s field must be at least %s characters long", field, err.Param())
		case "max":
			message = fmt.Sprintf("the %s field must be at most %s characters long", field, err.Param())
		default:
			message = fmt.Sprintf("the %s field is invalid", field)
		}
		errors[field] = message
	}
	return errors
}
