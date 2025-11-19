package interfaces

import "social-backend/internal/model"

type UserRepository interface {
	CreateUser(user *model.User) error
	GetUserByEmail(email string) (*model.User, error)
	GetUserByID(id uint) (*model.User, error)
	UpdateUser(user *model.User) error
	DeleteUser(id uint) error
}
