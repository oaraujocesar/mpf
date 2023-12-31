-- name: CreateAccount :one
INSERT INTO accounts (name, balance, user_id)
VALUES ($1, $2, $3)
RETURNING *;

-- name: ListAccounts :many
SELECT *
FROM accounts
LIMIT $1 OFFSET $2;

-- name: GetAccountById :one
SELECT *
FROM accounts
WHERE id = $1 AND deleted_at IS NULL
LIMIT 1;

-- name: UpdateBalance :one
UPDATE accounts
SET balance = $2,
    updated_at = NOW()
WHERE id = $1
RETURNING *;

-- name: UpdateAccount :one
UPDATE accounts
SET name = $2,
    balance = $3,
    updated_at = NOW()
WHERE id = $1 AND deleted_at IS NULL
RETURNING *;

-- name: DeleteAccount :exec
UPDATE accounts
SET deleted_at = NOW()
WHERE id = $1;
