-- name: CreateCategory :one
INSERT INTO categories (name)
VALUES ($1)
RETURNING *;

-- name: ListCategories :many
SELECT *
FROM categories
LIMIT $1 OFFSET $2;

-- name: GetCategoryById :one
SELECT *
FROM categories
WHERE id = $1 AND deleted_at IS NULL
LIMIT 1;

-- name: UpdateCategory :one
UPDATE categories
SET name = $2,
    updated_at = NOW()
WHERE id = $1 AND deleted_at IS NULL
RETURNING *;

-- name: DeleteCategory :exec
UPDATE categories
SET deleted_at = NOW()
WHERE id = $1;
