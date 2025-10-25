package httpx

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/go-playground/validator/v10"
)

const (
	NotFoundError       = "the requested resource was not found"
	UnauthorizedError   = "you are not authorized to access this resource"
	ForbiddenError      = "access to this resource is forbidden"
	InternalServerError = "an internal server error occurred"
	ConflictError       = "a conflict occurred with the current state of the resource"
	BadRequestError     = "the request could not be understood or was missing required parameters"
	InvalidIDError      = "the provided ID is invalid"
)

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
