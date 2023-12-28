-- name: CreateFamily :one
INSERT INTO families (name, user_id)
VALUES ($1, $2)
RETURNING *;

-- name: ListFamilies :many
SELECT *
FROM families
ORDER BY $1
LIMIT $2 OFFSET $3;

-- name: GetFamilyById :one
SELECT *
FROM families
WHERE id = $1 AND deleted_at IS NULL
LIMIT 1;

-- name: UpdateFamily :one
UPDATE families
SET name = $2,
    updated_at = NOW()
WHERE id = $1 AND deleted_at IS NULL
RETURNING *;

-- name: DeleteFamily :exec
UPDATE families
SET deleted_at = NOW()
WHERE id = $1;