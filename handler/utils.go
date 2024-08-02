package handler

import (
	"golang.org/x/crypto/bcrypt"
)

// Helper function to hash passwords
func HashPassword(password string) (string, error) {
	if hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost); err != nil {
		return "", err
	}
	return string(hashedPassword), nil
}
