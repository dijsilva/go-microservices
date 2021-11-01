package appErrors

import (
	"net/http"
)

type ErrorResponse struct {
	Status  int         `json:"status"`
	Message string      `json:"message"`
	Details interface{} `json:"details,omitempty"`
}

func (e ErrorResponse) Error() string {
	return e.Message
}

func (e ErrorResponse) StatusCode() int {
	return e.Status
}

func Unauthorized(message string) ErrorResponse {
	if message == "" {
		message = "You are not authenticated to perform the requested action."
	}
	return ErrorResponse{
		Status:  http.StatusUnauthorized,
		Message: message,
	}
}

func IncorrectCredentials(message string) ErrorResponse {
	if message == "" {
		message = "Incorrect credentials."
	}
	return ErrorResponse{
		Status:  http.StatusUnauthorized,
		Message: message,
	}
}

func NotFound(message string) ErrorResponse {
	if message == "" {
		message = "Resource not found."
	}
	return ErrorResponse{
		Status:  http.StatusNotFound,
		Message: message,
	}
}

func AccessDenied(message string) ErrorResponse {
	if message == "" {
		message = "Access denied."
	}
	return ErrorResponse{
		Status:  http.StatusUnauthorized,
		Message: message,
	}
}

func BadInput(message string) ErrorResponse {
	if message == "" {
		message = "Bad input."
	}
	return ErrorResponse{
		Status:  http.StatusBadRequest,
		Message: message,
	}
}

func AlreadyExists(message string) ErrorResponse {
	if message == "" {
		message = "Resource already exists."
	}
	return ErrorResponse{
		Status:  http.StatusBadRequest,
		Message: message,
	}
}

func InternalServerError(message string) ErrorResponse {
	if message == "" {
		message = "Internal Server Error"
	}
	return ErrorResponse{
		Status:  http.StatusInternalServerError,
		Message: message,
	}
}
