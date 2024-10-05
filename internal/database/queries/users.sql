-- name: GetUsers :many
SELECT * FROM users;

-- name: CreateUser :one
INSERT INTO users (name, surname, email, password)
VALUES ($1, $2, $3, $4)
RETURNING *;
