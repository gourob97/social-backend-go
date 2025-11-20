package errutil

import "errors"

var (
	ErrUserNotFound       = errors.New("user not found")
	ErrInvalidCredentials = errors.New("invalid credentials")
	ErrUserAlreadyExists  = errors.New("user already exists")
	ErrCreatingUser     = errors.New("error creating user")
	ErrInternalServerError   = errors.New("internal server error")
)