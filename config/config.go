package config

import (
	"database/sql"
	"os"

	"github.com/balagrivine/go_auth/internal/database"
	_"github.com/lib/pq"


)

type APIConfig struct {
	DB *database.Queries
}

// Initializes a database connection
func InitConfig() (*APIConfig, error) {

	dbURL := os.Getenv("DB_URL")

	// Make a database connection
	conn, err := sql.Open("postgres", dbURL)
	if err != nil {
		return nil, err
	}

	if err = conn.Ping(); err != nil {
		return nil, err
	}

	queries := database.New(conn)

	return &APIConfig{
		DB: queries,
	}, nil
}
