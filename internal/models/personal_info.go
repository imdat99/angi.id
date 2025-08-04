package models

import (
	"time"

	"github.com/google/uuid"
)

type PersonalInfo struct {
	Id              uuid.UUID  `json:"id"`
	UserID          uuid.UUID  `json:"user_id"`
	FirstName       *string    `json:"first_name"`
	LastName        *string    `json:"last_name"`
	DateOfBirth     *time.Time `json:"date_of_birth"`
	Gender          *string    `json:"gender"`
	PhoneNumber     *string    `json:"phone_number"`
	AlternateEmails []string   `json:"alternate_emails"`
}
