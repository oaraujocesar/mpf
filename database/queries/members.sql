-- name: CreateMember :one
INSERT INTO members (family_id, user_id)
VALUES ($1, $2)
RETURNING *;

-- name: ListMembers :many
SELECT *
FROM members
ORDER BY $1
LIMIT $2 OFFSET $3;

-- name: GetMemberById :one
SELECT *
FROM members
WHERE id = $1 AND deleted_at IS NULL
LIMIT 1;

-- name: DeleteMember :exec
UPDATE members
SET deleted_at = NOW()
WHERE id = $1;