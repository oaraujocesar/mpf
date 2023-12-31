// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.24.0
// source: categories.sql

package database

import (
	"context"
)

const createCategory = `-- name: CreateCategory :one
INSERT INTO categories (name)
VALUES ($1)
RETURNING id, name, created_at, updated_at, deleted_at
`

func (q *Queries) CreateCategory(ctx context.Context, name string) (Category, error) {
	row := q.db.QueryRowContext(ctx, createCategory, name)
	var i Category
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.DeletedAt,
	)
	return i, err
}

const deleteCategory = `-- name: DeleteCategory :exec
UPDATE categories
SET deleted_at = NOW()
WHERE id = $1
`

func (q *Queries) DeleteCategory(ctx context.Context, id int64) error {
	_, err := q.db.ExecContext(ctx, deleteCategory, id)
	return err
}

const getCategoryById = `-- name: GetCategoryById :one
SELECT id, name, created_at, updated_at, deleted_at
FROM categories
WHERE id = $1 AND deleted_at IS NULL
LIMIT 1
`

func (q *Queries) GetCategoryById(ctx context.Context, id int64) (Category, error) {
	row := q.db.QueryRowContext(ctx, getCategoryById, id)
	var i Category
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.DeletedAt,
	)
	return i, err
}

const hardDeleteCategory = `-- name: HardDeleteCategory :exec
DELETE FROM categories
WHERE id = $1
`

func (q *Queries) HardDeleteCategory(ctx context.Context, id int64) error {
	_, err := q.db.ExecContext(ctx, hardDeleteCategory, id)
	return err
}

const listCategories = `-- name: ListCategories :many
SELECT id, name, created_at, updated_at, deleted_at
FROM categories
LIMIT $1 OFFSET $2
`

type ListCategoriesParams struct {
	Limit  int32 `json:"limit"`
	Offset int32 `json:"offset"`
}

func (q *Queries) ListCategories(ctx context.Context, arg ListCategoriesParams) ([]Category, error) {
	rows, err := q.db.QueryContext(ctx, listCategories, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []Category{}
	for rows.Next() {
		var i Category
		if err := rows.Scan(
			&i.ID,
			&i.Name,
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

const updateCategory = `-- name: UpdateCategory :one
UPDATE categories
SET name = $2,
    updated_at = NOW()
WHERE id = $1 AND deleted_at IS NULL
RETURNING id, name, created_at, updated_at, deleted_at
`

type UpdateCategoryParams struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
}

func (q *Queries) UpdateCategory(ctx context.Context, arg UpdateCategoryParams) (Category, error) {
	row := q.db.QueryRowContext(ctx, updateCategory, arg.ID, arg.Name)
	var i Category
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.DeletedAt,
	)
	return i, err
}
