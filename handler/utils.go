package handler

import (
	"golang.org/x/crypto/bcrypt"
	"github.com/balagrivine/go_auth/internal/database"
	"context"
	"database/sql"
	"errors"
)

var ErrUserExists error = errors.New("user with this email already exists")

// Helper function to hash passwords
func HashPassword(password string) (string, error) {

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashedPassword), nil
}

// Helper function to check if a user is already existent in the database
func CheckDuplicateUserByEmail(ctx context.Context, email string, db *database.Queries) error {
	
	// Get user by email
	_, err := db.GetUserByEmail(ctx, email)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil
		}
		return err
	}
	return ErrUserExists
}
