package dto

// User response DTOs - safe for API responses
type UserResponse struct {
	ID       uint   `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
}

// Authentication response DTOs
type LoginResponse struct {
	Message string       `json:"message"`
	User    UserResponse `json:"user"`
	Token   string       `json:"token"`
}

type RegisterResponse struct {
	Message string `json:"message"`
	Success bool   `json:"success"`
}

// Standard API response wrapper
type APIResponse struct {
	Success bool        `json:"success"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
	Error   string      `json:"error,omitempty"`
}
