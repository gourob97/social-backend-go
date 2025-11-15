package service

import (
    "golang.org/x/crypto/bcrypt"
    "social-backend/internal/model"
    "social-backend/internal/repository"
)

func RegisterUser(username, email, password string) error {
    hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
    if err != nil {
        return err
    }
    user := &model.User{
        Username: username,
        Email:    email,
        Password: string(hashedPassword),
    }
    return repository.CreateUser(user)
}

func LoginUser(email, password string) (*model.User, error) {
    user, err := repository.GetUserByEmail(email)
    if err != nil {
        return nil, err
    }

    err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
    if err != nil {
        return nil, err
    }

    return user, nil
}