package service

import (
	"social-backend/internal/auth"
	"social-backend/internal/dto"
	"social-backend/internal/model"
	"social-backend/internal/repository/interfaces"
	serviceInterfaces "social-backend/internal/service/interfaces"

	"golang.org/x/crypto/bcrypt"
)

type userService struct {
	userRepo   interfaces.UserRepository
	jwtManager *auth.JWTManager
}

func NewUserService(userRepo interfaces.UserRepository, jwtSecret string) serviceInterfaces.UserService {
	return &userService{
		userRepo:   userRepo,
		jwtManager: auth.NewJWTManager(jwtSecret),
	}
}

func (s *userService) RegisterUser(username, email, password string) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	user := &model.User{
		Username: username,
		Email:    email,
		Password: string(hashedPassword),
	}
	return s.userRepo.CreateUser(user)
}

func (s *userService) LoginUser(email, password string) (*dto.LoginResponse, error) {
	user, err := s.userRepo.GetUserByEmail(email)
	if err != nil {
		return nil, err
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return nil, err
	}

	// Generate JWT token
	token, err := s.jwtManager.GenerateToken(user.ID, user.Username, user.Email)
	if err != nil {
		return nil, err
	}

	// Return safe user data with token
	return &dto.LoginResponse{
		Message: "Login successful",
		User: dto.UserResponse{
			ID:       user.ID,
			Username: user.Username,
			Email:    user.Email,
		},
		Token: token,
	}, nil
}

func (s *userService) GetUserByID(id uint) (*dto.UserResponse, error) {
	user, err := s.userRepo.GetUserByID(id)
	if err != nil {
		return nil, err
	}

	return &dto.UserResponse{
		ID:       user.ID,
		Username: user.Username,
		Email:    user.Email,
	}, nil
}
