package appErrors

import "net/http"

type ErrorResponse struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}

func (e *ErrorResponse) Error() string {
	return e.Message
}

func (e *ErrorResponse) StatusCode() int {
	return e.Status
}

func DefaultBadRequest(message string) ErrorResponse {
	return ErrorResponse{
		Status:  http.StatusBadRequest,
		Message: "Bad request",
	}
}

func BadRequest(message string) ErrorResponse {
	return ErrorResponse{
		Status:  http.StatusBadRequest,
		Message: message,
	}
}

func InternalServerError(message string) ErrorResponse {
	if message == "" {
		message = "Internal server error"
	}
	return ErrorResponse{
		Status:  http.StatusInternalServerError,
		Message: message,
	}
}

func NotFound(message string) ErrorResponse {
	if message == "" {
		message = "Not found"
	}
	return ErrorResponse{
		Status:  http.StatusNotFound,
		Message: message,
	}
}
