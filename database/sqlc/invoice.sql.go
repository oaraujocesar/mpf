// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.24.0
// source: invoice.sql

package database

import (
	"context"
	"database/sql"
	"time"
)

const createInvoice = `-- name: CreateInvoice :one
INSERT INTO invoices (amount, account_id, close_at, card_id, due_at)
VALUES ($1, $2, $3, $4, $5)
RETURNING id, amount, account_id, close_at, card_id, due_at, paid_at, created_at, updated_at, deleted_at
`

type CreateInvoiceParams struct {
	Amount    float64   `json:"amount"`
	AccountID int64     `json:"account_id"`
	CloseAt   time.Time `json:"close_at"`
	CardID    int64     `json:"card_id"`
	DueAt     time.Time `json:"due_at"`
}

func (q *Queries) CreateInvoice(ctx context.Context, arg CreateInvoiceParams) (Invoice, error) {
	row := q.db.QueryRowContext(ctx, createInvoice,
		arg.Amount,
		arg.AccountID,
		arg.CloseAt,
		arg.CardID,
		arg.DueAt,
	)
	var i Invoice
	err := row.Scan(
		&i.ID,
		&i.Amount,
		&i.AccountID,
		&i.CloseAt,
		&i.CardID,
		&i.DueAt,
		&i.PaidAt,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.DeletedAt,
	)
	return i, err
}

const deleteInvoice = `-- name: DeleteInvoice :exec
UPDATE invoices
SET deleted_at = NOW()
WHERE id = $1
`

func (q *Queries) DeleteInvoice(ctx context.Context, id int64) error {
	_, err := q.db.ExecContext(ctx, deleteInvoice, id)
	return err
}

const getInvoiceById = `-- name: GetInvoiceById :one
SELECT id, amount, account_id, close_at, card_id, due_at, paid_at, created_at, updated_at, deleted_at
FROM invoices
WHERE id = $1 AND deleted_at IS NULL
LIMIT 1
`

func (q *Queries) GetInvoiceById(ctx context.Context, id int64) (Invoice, error) {
	row := q.db.QueryRowContext(ctx, getInvoiceById, id)
	var i Invoice
	err := row.Scan(
		&i.ID,
		&i.Amount,
		&i.AccountID,
		&i.CloseAt,
		&i.CardID,
		&i.DueAt,
		&i.PaidAt,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.DeletedAt,
	)
	return i, err
}

const listInvoices = `-- name: ListInvoices :many
SELECT id, amount, account_id, close_at, card_id, due_at, paid_at, created_at, updated_at, deleted_at
FROM invoices
ORDER BY $1
LIMIT $2 OFFSET $3
`

type ListInvoicesParams struct {
	Column1 interface{} `json:"column_1"`
	Limit   int32       `json:"limit"`
	Offset  int32       `json:"offset"`
}

func (q *Queries) ListInvoices(ctx context.Context, arg ListInvoicesParams) ([]Invoice, error) {
	rows, err := q.db.QueryContext(ctx, listInvoices, arg.Column1, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []Invoice{}
	for rows.Next() {
		var i Invoice
		if err := rows.Scan(
			&i.ID,
			&i.Amount,
			&i.AccountID,
			&i.CloseAt,
			&i.CardID,
			&i.DueAt,
			&i.PaidAt,
			&i.CreatedAt,
			&i.UpdatedAt,
			&i.DeletedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const updateInvoice = `-- name: UpdateInvoice :one
UPDATE invoices
SET due_at = $2,
    close_at = $3,
    account_id = $4,
    paid_at = $5,
    updated_at = NOW()
WHERE id = $1 AND deleted_at IS NULL
RETURNING id, amount, account_id, close_at, card_id, due_at, paid_at, created_at, updated_at, deleted_at
`

type UpdateInvoiceParams struct {
	ID        int64        `json:"id"`
	DueAt     time.Time    `json:"due_at"`
	CloseAt   time.Time    `json:"close_at"`
	AccountID int64        `json:"account_id"`
	PaidAt    sql.NullTime `json:"paid_at"`
}

func (q *Queries) UpdateInvoice(ctx context.Context, arg UpdateInvoiceParams) (Invoice, error) {
	row := q.db.QueryRowContext(ctx, updateInvoice,
		arg.ID,
		arg.DueAt,
		arg.CloseAt,
		arg.AccountID,
		arg.PaidAt,
	)
	var i Invoice
	err := row.Scan(
		&i.ID,
		&i.Amount,
		&i.AccountID,
		&i.CloseAt,
		&i.CardID,
		&i.DueAt,
		&i.PaidAt,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.DeletedAt,
	)
	return i, err
}

const updateInvoiceAmount = `-- name: UpdateInvoiceAmount :one
UPDATE invoices
SET amount = $2,
    updated_at = NOW()
WHERE id = $1 AND deleted_at IS NULL
RETURNING id, amount, account_id, close_at, card_id, due_at, paid_at, created_at, updated_at, deleted_at
`

type UpdateInvoiceAmountParams struct {
	ID     int64   `json:"id"`
	Amount float64 `json:"amount"`
}

func (q *Queries) UpdateInvoiceAmount(ctx context.Context, arg UpdateInvoiceAmountParams) (Invoice, error) {
	row := q.db.QueryRowContext(ctx, updateInvoiceAmount, arg.ID, arg.Amount)
	var i Invoice
	err := row.Scan(
		&i.ID,
		&i.Amount,
		&i.AccountID,
		&i.CloseAt,
		&i.CardID,
		&i.DueAt,
		&i.PaidAt,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.DeletedAt,
	)
	return i, err
}
