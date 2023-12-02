// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.24.0
// source: query.sql

package database

import (
	"context"

	"github.com/jackc/pgx/v5/pgtype"
)

const createAuthor = `-- name: CreateAuthor :one
INSERT INTO users (name, avatar, email, password)
VALUES ($1, $2, $3, $4)
RETURNING id, name, avatar, email, password, created_at, updated_at, deleted_at
`

type CreateAuthorParams struct {
	Name     string      `json:"name"`
	Avatar   pgtype.Text `json:"avatar"`
	Email    string      `json:"email"`
	Password string      `json:"password"`
}

func (q *Queries) CreateAuthor(ctx context.Context, arg CreateAuthorParams) (User, error) {
	row := q.db.QueryRow(ctx, createAuthor,
		arg.Name,
		arg.Avatar,
		arg.Email,
		arg.Password,
	)
	var i User
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Avatar,
		&i.Email,
		&i.Password,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.DeletedAt,
	)
	return i, err
}

const createCategory = `-- name: CreateCategory :one
INSERT INTO categories (name)
VALUES ($1) 
RETURNING id, name, created_at, updated_at, deleted_at
`

func (q *Queries) CreateCategory(ctx context.Context, name string) (Category, error) {
	row := q.db.QueryRow(ctx, createCategory, name)
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

const createEntry = `-- name: CreateEntry :one
INSERT INTO entries (title, amount, type, user_id, category_id)
VALUES ($1, $2, $3, $4, $5)
RETURNING id, title, amount, type, user_id, category_id, created_at, updated_at, deleted_at
`

type CreateEntryParams struct {
	Title      string         `json:"title"`
	Amount     pgtype.Numeric `json:"amount"`
	Type       EntryType      `json:"type"`
	UserID     pgtype.UUID    `json:"user_id"`
	CategoryID pgtype.UUID    `json:"category_id"`
}

func (q *Queries) CreateEntry(ctx context.Context, arg CreateEntryParams) (Entry, error) {
	row := q.db.QueryRow(ctx, createEntry,
		arg.Title,
		arg.Amount,
		arg.Type,
		arg.UserID,
		arg.CategoryID,
	)
	var i Entry
	err := row.Scan(
		&i.ID,
		&i.Title,
		&i.Amount,
		&i.Type,
		&i.UserID,
		&i.CategoryID,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.DeletedAt,
	)
	return i, err
}

const deleteAuthor = `-- name: DeleteAuthor :exec
DELETE FROM users
WHERE id = $1
`

func (q *Queries) DeleteAuthor(ctx context.Context, id pgtype.UUID) error {
	_, err := q.db.Exec(ctx, deleteAuthor, id)
	return err
}

const deleteCategory = `-- name: DeleteCategory :exec
DELETE FROM categories
WHERE id = $1
`

func (q *Queries) DeleteCategory(ctx context.Context, id pgtype.UUID) error {
	_, err := q.db.Exec(ctx, deleteCategory, id)
	return err
}

const deleteEntry = `-- name: DeleteEntry :exec
DELETE FROM entries
WHERE id = $1
`

func (q *Queries) DeleteEntry(ctx context.Context, id pgtype.UUID) error {
	_, err := q.db.Exec(ctx, deleteEntry, id)
	return err
}

const getAuthor = `-- name: GetAuthor :one
SELECT id, name, avatar, email, password, created_at, updated_at, deleted_at
FROM users
WHERE email = $1
LIMIT 1
`

func (q *Queries) GetAuthor(ctx context.Context, email string) (User, error) {
	row := q.db.QueryRow(ctx, getAuthor, email)
	var i User
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Avatar,
		&i.Email,
		&i.Password,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.DeletedAt,
	)
	return i, err
}

const getCategory = `-- name: GetCategory :one
SELECT id, name, created_at, updated_at, deleted_at
FROM categories
WHERE name = $1
LIMIT 1
`

