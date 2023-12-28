// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.24.0
// source: families.sql

package database

import (
	"context"
)

const createFamily = `-- name: CreateFamily :one
INSERT INTO families (name, user_id)
VALUES ($1, $2)
RETURNING id, name, user_id, created_at, updated_at, deleted_at
`

type CreateFamilyParams struct {
	Name   string `json:"name"`
	UserID int64  `json:"user_id"`
}

func (q *Queries) CreateFamily(ctx context.Context, arg CreateFamilyParams) (Family, error) {
	row := q.db.QueryRow(ctx, createFamily, arg.Name, arg.UserID)
	var i Family
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.UserID,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.DeletedAt,
	)
	return i, err
}

const deleteFamily = `-- name: DeleteFamily :exec
UPDATE families
SET deleted_at = NOW()
WHERE id = $1
`

func (q *Queries) DeleteFamily(ctx context.Context, id int64) error {
	_, err := q.db.Exec(ctx, deleteFamily, id)
	return err
}

const getFamilyById = `-- name: GetFamilyById :one
SELECT id, name, user_id, created_at, updated_at, deleted_at
FROM families
WHERE id = $1 AND deleted_at IS NULL
LIMIT 1
`

func (q *Queries) GetFamilyById(ctx context.Context, id int64) (Family, error) {
	row := q.db.QueryRow(ctx, getFamilyById, id)
	var i Family
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.UserID,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.DeletedAt,
	)
	return i, err
}

const listFamilies = `-- name: ListFamilies :many
SELECT id, name, user_id, created_at, updated_at, deleted_at
FROM families
ORDER BY $1
LIMIT $2 OFFSET $3
`

type ListFamiliesParams struct {
	Column1 interface{} `json:"column_1"`
	Limit   int32       `json:"limit"`
	Offset  int32       `json:"offset"`
}

func (q *Queries) ListFamilies(ctx context.Context, arg ListFamiliesParams) ([]Family, error) {
	rows, err := q.db.Query(ctx, listFamilies, arg.Column1, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []Family{}
	for rows.Next() {
		var i Family
		if err := rows.Scan(
			&i.ID,
			&i.Name,
			&i.UserID,
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

const updateFamily = `-- name: UpdateFamily :one
UPDATE families
SET name = $2,
    updated_at = NOW()
WHERE id = $1 AND deleted_at IS NULL
RETURNING id, name, user_id, created_at, updated_at, deleted_at
`

type UpdateFamilyParams struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
}

func (q *Queries) UpdateFamily(ctx context.Context, arg UpdateFamilyParams) (Family, error) {
	row := q.db.QueryRow(ctx, updateFamily, arg.ID, arg.Name)
	var i Family
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.UserID,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.DeletedAt,
	)
	return i, err
}
