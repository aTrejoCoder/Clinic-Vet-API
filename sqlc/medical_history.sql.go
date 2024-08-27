// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0
// source: medical_history.sql

package sqlc

import (
	"context"

	"github.com/jackc/pgx/v5/pgtype"
)

const createMedicalHistory = `-- name: CreateMedicalHistory :one
INSERT INTO medical_histories (pet_id, date, description, vet_id, created_at, updated_at)
VALUES ($1, $2, $3, $4, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP)
RETURNING id, pet_id, date, description, vet_id, created_at, updated_at
`

type CreateMedicalHistoryParams struct {
	PetID       int32
	Date        pgtype.Timestamp
	Description pgtype.Text
	VetID       int32
}

func (q *Queries) CreateMedicalHistory(ctx context.Context, arg CreateMedicalHistoryParams) (MedicalHistory, error) {
	row := q.db.QueryRow(ctx, createMedicalHistory,
		arg.PetID,
		arg.Date,
		arg.Description,
		arg.VetID,
	)
	var i MedicalHistory
	err := row.Scan(
		&i.ID,
		&i.PetID,
		&i.Date,
		&i.Description,
		&i.VetID,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const deleteMedicalHistory = `-- name: DeleteMedicalHistory :exec
DELETE FROM medical_histories
WHERE id = $1
`

func (q *Queries) DeleteMedicalHistory(ctx context.Context, id int32) error {
	_, err := q.db.Exec(ctx, deleteMedicalHistory, id)
	return err
}

const getMedicalHistoryByID = `-- name: GetMedicalHistoryByID :one
SELECT id, pet_id, date, description, vet_id, created_at, updated_at
FROM medical_histories
WHERE id = $1
`

func (q *Queries) GetMedicalHistoryByID(ctx context.Context, id int32) (MedicalHistory, error) {
	row := q.db.QueryRow(ctx, getMedicalHistoryByID, id)
	var i MedicalHistory
	err := row.Scan(
		&i.ID,
		&i.PetID,
		&i.Date,
		&i.Description,
		&i.VetID,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const listMedicalHistoriesByPetID = `-- name: ListMedicalHistoriesByPetID :many
SELECT id, pet_id, date, description, vet_id, created_at, updated_at
FROM medical_histories
WHERE pet_id = $1
ORDER BY date
`

func (q *Queries) ListMedicalHistoriesByPetID(ctx context.Context, petID int32) ([]MedicalHistory, error) {
	rows, err := q.db.Query(ctx, listMedicalHistoriesByPetID, petID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []MedicalHistory
	for rows.Next() {
		var i MedicalHistory
		if err := rows.Scan(
			&i.ID,
			&i.PetID,
			&i.Date,
			&i.Description,
			&i.VetID,
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

const listMedicalHistoriesByVetID = `-- name: ListMedicalHistoriesByVetID :many
SELECT id, pet_id, date, description, vet_id, created_at, updated_at
FROM medical_histories
WHERE vet_id = $1
ORDER BY date
`

func (q *Queries) ListMedicalHistoriesByVetID(ctx context.Context, vetID int32) ([]MedicalHistory, error) {
	rows, err := q.db.Query(ctx, listMedicalHistoriesByVetID, vetID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []MedicalHistory
	for rows.Next() {
		var i MedicalHistory
		if err := rows.Scan(
			&i.ID,
			&i.PetID,
			&i.Date,
			&i.Description,
			&i.VetID,
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

const updateMedicalHistory = `-- name: UpdateMedicalHistory :exec
UPDATE medical_histories
SET pet_id = $2, date = $3, description = $4, vet_id = $5, updated_at = CURRENT_TIMESTAMP
WHERE id = $1
`

type UpdateMedicalHistoryParams struct {
	ID          int32
	PetID       int32
	Date        pgtype.Timestamp
	Description pgtype.Text
	VetID       int32
}

func (q *Queries) UpdateMedicalHistory(ctx context.Context, arg UpdateMedicalHistoryParams) error {
	_, err := q.db.Exec(ctx, updateMedicalHistory,
		arg.ID,
		arg.PetID,
		arg.Date,
		arg.Description,
		arg.VetID,
	)
	return err
}
