// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.25.0
// source: veterinarian.sql

package sqlc

import (
	"context"

	"github.com/jackc/pgx/v5/pgtype"
)

const createVeterinarian = `-- name: CreateVeterinarian :one
INSERT INTO veterinarians (name, photo, specialty, user_id, created_at, updated_at)
VALUES ($1, $2, $3, $4, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP)
RETURNING id, name, photo, specialty, user_id, created_at, updated_at
`

type CreateVeterinarianParams struct {
	Name      string
	Photo     pgtype.Text
	Specialty pgtype.Text
	UserID    pgtype.Int4
}

func (q *Queries) CreateVeterinarian(ctx context.Context, arg CreateVeterinarianParams) (Veterinarian, error) {
	row := q.db.QueryRow(ctx, createVeterinarian,
		arg.Name,
		arg.Photo,
		arg.Specialty,
		arg.UserID,
	)
	var i Veterinarian
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Photo,
		&i.Specialty,
		&i.UserID,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const deleteVeterinarian = `-- name: DeleteVeterinarian :exec
DELETE FROM veterinarians
WHERE id = $1
`

func (q *Queries) DeleteVeterinarian(ctx context.Context, id int32) error {
	_, err := q.db.Exec(ctx, deleteVeterinarian, id)
	return err
}

const getVeterinarianByID = `-- name: GetVeterinarianByID :one
SELECT id, name, photo, specialty, user_id, created_at, updated_at
FROM veterinarians
WHERE id = $1
`

func (q *Queries) GetVeterinarianByID(ctx context.Context, id int32) (Veterinarian, error) {
	row := q.db.QueryRow(ctx, getVeterinarianByID, id)
	var i Veterinarian
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Photo,
		&i.Specialty,
		&i.UserID,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const getVeterinarianByUserID = `-- name: GetVeterinarianByUserID :one
SELECT id, name, photo, specialty, user_id, created_at, updated_at
FROM veterinarians
WHERE user_id = $1
`

func (q *Queries) GetVeterinarianByUserID(ctx context.Context, userID pgtype.Int4) (Veterinarian, error) {
	row := q.db.QueryRow(ctx, getVeterinarianByUserID, userID)
	var i Veterinarian
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Photo,
		&i.Specialty,
		&i.UserID,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const listVeterinarians = `-- name: ListVeterinarians :many
SELECT id, name, photo, specialty, user_id, created_at, updated_at
FROM veterinarians
ORDER BY id
`

func (q *Queries) ListVeterinarians(ctx context.Context) ([]Veterinarian, error) {
	rows, err := q.db.Query(ctx, listVeterinarians)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Veterinarian
	for rows.Next() {
		var i Veterinarian
		if err := rows.Scan(
			&i.ID,
			&i.Name,
			&i.Photo,
			&i.Specialty,
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

const updateVeterinarian = `-- name: UpdateVeterinarian :exec
UPDATE veterinarians
SET name = $2, photo = $3, specialty = $4, user_id = $5, updated_at = CURRENT_TIMESTAMP
WHERE id = $1
`

type UpdateVeterinarianParams struct {
	ID        int32
	Name      string
	Photo     pgtype.Text
	Specialty pgtype.Text
	UserID    pgtype.Int4
}

func (q *Queries) UpdateVeterinarian(ctx context.Context, arg UpdateVeterinarianParams) error {
	_, err := q.db.Exec(ctx, updateVeterinarian,
		arg.ID,
		arg.Name,
		arg.Photo,
		arg.Specialty,
		arg.UserID,
	)
	return err
}

const updateVeterinarianUserID = `-- name: UpdateVeterinarianUserID :exec
UPDATE veterinarians
SET user_id = $2, updated_at = CURRENT_TIMESTAMP
WHERE id = $1
`

type UpdateVeterinarianUserIDParams struct {
	ID     int32
	UserID pgtype.Int4
}

func (q *Queries) UpdateVeterinarianUserID(ctx context.Context, arg UpdateVeterinarianUserIDParams) error {
	_, err := q.db.Exec(ctx, updateVeterinarianUserID, arg.ID, arg.UserID)
	return err
}
