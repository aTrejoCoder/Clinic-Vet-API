// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.25.0
// source: pets.sql

package sqlc

import (
	"context"

	"github.com/jackc/pgx/v5/pgtype"
)

const createPet = `-- name: CreatePet :one
INSERT INTO pets (name, photo, species, breed, age, owner_id, created_at, updated_at)
VALUES ($1, $2, $3, $4, $5, $6, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP)
RETURNING id, name, photo, species, breed, age, owner_id, created_at, updated_at
`

type CreatePetParams struct {
	Name    string
	Photo   pgtype.Text
	Species string
	Breed   pgtype.Text
	Age     pgtype.Int4
	OwnerID int32
}

func (q *Queries) CreatePet(ctx context.Context, arg CreatePetParams) (Pet, error) {
	row := q.db.QueryRow(ctx, createPet,
		arg.Name,
		arg.Photo,
		arg.Species,
		arg.Breed,
		arg.Age,
		arg.OwnerID,
	)
	var i Pet
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Photo,
		&i.Species,
		&i.Breed,
		&i.Age,
		&i.OwnerID,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const deletePet = `-- name: DeletePet :exec
DELETE FROM pets
WHERE id = $1
`

func (q *Queries) DeletePet(ctx context.Context, id int32) error {
	_, err := q.db.Exec(ctx, deletePet, id)
	return err
}

const getPetByID = `-- name: GetPetByID :one
SELECT id, name, photo, species, breed, age, owner_id, created_at, updated_at
FROM pets
WHERE id = $1
`

func (q *Queries) GetPetByID(ctx context.Context, id int32) (Pet, error) {
	row := q.db.QueryRow(ctx, getPetByID, id)
	var i Pet
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Photo,
		&i.Species,
		&i.Breed,
		&i.Age,
		&i.OwnerID,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const listPets = `-- name: ListPets :many
SELECT id, name, photo, species, breed, age, owner_id, created_at, updated_at
FROM pets
ORDER BY id
`

func (q *Queries) ListPets(ctx context.Context) ([]Pet, error) {
	rows, err := q.db.Query(ctx, listPets)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Pet
	for rows.Next() {
		var i Pet
		if err := rows.Scan(
			&i.ID,
			&i.Name,
			&i.Photo,
			&i.Species,
			&i.Breed,
			&i.Age,
			&i.OwnerID,
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

const listPetsByOwnerByID = `-- name: ListPetsByOwnerByID :many
SELECT id, name, photo, species, breed, age, owner_id, created_at, updated_at
FROM pets
WHERE owner_id = $1
`

func (q *Queries) ListPetsByOwnerByID(ctx context.Context, ownerID int32) ([]Pet, error) {
	rows, err := q.db.Query(ctx, listPetsByOwnerByID, ownerID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Pet
	for rows.Next() {
		var i Pet
		if err := rows.Scan(
			&i.ID,
			&i.Name,
			&i.Photo,
			&i.Species,
			&i.Breed,
			&i.Age,
			&i.OwnerID,
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

const updatePet = `-- name: UpdatePet :exec
UPDATE pets
SET name = $2, photo = $3, species = $4, breed = $5, age = $6, owner_id = $7, updated_at = CURRENT_TIMESTAMP
WHERE id = $1
`

type UpdatePetParams struct {
	ID      int32
	Name    string
	Photo   pgtype.Text
	Species string
	Breed   pgtype.Text
	Age     pgtype.Int4
	OwnerID int32
}

func (q *Queries) UpdatePet(ctx context.Context, arg UpdatePetParams) error {
	_, err := q.db.Exec(ctx, updatePet,
		arg.ID,
		arg.Name,
		arg.Photo,
		arg.Species,
		arg.Breed,
		arg.Age,
		arg.OwnerID,
	)
	return err
}
