// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.24.0
// source: installments.sql

package database

import (
	"context"

	"github.com/jackc/pgx/v5/pgtype"
)

const createInstallment = `-- name: CreateInstallment :one
INSERT INTO installments (amount, entry_id, payday, paid_at)
VALUES ($1, $2, $3, $4)
RETURNING id, amount, entry_id, payday, paid_at, created_at, updated_at, deleted_at
`

type CreateInstallmentParams struct {
	Amount  float64            `json:"amount"`
	EntryID int64              `json:"entry_id"`
	Payday  pgtype.Timestamptz `json:"payday"`
	PaidAt  pgtype.Timestamptz `json:"paid_at"`
}

func (q *Queries) CreateInstallment(ctx context.Context, arg CreateInstallmentParams) (Installment, error) {
	row := q.db.QueryRow(ctx, createInstallment,
		arg.Amount,
		arg.EntryID,
		arg.Payday,
		arg.PaidAt,
	)
	var i Installment
	err := row.Scan(
		&i.ID,
		&i.Amount,
		&i.EntryID,
		&i.Payday,
		&i.PaidAt,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.DeletedAt,
	)
	return i, err
}

const deleteInstallment = `-- name: DeleteInstallment :exec
UPDATE installments
SET deleted_at = NOW()
WHERE id = $1
`

func (q *Queries) DeleteInstallment(ctx context.Context, id int64) error {
	_, err := q.db.Exec(ctx, deleteInstallment, id)
	return err
}

const getInstallmentById = `-- name: GetInstallmentById :one
SELECT id, amount, entry_id, payday, paid_at, created_at, updated_at, deleted_at
FROM installments
WHERE id = $1 AND deleted_at IS NULL
LIMIT 1
`

func (q *Queries) GetInstallmentById(ctx context.Context, id int64) (Installment, error) {
	row := q.db.QueryRow(ctx, getInstallmentById, id)
	var i Installment
	err := row.Scan(
		&i.ID,
		&i.Amount,
		&i.EntryID,
		&i.Payday,
		&i.PaidAt,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.DeletedAt,
	)
	return i, err
}

const listInstallments = `-- name: ListInstallments :many
SELECT id, amount, entry_id, payday, paid_at, created_at, updated_at, deleted_at
FROM installments
WHERE entry_id = $1
ORDER BY $2
LIMIT $3 OFFSET $4
`

type ListInstallmentsParams struct {
	EntryID int64       `json:"entry_id"`
	Column2 interface{} `json:"column_2"`
	Limit   int32       `json:"limit"`
	Offset  int32       `json:"offset"`
}

func (q *Queries) ListInstallments(ctx context.Context, arg ListInstallmentsParams) ([]Installment, error) {
	rows, err := q.db.Query(ctx, listInstallments,
		arg.EntryID,
		arg.Column2,
		arg.Limit,
		arg.Offset,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []Installment{}
	for rows.Next() {
		var i Installment
		if err := rows.Scan(
			&i.ID,
			&i.Amount,
			&i.EntryID,
			&i.Payday,
			&i.PaidAt,
			&i.CreatedAt,
			&i.UpdatedAt,
			&i.DeletedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const updateInstallment = `-- name: UpdateInstallment :one
UPDATE installments
SET paid_at = $2,
    updated_at = NOW()
WHERE id = $1 AND deleted_at IS NULL
RETURNING id, amount, entry_id, payday, paid_at, created_at, updated_at, deleted_at
`

type UpdateInstallmentParams struct {
	ID     int64              `json:"id"`
	PaidAt pgtype.Timestamptz `json:"paid_at"`
}

func (q *Queries) UpdateInstallment(ctx context.Context, arg UpdateInstallmentParams) (Installment, error) {
	row := q.db.QueryRow(ctx, updateInstallment, arg.ID, arg.PaidAt)
	var i Installment
	err := row.Scan(
		&i.ID,
		&i.Amount,
		&i.EntryID,
		&i.Payday,
		&i.PaidAt,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.DeletedAt,
	)
	return i, err
}
