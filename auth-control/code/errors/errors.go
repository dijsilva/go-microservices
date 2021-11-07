package appErrors

import (
	"net/http"
)

type ErrorResponse struct {
	Status  int         `json:"status"`
	Message string      `json:"message"`
	Details interface{} `json:"details,omitempty"`
}

// mandatory to default error interface
func (e ErrorResponse) Error() string {
	return e.Message
}

func (e ErrorResponse) StatusCode() int {
	return e.Status
}

func DefaultBadRequest(message string) ErrorResponse {
	return ErrorResponse{
		Status:  http.StatusUnauthorized,
		Message: "Bad request",
	}
}

func InternalServerError(message string) ErrorResponse {
	return ErrorResponse{
		Status:  http.StatusInternalServerError,
		Message: "Internal server error",
	}
}

func DatabaseOperationError(message string) ErrorResponse {
	return ErrorResponse{
		Status:  http.StatusInternalServerError,
		Message: "Internal server error - Database operation",
	}
}

func Unauthorized(message string) ErrorResponse {
	return ErrorResponse{
		Status:  http.StatusUnauthorized,
		Message: message,
	}
}
