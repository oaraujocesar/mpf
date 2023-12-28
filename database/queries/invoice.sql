-- name: CreateInvoice :one
INSERT INTO invoices (amount, account_id, close_at, card_id, due_at)
VALUES ($1, $2, $3, $4, $5)
RETURNING *;

-- name: ListInvoices :many
SELECT *
FROM invoices
ORDER BY $1
LIMIT $2 OFFSET $3;

-- name: GetInvoiceById :one
SELECT *
FROM invoices
WHERE id = $1 AND deleted_at IS NULL
LIMIT 1;

-- name: UpdateInvoiceAmount :one
UPDATE invoices
SET amount = $2,
    updated_at = NOW()
WHERE id = $1 AND deleted_at IS NULL
RETURNING *;

-- name: UpdateInvoice :one
UPDATE invoices
SET due_at = $2,
    close_at = $3,
    account_id = $4,
    paid_at = $5,
    updated_at = NOW()
WHERE id = $1 AND deleted_at IS NULL
RETURNING *;

-- name: DeleteInvoice :exec
UPDATE invoices
SET deleted_at = NOW()
WHERE id = $1;