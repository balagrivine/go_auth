-- name: CreateUser :one
INSERT INTO users 
(id, username, email, password, first_name, last_name, created_at, updated_at, verified, reset_password_token, reset_password_token_sent_at)
VALUES (
	$1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11
)
RETURNING id, username, email, password, first_name, last_name, created_at, updated_at, verified, reset_password_token, reset_password_token_sent_at;

-- name: GetUserByEmail :one
SELECT * FROM users WHERE email = $1;

-- name: CreateResetPasswordToken :exec
UPDATE users
SET reset_password_token=$2, reset_password_token_sent_at = $3
WHERE email = $1;

-- name: GetResetPasswordToken :one
SELECT reset_password_token, reset_password_token_sent_at
FROM users
WHERE email = $1;

-- name: ResetPassword :exec
UPDATE users
SET password = $2, reset_password_token = ''
WHERE reset_password_token = $1;
