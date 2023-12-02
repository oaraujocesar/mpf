-- name: CreateAuthor :one
INSERT INTO users (name, avatar, email, password)
VALUES ($1, $2, $3, $4)
RETURNING *;

-- name: GetAuthor :one
SELECT *
FROM users
WHERE email = $1
LIMIT 1;

-- name: ListAuthors :many
SELECT *
FROM users
ORDER BY $1
LIMIT $2 OFFSET $3;

-- name: UpdateAuthor :one
UPDATE users
SET name = $2,
    avatar = $3,
    email = $4,
    password = $5,
    updated_at = NOW()
WHERE id = $1
RETURNING *;

-- name: DeleteAuthor :exec
DELETE FROM users
WHERE id = $1;

-- name: CreateCategory :one
INSERT INTO categories (name)
VALUES ($1) 
RETURNING *;

-- name: GetCategory :one
SELECT *
FROM categories
WHERE name = $1
LIMIT 1;

-- name: ListCategories :many
SELECT *
FROM categories
ORDER BY $1
LIMIT $2 OFFSET $3;

-- name: UpdateCategory :one
UPDATE categories
SET name = $2,
    updated_at = NOW()
WHERE id = $1
RETURNING *;

-- name: DeleteCategory :exec
DELETE FROM categories
WHERE id = $1;

-- name: CreateEntry :one
INSERT INTO entries (title, amount, type, user_id, category_id)
VALUES ($1, $2, $3, $4, $5)
RETURNING *;

-- name: GetEntry :one
SELECT *
FROM entries
WHERE id = $1
LIMIT 1;

-- name: ListEntries :many
SELECT *
FROM entries
ORDER BY $1
LIMIT $2 OFFSET $3;

-- name: UpdateEntry :one
UPDATE entries
SET title = $2,
    amount = $3,
    type = $4,
    user_id = $5,
    category_id = $6,
    updated_at = NOW()
WHERE id = $1
RETURNING *;

-- name: DeleteEntry :exec
DELETE FROM entries
WHERE id = $1;