-- name: CreateInstallment :one
INSERT INTO installments (amount, entry_id, payday, paid_at)
VALUES ($1, $2, $3, $4)
RETURNING *;

-- name: ListInstallments :many
SELECT *
FROM installments
WHERE entry_id = $1
ORDER BY $2
LIMIT $3 OFFSET $4;

-- name: GetInstallmentById :one
SELECT *
FROM installments
WHERE id = $1 AND deleted_at IS NULL
LIMIT 1;

-- name: UpdateInstallment :one
UPDATE installments
SET paid_at = $2,
    updated_at = NOW()
WHERE id = $1 AND deleted_at IS NULL
RETURNING *;

-- name: DeleteInstallment :exec
UPDATE installments
SET deleted_at = NOW()
WHERE id = $1;