package errors

import (
	"net/http"
)

type FieldError struct {
	Field   string `json:"field"`
	Message string `json:"message"`
}
type ApiError struct {
	Code    int          `json:"code"`
	Message string       `json:"message,omitempty"`
	Details []FieldError `json:"details,omitempty"`
}

func NotFoundError(message string) *ApiError {
	return &ApiError{
		Code:    http.StatusNotFound,
		Message: message,
		Details: []FieldError{},
	}
}

func InternalServerError(message string) *ApiError {
	return &ApiError{
		Code:    http.StatusInternalServerError,
		Message: message,
		Details: []FieldError{},
	}
}

func ConfilctError(message string) *ApiError {
	return &ApiError{
		Code:    http.StatusConflict,
		Message: message,
		Details: []FieldError{},
	}
}

func ValidationError(message string, details []FieldError) *ApiError {
	return &ApiError{
		Code:    http.StatusBadRequest,
		Message: message,
		Details: details,
	}
}
