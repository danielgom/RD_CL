// Package validator used to make validations on structs
package validator

import "C"

import (
	"fmt"
	"strings"
	"unicode"

	"github.com/go-playground/validator/v10"
)

// CustomValidator sets the validator for struct fields.
type CustomValidator struct {
	*validator.Validate
}

// GetValidator gets a new validator instance.
func GetValidator() *CustomValidator {
	return &CustomValidator{validator.New()}
}

// AddValidators registers all validators for struct fields.
func AddValidators(v *validator.Validate) error {
	err := v.RegisterValidation("password", ValidatePassword)
	if err != nil {
		return fmt.Errorf("validator register error: %w", err)
	}

	return nil
}

// ValidatePassword is a custom password validator.
func ValidatePassword(fl validator.FieldLevel) bool {
	password := fl.Field().String()
	var hasSpecial bool

	hasNumber := strings.ContainsAny(password, "123456789")
	hasUpper := strings.ContainsAny(password, "ABCDEFGHIJKLMNOPQRSTVWXYZ")
	hasLower := strings.ContainsAny(password, "abcdefghijklmnopqrstvwxyz")

	for _, char := range password {
		if unicode.IsPunct(char) || unicode.IsSymbol(char) {
			hasSpecial = true
		}
	}
	return hasUpper && hasLower && hasNumber && hasSpecial
}
