package handler

import (
	"net/http"
	"database/sql"
	"encoding/json"
	"time"
	"fmt"

	"github.com/balagrivine/go_auth/config"
	"github.com/balagrivine/go_auth/internal/database"
	"github.com/go-playground/validator/v10"
)

type CreateUser struct {
	Username string  `json:"username" validate:"required"`
	Email string     `json:"email" validate:"required"`
	Password string  `json:"password" validate:"required"`
	FirstName string `json:"first_name" validate:"required"`
	LastName string  `json:"last_name" validate:"required"`
}

// Handler function to check server health
func HandleHealth(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Will this code work")
	type Response struct {
		message string
	}
	resp := Response{
		message: "StatusOK",
	}
	RespondWithJSON(w, 200, resp)
}

// Handler function to create a new user
func HandleCreateUser(apiCfg *config.APIConfig) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Another test")
		decoder := json.NewDecoder(r.Body)
		var param CreateUser

		// Decode the request body into the param variable
		if err := decoder.Decode(&param); err != nil {
			http.Error(w, fmt.Sprintf("Error parsing JSON %s", err), http.StatusBadRequest)
			return
		}

		// Create a new validator instance
		validate := validator.New(validator.WithRequiredStructEnabled())

		// Validate the decoded user data
		if err := validate.Struct(param); err != nil {
			http.Error(w, fmt.Sprintf("Validation failed %s", err), http.StatusBadRequest)
			return
		}

		// Check for duplicate user
		err := CheckDuplicateUserByEmail(r.Context(), param.Email, apiCfg.DB)
		if err != nil {
			if err == ErrUserExists{
				http.Error(w, "Email already registered", http.StatusConflict)
				return
			}
			http.Error(w, fmt.Sprintf("Error while checking duplicate user: %s", err), http.StatusBadRequest)
			return
		}

		pass, err := HashPassword(param.Password)
		if err != nil {
			http.Error(w, fmt.Sprintf("error hashing password: %s", err), http.StatusInternalServerError)
			return
		}

		currentTime := time.Now().UTC()
		nullTime := sql.NullTime{
			Time: currentTime,
			Valid: true,
		}

		user, err := apiCfg.DB.CreateUser(r.Context(), database.CreateUserParams{
			Username:   param.Username,
			Email:      param.Email,
			Password:   pass,
			FirstName:  param.FirstName,
			LastName:   param.LastName,
			CreatedAt:  nullTime,
			UpdatedAt:  nullTime,
			Verified:   sql.NullBool{Bool: false, Valid: true},
		})
		if err != nil {
			http.Error(w, fmt.Sprintf("Couldn't create user: %s", err), http.StatusBadRequest)
			return
		}

		RespondWithJSON(w, 201, user)
	}
}

type LoginUser struct {
	Email string     `json:"email" validate:"required"`
	Password string  `json:"password" validate:"required"`
}

func HandleLoginUser(apiCfg *config.APIConfig) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		decoder := json.NewDecoder(r.Body)
		var param LoginUser

		// Decode the request body into the param variable
		if err := decoder.Decode(&param); err != nil {
			http.Error(w, fmt.Sprintf("Error while parsing JSON: %s", err), http.StatusBadRequest)
			return
		}

		validate := validator.New()

		if err := validate.Struct(param); err != nil {
			http.Error(w, fmt.Sprintf("Validation failed: %s", err), http.StatusBadRequest)
			return
		}

		user, err := apiCfg.DB.GetUserByEmail(r.Context(), param.Email)
		if err != nil {
			if err == sql.ErrNoRows {
				http.Error(w, fmt.Sprintf("Invalid email or username"), http.StatusUnauthorized)
			}
			http.Error(w, fmt.Sprintf("Error checking user existence: %s", err), http.StatusInternalServerError)
			return
		}

		// Validate provided password against the stored hash
		if !CheckPasswordHash(param.Password, user.Password) {
			http.Error(w, fmt.Sprintf("Invalid email or password"), http.StatusUnauthorized)
			return
		}

		// Create access token for the user
		token, err := createAccessToken(param.Email)
		if err != nil {
			http.Error(w, fmt.Sprintf("Error creating access token: %s", err), http.StatusInternalServerError)
		}

		response := map[string]interface{} {
			"user_id":      user.ID,
			"token_type":   "Bearer",
			"access_token": token,
			"verified":     user.Verified,
		}

		RespondWithJSON(w, 200, response)
	}
}
