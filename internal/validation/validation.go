package validation

import (
	"fmt"
	"strings"

	"github.com/go-playground/validator/v10"
)

type Validator struct {
	validator *validator.Validate
}

func NewValidator() *Validator {
	return &Validator{
		validator: validator.New(),
	}
}

func (v *Validator) Validate(i interface{}) error {
	if err := v.validator.Struct(i); err != nil {
		return fmt.Errorf("validation failed: %s", v.formatValidationErrors(err))
	}
	return nil
}

func (v *Validator) formatValidationErrors(err error) string {
	var errors []string

	for _, err := range err.(validator.ValidationErrors) {
		fieldName := strings.ToLower(err.Field())
		switch err.Tag() {
		case "required":
			errors = append(errors, fmt.Sprintf("%s is required", fieldName))
		case "email":
			errors = append(errors, fmt.Sprintf("%s must be a valid email", fieldName))
		case "min":
			errors = append(errors, fmt.Sprintf("%s must be at least %s characters", fieldName, err.Param()))
		case "max":
			errors = append(errors, fmt.Sprintf("%s must not exceed %s characters", fieldName, err.Param()))
		default:
			errors = append(errors, fmt.Sprintf("%s is invalid", fieldName))
		}
	}

	return strings.Join(errors, ", ")
}
