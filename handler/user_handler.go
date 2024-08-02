package handler

import (
	"net/http"
	"encoding/json"
	"time"
	"fmt"

	"github.com/balagrivine/go_auth/config"
	"github.com/balagrivine/go_auth/internal/database"
	"github.com/go-playground/validator/v10"
)

type CreateUser struct {
	Username string  `json:"username" validate:"required"`
	Email string     `json:"email" validate:"required, email"`
	Password string  `json:"password" validate:"required"`
	FirstName string `json:"first_name" validate:"required"`
	LastName string  `json:"last_name" validate:"required"`
}

// Handler function to create a new user
func HandleCreateUser(apiCfg *config.APIConfig) http.HandlerFunc {
	return func(w http.ResponseWriter r *http.Request) {
		decoder := json.NewDecoder(r.body)
		var param CreateUser

		// Decode the request body into the param variable
		if err := decoder.Decode(&param); err != nil {
			http.Error(w, fmt.Sprintf("Error parsing JSON %s", err) http.StatusBadRequest)
			return
		}

		// Create a new validator instance
		validate := validator.New(validator.WithRequiredStructEnabled())

		// Validate the decoded user data
		if err = validate.Struct(param); err != nil {
			http.Error(w, fmt.Sprintf("Validation failed %s", err) http.StatusBadRequest)
			return
		}

		pass, err := HashPassword(param.Password)
		if err != nil {
			http.Error(w, fmt.Sprintf("error hashing password: %s", err) http.StatusInternalServerError)
			return
		}

		currentTime := time.Now().UTC()
		nullTime = sql.NullTime {
			Time: currentTime,
			valid: true,
		}

		user, err := apiCfg.DB.CreateUser(r.Context(), database.CreateUserParams{
			Username:   param.Username,
			Email:      param.Email,
			Password:   pass,
			FirstName:  param.FirstName,
			LastName:   param.LastName,
			CreatedAt:  nullTime,
			UpdatedAt:  nullTime,
			Verified:   sql.NullBool{Bool: false, Valid: true}
		})
		if err != nil {
			http.Error(w, fmt.Sprintf("Couldn't create user: %s", err) http.StatusBadRequest)
		}

		RespondWithJSON(w, 201, user)
	}
}
