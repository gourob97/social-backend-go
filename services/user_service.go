package services

import (
	"social-backend/model"
	"social-backend/repository/db"
	"social-backend/utils/errutil"

	"golang.org/x/crypto/bcrypt"
)

type UserService struct {
	repository db.Repository
}

func ForNowUserService(repo db.Repository) *UserService {
	return &UserService{
		repository: repo,
	}
}

func (service *UserService) RegisterUser(username, email, password string) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return errutil.ErrInternalServerError
	}
	user := &model.User{
		Username: username,
		Email:    email,
		Password: string(hashedPassword),
	}
	err = service.repository.CreateUser(user)
	if err != nil {
		return errutil.ErrCreatingUser
	}
	return nil
}

func (service *UserService) LoginUser(email, password string) (*model.User, error) {
	user, err := service.repository.GetUserByEmail(email)
	if err != nil {
		return nil, errutil.ErrUserNotFound
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return nil, errutil.ErrInvalidCredentials
	}

	return user, nil
}
