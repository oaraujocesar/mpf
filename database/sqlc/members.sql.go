// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.24.0
// source: members.sql

package database

import (
	"context"
)

const createMember = `-- name: CreateMember :one
INSERT INTO members (family_id, user_id)
VALUES ($1, $2)
RETURNING id, family_id, user_id, created_at, updated_at, deleted_at
`

type CreateMemberParams struct {
	FamilyID int64 `json:"family_id"`
	UserID   int64 `json:"user_id"`
}

func (q *Queries) CreateMember(ctx context.Context, arg CreateMemberParams) (Member, error) {
	row := q.db.QueryRowContext(ctx, createMember, arg.FamilyID, arg.UserID)
	var i Member
	err := row.Scan(
		&i.ID,
		&i.FamilyID,
		&i.UserID,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.DeletedAt,
	)
	return i, err
}

const deleteMember = `-- name: DeleteMember :exec
UPDATE members
SET deleted_at = NOW()
WHERE id = $1
`

func (q *Queries) DeleteMember(ctx context.Context, id int64) error {
	_, err := q.db.ExecContext(ctx, deleteMember, id)
	return err
}

const getMemberById = `-- name: GetMemberById :one
SELECT id, family_id, user_id, created_at, updated_at, deleted_at
FROM members
WHERE id = $1 AND deleted_at IS NULL
LIMIT 1
`

func (q *Queries) GetMemberById(ctx context.Context, id int64) (Member, error) {
	row := q.db.QueryRowContext(ctx, getMemberById, id)
	var i Member
	err := row.Scan(
		&i.ID,
		&i.FamilyID,
		&i.UserID,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.DeletedAt,
	)
	return i, err
}

const listMembers = `-- name: ListMembers :many
SELECT id, family_id, user_id, created_at, updated_at, deleted_at
FROM members
ORDER BY $1
LIMIT $2 OFFSET $3
`

type ListMembersParams struct {
	Column1 interface{} `json:"column_1"`
	Limit   int32       `json:"limit"`
	Offset  int32       `json:"offset"`
}

func (q *Queries) ListMembers(ctx context.Context, arg ListMembersParams) ([]Member, error) {
	rows, err := q.db.QueryContext(ctx, listMembers, arg.Column1, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []Member{}
	for rows.Next() {
		var i Member
		if err := rows.Scan(
			&i.ID,
			&i.FamilyID,
			&i.UserID,
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
