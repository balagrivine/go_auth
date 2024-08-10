package handler

import (
	"time"
	"os"
	"fmt"
	"errors"
	"context"
	"database/sql"

	"golang.org/x/crypto/bcrypt"
	"github.com/dgrijalva/jwt-go"
	"github.com/joho/godotenv"

	"github.com/balagrivine/go_auth/internal/database"
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

func CheckPasswordHash(password string, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

// Function to generate access tokens
func createAccessToken(email string) (string, error) {
	err := godotenv.Load(".env")
	if err != nil {
		return "", err
	}

	// Define token claims
	claims := jwt.MapClaims{
		"sub": email,
		"exp": time.Now().Add(time.Hour * 72).Unix(),
	}

	// Create new JWT token with claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	secretKey := os.Getenv("SECRET_KEY")
	if secretKey == "" {
		return "", fmt.Errorf("SECRET_KEY is not set in the environment variables")
	}

	// Sign the token with the secret key
	tokenString, err := token.SignedString([]byte(secretKey))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
