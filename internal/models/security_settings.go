package models

import "github.com/google/uuid"

type SecuritySettings struct {
	Id               uuid.UUID `json:"id"`
	UserID           uuid.UUID `json:"user_id"`
	RecoveryEmail    *string   `json:"recovery_email"`
	RecoveryPhone    *string   `json:"recovery_phone"`
	TwoFactorEnabled bool      `json:"two_factor_enabled"`
	BackupCodes      []string  `json:"backup_codes"`

	// Relation
	SecurityQuestions []SecurityQuestion `json:"security_questions,omitempty"`
}
