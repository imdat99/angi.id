package models

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	Id             uuid.UUID  `json:"id"`
	Email          string     `json:"email"`
	Name           string     `json:"name"`
	ProfilePicture *string    `json:"profile_picture_url"`
	CreationDate   time.Time  `json:"creation_date"`
	LastLogin      *time.Time `json:"last_login"`
	IsActive       bool       `json:"is_active"`
	PasswordHash   string     `json:"password_hash"`

	// Relations
	SecuritySettings *SecuritySettings `json:"security_settings,omitempty"`
	PersonalInfo     *PersonalInfo     `json:"personal_info,omitempty"`
	PrivacySettings  *PrivacySettings  `json:"privacy_settings,omitempty"`
	Devices          []Device          `json:"devices,omitempty"`
	ActivityLogs     []ActivityLog     `json:"activity_logs,omitempty"`
}

type UserFilter struct {
	Email           *string    `json:"email,omitempty"`
	IsActive        *bool      `json:"is_active,omitempty"`
	CreatedAfter    *time.Time `json:"created_after,omitempty"`
	CreatedBefore   *time.Time `json:"created_before,omitempty"`
	LastLoginAfter  *time.Time `json:"last_login_after,omitempty"`
	LastLoginBefore *time.Time `json:"last_login_before,omitempty"`
	Limit           int        `json:"limit,omitempty"`
	Offset          int        `json:"offset,omitempty"`
	Name            *string    `json:"name,omitempty"`
}
