package interfaces

import "social-backend/internal/dto"

type UserService interface {
	RegisterUser(username, email, password string) error
	LoginUser(email, password string) (*dto.LoginResponse, error)
	GetUserByID(id uint) (*dto.UserResponse, error)
}
