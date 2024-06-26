-- name: CreateUser :one
INSERT INTO users (name, email, password, user_id, role, created_at, updated_at)
VALUES ($1, $2, $3, $4, $5, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP)
RETURNING id, name, email, password, user_id, role, created_at, updated_at;

-- name: GetUserByID :one
SELECT id, name, email, password, user_id, role, created_at, updated_at
FROM users
WHERE id = $1;

-- name: ListUsers :many
SELECT id, name, email, password, user_id, role, created_at, updated_at
FROM users
ORDER BY id;

-- name: UpdateUser :exec
UPDATE users
SET name = $2, email = $3, password = $4, user_id = $5, role = $6, updated_at = CURRENT_TIMESTAMP
WHERE id = $1;

-- name: DeleteUser :exec
DELETE FROM users
WHERE id = $1;