func (q *Queries) GetCategory(ctx context.Context, name string) (Category, error) {
	row := q.db.QueryRow(ctx, getCategory, name)
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

const getEntry = `-- name: GetEntry :one
SELECT id, title, amount, type, user_id, category_id, created_at, updated_at, deleted_at
FROM entries
WHERE id = $1
LIMIT 1
`

func (q *Queries) GetEntry(ctx context.Context, id pgtype.UUID) (Entry, error) {
	row := q.db.QueryRow(ctx, getEntry, id)
	var i Entry
	err := row.Scan(
		&i.ID,
		&i.Title,
		&i.Amount,
		&i.Type,
		&i.UserID,
		&i.CategoryID,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.DeletedAt,
	)
	return i, err
}

const listAuthors = `-- name: ListAuthors :many
SELECT id, name, avatar, email, password, created_at, updated_at, deleted_at
FROM users
ORDER BY $1
LIMIT $2 OFFSET $3
`

type ListAuthorsParams struct {
	Column1 interface{} `json:"column_1"`
	Limit   int32       `json:"limit"`
	Offset  int32       `json:"offset"`
}

func (q *Queries) ListAuthors(ctx context.Context, arg ListAuthorsParams) ([]User, error) {
	rows, err := q.db.Query(ctx, listAuthors, arg.Column1, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []User{}
	for rows.Next() {
		var i User
		if err := rows.Scan(
			&i.ID,
			&i.Name,
			&i.Avatar,
			&i.Email,
			&i.Password,
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

const listCategories = `-- name: ListCategories :many
SELECT id, name, created_at, updated_at, deleted_at
FROM categories
ORDER BY $1
LIMIT $2 OFFSET $3
`

type ListCategoriesParams struct {
	Column1 interface{} `json:"column_1"`
	Limit   int32       `json:"limit"`
	Offset  int32       `json:"offset"`
}

func (q *Queries) ListCategories(ctx context.Context, arg ListCategoriesParams) ([]Category, error) {
	rows, err := q.db.Query(ctx, listCategories, arg.Column1, arg.Limit, arg.Offset)
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
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const listEntries = `-- name: ListEntries :many
SELECT id, title, amount, type, user_id, category_id, created_at, updated_at, deleted_at
FROM entries
ORDER BY $1
LIMIT $2 OFFSET $3
`

type ListEntriesParams struct {
	Column1 interface{} `json:"column_1"`
	Limit   int32       `json:"limit"`
	Offset  int32       `json:"offset"`
}

func (q *Queries) ListEntries(ctx context.Context, arg ListEntriesParams) ([]Entry, error) {
	rows, err := q.db.Query(ctx, listEntries, arg.Column1, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []Entry{}
	for rows.Next() {
		var i Entry
		if err := rows.Scan(
			&i.ID,
			&i.Title,
			&i.Amount,
			&i.Type,
			&i.UserID,
			&i.CategoryID,
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

const updateAuthor = `-- name: UpdateAuthor :one
UPDATE users
SET name = $2,
    avatar = $3,
    email = $4,
    password = $5,
    updated_at = NOW()
WHERE id = $1
RETURNING id, name, avatar, email, password, created_at, updated_at, deleted_at
`

type UpdateAuthorParams struct {
	ID       pgtype.UUID `json:"id"`
	Name     string      `json:"name"`
	Avatar   pgtype.Text `json:"avatar"`
	Email    string      `json:"email"`
	Password string      `json:"password"`
}

func (q *Queries) UpdateAuthor(ctx context.Context, arg UpdateAuthorParams) (User, error) {
	row := q.db.QueryRow(ctx, updateAuthor,
		arg.ID,
		arg.Name,
		arg.Avatar,
		arg.Email,
		arg.Password,
	)
	var i User
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Avatar,
		&i.Email,
		&i.Password,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.DeletedAt,
	)
	return i, err
}

const updateCategory = `-- name: UpdateCategory :one
UPDATE categories
SET name = $2,
    updated_at = NOW()
WHERE id = $1
RETURNING id, name, created_at, updated_at, deleted_at
`

type UpdateCategoryParams struct {
	ID   pgtype.UUID `json:"id"`
	Name string      `json:"name"`
}

func (q *Queries) UpdateCategory(ctx context.Context, arg UpdateCategoryParams) (Category, error) {
	row := q.db.QueryRow(ctx, updateCategory, arg.ID, arg.Name)
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

const updateEntry = `-- name: UpdateEntry :one
UPDATE entries
SET title = $2,
    amount = $3,
    type = $4,
    user_id = $5,
    category_id = $6,
    updated_at = NOW()
WHERE id = $1
RETURNING id, title, amount, type, user_id, category_id, created_at, updated_at, deleted_at
`

type UpdateEntryParams struct {
	ID         pgtype.UUID    `json:"id"`
	Title      string         `json:"title"`
	Amount     pgtype.Numeric `json:"amount"`
	Type       EntryType      `json:"type"`
	UserID     pgtype.UUID    `json:"user_id"`
	CategoryID pgtype.UUID    `json:"category_id"`
}

func (q *Queries) UpdateEntry(ctx context.Context, arg UpdateEntryParams) (Entry, error) {
	row := q.db.QueryRow(ctx, updateEntry,
		arg.ID,
		arg.Title,
		arg.Amount,
		arg.Type,
		arg.UserID,
		arg.CategoryID,
	)
	var i Entry
	err := row.Scan(
		&i.ID,
		&i.Title,
		&i.Amount,
		&i.Type,
		&i.UserID,
		&i.CategoryID,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.DeletedAt,
	)
	return i, err
}
