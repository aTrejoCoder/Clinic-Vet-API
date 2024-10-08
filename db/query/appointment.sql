-- name: CreateAppointment :one
INSERT INTO appointments (pet_id, vet_id, owner_id, service, date, status, created_at, updated_at)
VALUES ($1, $2, $3, $4, $5, $6, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP)
RETURNING id, pet_id, vet_id, owner_id, service, date, status, created_at, updated_at;

-- name: RequestAppointment :one
INSERT INTO appointments (pet_id, owner_id, service, date, status, created_at, updated_at)
VALUES ($1, $2, $3, $4, $5, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP)
RETURNING id, pet_id, vet_id, owner_id, service, date, status, created_at, updated_at;

-- name: GetAppointmentByID :one
SELECT id, pet_id, vet_id, owner_id, service, date, status, created_at, updated_at
FROM appointments
WHERE id = $1;

-- name: ListAppointmentsByOwnerID :many
SELECT id, pet_id, vet_id, owner_id, service, date, status, created_at, updated_at
FROM appointments
WHERE owner_id = $1;

-- name: UpdateAppointment :exec
UPDATE appointments
SET pet_id = $2, vet_id = $3, owner_id = $4 ,service = $5, date = $6, status = $7,  updated_at = CURRENT_TIMESTAMP
WHERE id = $1;

-- name: UpdateOwnerAppointment :exec
UPDATE appointments
SET pet_id = $2, service = $3, date = $4, updated_at = CURRENT_TIMESTAMP
WHERE id = $1;

-- name: UpdateAppointmentStatus :exec
UPDATE appointments
SET status = $2, updated_at = CURRENT_TIMESTAMP
WHERE id = $1;

-- name: DeleteAppointment :exec
DELETE FROM appointments
WHERE id = $1;