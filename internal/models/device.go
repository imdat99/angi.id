package models

import (
	"time"

	"github.com/google/uuid"
)

type Device struct {
	Id              uuid.UUID `json:"id"`
	UserID          uuid.UUID `json:"user_id"`
	DeviceName      string    `json:"device_name"`
	DeviceType      string    `json:"device_type"`
	LastActive      time.Time `json:"last_active"`
	Location        *string   `json:"location"`
	IsCurrentDevice bool      `json:"is_current_device"`

	// Relation
	ActivityLogs []ActivityLog `json:"activity_logs,omitempty"`
}
