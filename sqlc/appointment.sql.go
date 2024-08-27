// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0
// source: appointment.sql

package sqlc

import (
	"context"

	"github.com/jackc/pgx/v5/pgtype"
)

const createAppointment = `-- name: CreateAppointment :one
INSERT INTO appointments (pet_id, owner_id, service, date, status, created_at, updated_at)
VALUES ($1, $2, $3, $4, $5, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP)
RETURNING id, pet_id, vet_id, owner_id, service, date, status, created_at, updated_at
`

type CreateAppointmentParams struct {
	PetID   int32
	OwnerID int32
	Service string
	Date    pgtype.Timestamp
	Status  string
}

func (q *Queries) CreateAppointment(ctx context.Context, arg CreateAppointmentParams) (Appointment, error) {
	row := q.db.QueryRow(ctx, createAppointment,
		arg.PetID,
		arg.OwnerID,
		arg.Service,
		arg.Date,
		arg.Status,
	)
	var i Appointment
	err := row.Scan(
		&i.ID,
		&i.PetID,
		&i.VetID,
		&i.OwnerID,
		&i.Service,
		&i.Date,
		&i.Status,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const deleteAppointment = `-- name: DeleteAppointment :exec
DELETE FROM appointments
WHERE id = $1
`

func (q *Queries) DeleteAppointment(ctx context.Context, id int32) error {
	_, err := q.db.Exec(ctx, deleteAppointment, id)
	return err
}

const getAppointmentByID = `-- name: GetAppointmentByID :one
SELECT id, pet_id, vet_id, owner_id, service, date, status, created_at, updated_at
FROM appointments
WHERE id = $1
`

func (q *Queries) GetAppointmentByID(ctx context.Context, id int32) (Appointment, error) {
	row := q.db.QueryRow(ctx, getAppointmentByID, id)
	var i Appointment
	err := row.Scan(
		&i.ID,
		&i.PetID,
		&i.VetID,
		&i.OwnerID,
		&i.Service,
		&i.Date,
		&i.Status,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const listAppointmentsByOwnerID = `-- name: ListAppointmentsByOwnerID :many
SELECT id, pet_id, vet_id, owner_id, service, date, status, created_at, updated_at
FROM appointments
WHERE owner_id = $1
`

func (q *Queries) ListAppointmentsByOwnerID(ctx context.Context, ownerID int32) ([]Appointment, error) {
	rows, err := q.db.Query(ctx, listAppointmentsByOwnerID, ownerID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Appointment
	for rows.Next() {
		var i Appointment
		if err := rows.Scan(
			&i.ID,
			&i.PetID,
			&i.VetID,
			&i.OwnerID,
			&i.Service,
			&i.Date,
			&i.Status,
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

const updateAppointment = `-- name: UpdateAppointment :exec
UPDATE appointments
SET pet_id = $2, vet_id = $3, owner_id = $4, vet_id = $5, service = $6, date = $7, status = $8,  updated_at = CURRENT_TIMESTAMP
WHERE id = $1
`

type UpdateAppointmentParams struct {
	ID      int32
	PetID   int32
	VetID   pgtype.Int4
	OwnerID int32
	VetID_2 pgtype.Int4
	Service string
	Date    pgtype.Timestamp
	Status  string
}

func (q *Queries) UpdateAppointment(ctx context.Context, arg UpdateAppointmentParams) error {
	_, err := q.db.Exec(ctx, updateAppointment,
		arg.ID,
		arg.PetID,
		arg.VetID,
		arg.OwnerID,
		arg.VetID_2,
		arg.Service,
		arg.Date,
		arg.Status,
	)
	return err
}

const updateAppointmentStatus = `-- name: UpdateAppointmentStatus :exec
UPDATE appointments
SET status = $2, updated_at = CURRENT_TIMESTAMP
WHERE id = $1
`

type UpdateAppointmentStatusParams struct {
	ID     int32
	Status string
}

func (q *Queries) UpdateAppointmentStatus(ctx context.Context, arg UpdateAppointmentStatusParams) error {
	_, err := q.db.Exec(ctx, updateAppointmentStatus, arg.ID, arg.Status)
	return err
}
