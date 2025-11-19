package apperrors

import (
	"net/http"
	"strings"
)

// Custom error types
type AppError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Type    string `json:"type"`
}

func (e *AppError) Error() string {
	return e.Message
}

// Predefined error types
var (
	ErrUserNotFound       = &AppError{Code: http.StatusNotFound, Message: "User not found", Type: "NOT_FOUND"}
	ErrPostNotFound       = &AppError{Code: http.StatusNotFound, Message: "Post not found", Type: "NOT_FOUND"}
	ErrInvalidCredentials = &AppError{Code: http.StatusUnauthorized, Message: "Invalid email or password", Type: "INVALID_CREDENTIALS"}
	ErrUsernameExists     = &AppError{Code: http.StatusConflict, Message: "Username already exists", Type: "DUPLICATE_USERNAME"}
	ErrEmailExists        = &AppError{Code: http.StatusConflict, Message: "Email already exists", Type: "DUPLICATE_EMAIL"}
	ErrInternalServer     = &AppError{Code: http.StatusInternalServerError, Message: "Internal server error", Type: "INTERNAL_ERROR"}
	ErrValidation         = &AppError{Code: http.StatusBadRequest, Message: "Validation failed", Type: "VALIDATION_ERROR"}
	ErrInvalidInput       = &AppError{Code: http.StatusBadRequest, Message: "Invalid input", Type: "INVALID_INPUT"}
	ErrUnauthorized       = &AppError{Code: http.StatusUnauthorized, Message: "Unauthorized access", Type: "UNAUTHORIZED"}
	ErrForbidden          = &AppError{Code: http.StatusForbidden, Message: "Access forbidden", Type: "FORBIDDEN"}
	ErrBadRequest         = &AppError{Code: http.StatusBadRequest, Message: "Bad request", Type: "BAD_REQUEST"}
)

// Parse database errors into app errors
func ParseDatabaseError(err error) error {
	errStr := err.Error()
	if strings.Contains(errStr, "Duplicate entry") {
		if strings.Contains(errStr, "username") {
			return ErrUsernameExists
		}
		if strings.Contains(errStr, "email") {
			return ErrEmailExists
		}
	}
	return ErrInternalServer
}

// Parse database errors into app errors with entity context
func ParseDBError(err error, entity string) error {
	errStr := err.Error()
	if strings.Contains(errStr, "Duplicate entry") {
		if strings.Contains(errStr, "username") {
			return ErrUsernameExists
		}
		if strings.Contains(errStr, "email") {
			return ErrEmailExists
		}
	}
	return ErrInternalServer
}

// Check if error is record not found
func IsRecordNotFound(err error) bool {
	return strings.Contains(err.Error(), "record not found")
}

// Create new validation error with custom message
func NewValidationError(message string) error {
	return &AppError{
		Code:    http.StatusBadRequest,
		Message: message,
		Type:    "VALIDATION_ERROR",
	}
}
