module github.com/balagrivine/go_auth

go 1.22.5

require (
	github.com/balagrivine/go_auth/config v0.0.0-20240803143509-5ad1d2301acc
	github.com/balagrivine/go_auth/handler v0.0.0-20240802190103-d7b63a39bfe8
	github.com/joho/godotenv v1.5.1
)

require (
	github.com/balagrivine/go_auth/internal/database v0.0.0-20240803143509-5ad1d2301acc // indirect
	github.com/dgrijalva/jwt-go v3.2.0+incompatible // indirect
	github.com/gabriel-vasile/mimetype v1.4.3 // indirect
	github.com/go-playground/locales v0.14.1 // indirect
	github.com/go-playground/universal-translator v0.18.1 // indirect
	github.com/go-playground/validator/v10 v10.22.0 // indirect
	github.com/leodido/go-urn v1.4.0 // indirect
	github.com/lib/pq v1.10.9 // indirect
	golang.org/x/crypto v0.25.0 // indirect
	golang.org/x/net v0.25.0 // indirect
	golang.org/x/sys v0.22.0 // indirect
	golang.org/x/text v0.16.0 // indirect
)

replace github.com/balagrivine/go_auth/handler => ./handler

replace github.com/balagrivine/go_auth/config => ./config
