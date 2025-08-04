package models

import (
	"net"
	"time"

	"github.com/google/uuid"
)

type ActivityType string

const (
	ActivityLogin            ActivityType = "LOGIN"
	ActivityPasswordChange   ActivityType = "PASSWORD_CHANGE"
	ActivitySecuritySetting  ActivityType = "SECURITY_SETTING_CHANGE"
	ActivityProfileUpdate    ActivityType = "PROFILE_UPDATE"
	ActivityDeviceConnection ActivityType = "DEVICE_CONNECTION"
)

type ActivityLog struct {
	Id           uuid.UUID    `json:"id"`
	UserID       uuid.UUID    `json:"user_id"`
	DeviceID     *uuid.UUID   `json:"device_id,omitempty"` // Nullable
	Timestamp    time.Time    `json:"timestamp"`
	ActivityType ActivityType `json:"activity_type"`
	IPAddress    net.IP       `json:"ip_address"`
	DeviceInfo   *string      `json:"device_info"`
	Location     *string      `json:"location"`
	Details      *string      `json:"details"`
}
