package response

import (
	"social-backend/internal/apperrors"

	"github.com/labstack/echo/v4"
)

// Standard API response structure
type APIResponse struct {
	Success bool        `json:"success"`
	Message string      `json:"message,omitempty"`
	Data    interface{} `json:"data,omitempty"`
	Error   interface{} `json:"error,omitempty"`
	Meta    interface{} `json:"meta,omitempty"`
}

// Success responses
func Success(c echo.Context, statusCode int, message string, data interface{}) error {
	return c.JSON(statusCode, APIResponse{
		Success: true,
		Message: message,
		Data:    data,
	})
}

func Created(c echo.Context, message string, data interface{}) error {
	return Success(c, 201, message, data)
}

func OK(c echo.Context, message string, data interface{}) error {
	return Success(c, 200, message, data)
}

// Error responses
func Error(c echo.Context, err error) error {
	if appErr, ok := err.(*apperrors.AppError); ok {
		return c.JSON(appErr.Code, APIResponse{
			Success: false,
			Error:   appErr.Message,
		})
	}

	// Default to internal server error for unknown errors
	return c.JSON(500, APIResponse{
		Success: false,
		Error:   "Internal server error",
	})
}

func BadRequest(c echo.Context, message string) error {
	return c.JSON(400, APIResponse{
		Success: false,
		Error:   message,
	})
}

func Unauthorized(c echo.Context, message string) error {
	return c.JSON(401, APIResponse{
		Success: false,
		Error:   message,
	})
}

func Forbidden(c echo.Context, message string) error {
	return c.JSON(403, APIResponse{
		Success: false,
		Error:   message,
	})
}

func NotFound(c echo.Context, message string) error {
	return c.JSON(404, APIResponse{
		Success: false,
		Error:   message,
	})
}

// Convenience functions to match controller usage
func SuccessResponse(c echo.Context, statusCode int, message string, data interface{}) error {
	return Success(c, statusCode, message, data)
}

func ErrorResponse(c echo.Context, err error) error {
	return Error(c, err)
}
