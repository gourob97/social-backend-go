package repository

import (
    "social-backend/internal/config"
    "social-backend/internal/model"
)

func CreateUser(user *model.User) error {
    return config.DB.Create(user).Error
}

func GetUserByEmail(email string) (*model.User, error) {
    var user model.User
    err := config.DB.Where("email = ?", email).First(&user).Error
    return &user, err
}


// 