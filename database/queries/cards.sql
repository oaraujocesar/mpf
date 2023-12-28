-- name: CreateCard :one
INSERT INTO cards (name, card_limit, due_date, user_id, family_id)
VALUES ($1, $2, $3, $4, $5)
RETURNING *;

-- name: ListCards :many
SELECT *
FROM cards
ORDER BY $1
LIMIT $2 OFFSET $3;

-- name: GetCardById :one
SELECT *
FROM cards
WHERE id = $1 AND deleted_at IS NULL
LIMIT 1;

-- name: UpdateCard :one
UPDATE cards
SET name = $2,
    card_limit = $3,
    due_date = $4,
    updated_at = NOW()
WHERE id = $1 AND deleted_at IS NULL
RETURNING *;

-- name: DeleteCard :exec
UPDATE cards
SET deleted_at = NOW()
WHERE id = $1;