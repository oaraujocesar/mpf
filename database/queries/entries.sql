-- name: CreateEntry :one
INSERT INTO entries (title, amount, account_id, installments, type, category_id, invoice_id, payday, paid_at)
VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)
RETURNING *;

-- name: ListEntries :many
SELECT *
FROM entries
ORDER BY $1
LIMIT $2 OFFSET $3;

-- name: GetEntryById :one
SELECT *
FROM entries
WHERE id = $1 AND deleted_at IS NULL
LIMIT 1;

-- name: UpdatePaidAt :one
UPDATE entries
SET paid_at = $2,
    updated_at = NOW()
WHERE id = $1 AND deleted_at IS NULL
RETURNING *;

-- name: UpdateEntry :one
UPDATE entries
SET title = $2,
    amount = $3,
    category_id = $4,
    payday = $5,
    updated_at = NOW()
WHERE id = $1 AND deleted_at IS NULL
RETURNING *;

-- name: DeleteEntry :exec
UPDATE entries
SET deleted_at = NOW()
WHERE id = $1;