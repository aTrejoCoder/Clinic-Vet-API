// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0
// source: reminder.sql

package sqlc

import (
	"context"
)

const createReminder = `-- name: CreateReminder :one
INSERT INTO reminders (appointment_id, method, time_before, created_at, updated_at)
VALUES ($1, $2, $3, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP)
RETURNING id, appointment_id, method, time_before, created_at, updated_at
`

type CreateReminderParams struct {
	AppointmentID int32
	Method        string
	TimeBefore    string
}

func (q *Queries) CreateReminder(ctx context.Context, arg CreateReminderParams) (Reminder, error) {
	row := q.db.QueryRow(ctx, createReminder, arg.AppointmentID, arg.Method, arg.TimeBefore)
	var i Reminder
	err := row.Scan(
		&i.ID,
		&i.AppointmentID,
		&i.Method,
		&i.TimeBefore,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const deleteReminder = `-- name: DeleteReminder :exec
DELETE FROM reminders
WHERE id = $1
`

func (q *Queries) DeleteReminder(ctx context.Context, id int32) error {
	_, err := q.db.Exec(ctx, deleteReminder, id)
	return err
}

const getReminderByID = `-- name: GetReminderByID :one
SELECT id, appointment_id, method, time_before, created_at, updated_at
FROM reminders
WHERE id = $1
`

func (q *Queries) GetReminderByID(ctx context.Context, id int32) (Reminder, error) {
	row := q.db.QueryRow(ctx, getReminderByID, id)
	var i Reminder
	err := row.Scan(
		&i.ID,
		&i.AppointmentID,
		&i.Method,
		&i.TimeBefore,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const listReminders = `-- name: ListReminders :many
SELECT id, appointment_id, method, time_before, created_at, updated_at
FROM reminders
ORDER BY id
`

func (q *Queries) ListReminders(ctx context.Context) ([]Reminder, error) {
	rows, err := q.db.Query(ctx, listReminders)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Reminder
	for rows.Next() {
		var i Reminder
		if err := rows.Scan(
			&i.ID,
			&i.AppointmentID,
			&i.Method,
			&i.TimeBefore,
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

const updateReminder = `-- name: UpdateReminder :exec
UPDATE reminders
SET appointment_id = $2, method = $3, time_before = $4, updated_at = CURRENT_TIMESTAMP
WHERE id = $1
`

type UpdateReminderParams struct {
	ID            int32
	AppointmentID int32
	Method        string
	TimeBefore    string
}

func (q *Queries) UpdateReminder(ctx context.Context, arg UpdateReminderParams) error {
	_, err := q.db.Exec(ctx, updateReminder,
		arg.ID,
		arg.AppointmentID,
		arg.Method,
		arg.TimeBefore,
	)
	return err
}
