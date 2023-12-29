-- name: CreateUser :one
INSERT INTO users (name, email, password, avatar)
VALUES ($1, $2, $3, $4)
RETURNING *;

-- name: ListUsers :many
SELECT *
FROM users
ORDER BY $1
LIMIT $2 OFFSET $3;

-- name: GetUserByEmail :one
SELECT *
FROM users
WHERE email = $1 AND deleted_at IS NULL
LIMIT 1;

-- name: GetUserById :one
SELECT *
FROM users
WHERE id = $1 AND deleted_at IS NULL
LIMIT 1;

-- name: UpdateUser :one
UPDATE users
SET name = $2,
    email = $3,
    password = $4,
    avatar = $5,
    updated_at = NOW()
WHERE id = $1 AND deleted_at IS NULL
RETURNING *;

-- name: DeleteUser :exec
UPDATE users
SET deleted_at = NOW()
WHERE id = $1;

-- name: HardDeleteUser :exec
DELETE FROM users
WHERE id = $1;
