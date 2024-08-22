// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.25.0

package sqlc

import (
	"github.com/jackc/pgx/v5/pgtype"
)

type Appointment struct {
	ID        int32
	PetID     int32
	VetID     int32
	Service   string
	Date      pgtype.Timestamp
	CreatedAt pgtype.Timestamp
	UpdatedAt pgtype.Timestamp
}

type MedicalHistory struct {
	ID          int32
	PetID       int32
	Date        pgtype.Timestamp
	Description pgtype.Text
	VetID       int32
	CreatedAt   pgtype.Timestamp
	UpdatedAt   pgtype.Timestamp
}

type Owner struct {
	ID        int32
	Photo     pgtype.Text
	Name      string
	LastName  string
	UserID    pgtype.Int4
	Birthday  pgtype.Date
	Genre     pgtype.Text
	CreatedAt pgtype.Timestamp
	UpdatedAt pgtype.Timestamp
}

type Payment struct {
	ID            int32
	AppointmentID int32
	Amount        pgtype.Numeric
	PaymentMethod string
	CreatedAt     pgtype.Timestamp
	UpdatedAt     pgtype.Timestamp
}

type Pet struct {
	ID        int32
	Name      string
	Photo     pgtype.Text
	Species   string
	Breed     pgtype.Text
	Age       pgtype.Int4
	OwnerID   int32
	CreatedAt pgtype.Timestamp
	UpdatedAt pgtype.Timestamp
}

type Reminder struct {
	ID            int32
	AppointmentID int32
	Method        string
	TimeBefore    string
	CreatedAt     pgtype.Timestamp
	UpdatedAt     pgtype.Timestamp
}

type User struct {
	ID          int32
	Name        string
	Email       string
	PhoneNumber string
	Password    string
	Role        string
	CreatedAt   pgtype.Timestamp
	UpdatedAt   pgtype.Timestamp
}

type Veterinarian struct {
	ID        int32
	Name      string
	Photo     pgtype.Text
	Specialty pgtype.Text
	UserID    pgtype.Int4
	CreatedAt pgtype.Timestamp
	UpdatedAt pgtype.Timestamp
}
