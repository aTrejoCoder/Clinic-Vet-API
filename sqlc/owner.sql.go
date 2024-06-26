// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0
// source: owner.sql

package sqlc

import (
	"context"

	"github.com/jackc/pgx/v5/pgtype"
)

const createOwner = `-- name: CreateOwner :one
INSERT INTO owners (photo, name, phone, user_id, created_at, updated_at)
VALUES ($1, $2, $3, $4, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP)
RETURNING id, photo, name, phone, user_id, created_at, updated_at
`

type CreateOwnerParams struct {
	Photo  pgtype.Text
	Name   string
	Phone  pgtype.Text
	UserID pgtype.Int4
}

func (q *Queries) CreateOwner(ctx context.Context, arg CreateOwnerParams) (Owner, error) {
	row := q.db.QueryRow(ctx, createOwner,
		arg.Photo,
		arg.Name,
		arg.Phone,
		arg.UserID,
	)
	var i Owner
	err := row.Scan(
		&i.ID,
		&i.Photo,
		&i.Name,
		&i.Phone,
		&i.UserID,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const deleteOwner = `-- name: DeleteOwner :exec
DELETE FROM owners
WHERE id = $1
`

func (q *Queries) DeleteOwner(ctx context.Context, id int32) error {
	_, err := q.db.Exec(ctx, deleteOwner, id)
	return err
}

const getOwnerByID = `-- name: GetOwnerByID :one
SELECT id, photo, name, phone, user_id, created_at, updated_at
FROM owners
WHERE id = $1
`

func (q *Queries) GetOwnerByID(ctx context.Context, id int32) (Owner, error) {
	row := q.db.QueryRow(ctx, getOwnerByID, id)
	var i Owner
	err := row.Scan(
		&i.ID,
		&i.Photo,
		&i.Name,
		&i.Phone,
		&i.UserID,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const listOwners = `-- name: ListOwners :many
SELECT id, photo, name, phone, user_id, created_at, updated_at
FROM owners
ORDER BY id
`

func (q *Queries) ListOwners(ctx context.Context) ([]Owner, error) {
	rows, err := q.db.Query(ctx, listOwners)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Owner
	for rows.Next() {
		var i Owner
		if err := rows.Scan(
			&i.ID,
			&i.Photo,
			&i.Name,
			&i.Phone,
			&i.UserID,
			&i.CreatedAt,
			&i.UpdatedAt,
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

const updateOwner = `-- name: UpdateOwner :exec
UPDATE owners
SET photo = $2, name = $3, phone = $4, user_id = $5, updated_at = CURRENT_TIMESTAMP
WHERE id = $1
`

type UpdateOwnerParams struct {
	ID     int32
	Photo  pgtype.Text
	Name   string
	Phone  pgtype.Text
	UserID pgtype.Int4
}

func (q *Queries) UpdateOwner(ctx context.Context, arg UpdateOwnerParams) error {
	_, err := q.db.Exec(ctx, updateOwner,
		arg.ID,
		arg.Photo,
		arg.Name,
		arg.Phone,
		arg.UserID,
	)
	return err
}