package msgutil

import "social-backend/model"

type Data map[string]interface{}

type Message struct {
	Data    Data
}

func NewMessage() Message {
	return Message{
		Data: make(Data),
	}
}

func (m Message) Set(key string, value interface{}) Message {
	m.Data[key] = value
	return m
}

func (m Message) Done() Data {
	return m.Data
}


func RequestBodyParseErrorMessage() Data {
	return NewMessage().
		Set("message", "Failed to parse request body").
		Done()
}

func InternalServerErrorMessage() Data {
	return NewMessage().
		Set("message", "Internal server error").
		Done()
}

func ResourceNotFoundErrorMessage() Data {
	return NewMessage().
		Set("message", "Requested resource not found").
		Done()
}

func LoginSuccessMessage(user model.User) Data {
	return NewMessage().
		Set("message", "Login successful").
		Set("user", user).
		Done()
}

func RegistrationSuccessMessage() Data {
	return NewMessage().
		Set("message", "User registered successfully").
		Done()
}

func LoginFailureMessage() Data {
	return NewMessage().
		Set("message", "Invalid email or password").
		Done()
}

func RegistrationFailureMessage() Data {
	return NewMessage().
		Set("message", "User registration failed").
		Done()
}

func InvalidCredentialsMessage() Data {
	return NewMessage().
		Set("message", "Invalid email or password").
		Done()
}

func UserNotFoundMessage() Data {
	return NewMessage().
		Set("message", "User not found").
		Done()
}