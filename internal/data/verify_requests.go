package data

import (
	"time"

	"github.com/google/uuid"
)

type VerifyRequestsQ interface {
	New() VerifyRequestsQ

	Get() (*VerifyRequest, error)
	Select() ([]VerifyRequest, error)
	Update(*VerifyRequest) error
	Insert(*VerifyRequest) error
	Delete() error

	WhereID(id uuid.UUID) VerifyRequestsQ
	WhereCreatedAtLt(createdAt time.Time) VerifyRequestsQ
}

type VerifyRequest struct {
	ID           uuid.UUID          `db:"id"            structs:"id"`
	Status       VerificationStatus `db:"status"        structs:"status"`
	CallbackData []byte             `db:"callback_data" structs:"callback_data"`
	CreatedAt    time.Time          `db:"created_at"    structs:"created_at"`
	UpdatedAt    time.Time          `db:"updated_at"    structs:"updated_at"`
}

type VerificationStatus string

func (v VerificationStatus) String() string {
	return string(v)
}

const (
	VerificationStatusInitialized VerificationStatus = "initialized"
	VerificationStatusVerified    VerificationStatus = "verified"
)
