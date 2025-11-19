package db

import (
	"social-backend/model"

	"gorm.io/gorm"
)

type Repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{db: db}
}

func (repo *Repository) CreateUser(user *model.User) error {
	return repo.db.Create(user).Error
}

func (repo *Repository) GetUserByEmail(email string) (*model.User, error) {
	var user model.User
	err := repo.db.Where("email = ?", email).First(&user).Error
	return &user, err
}
