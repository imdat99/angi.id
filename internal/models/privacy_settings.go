package models

import "github.com/google/uuid"

type VisibilityType string

const (
	VisibilityPublic  VisibilityType = "PUBLIC"
	VisibilityFriends VisibilityType = "FRIENDS"
	VisibilityPrivate VisibilityType = "PRIVATE"
)

type PrivacySettings struct {
	Id                          uuid.UUID      `json:"id"`
	UserID                      uuid.UUID      `json:"user_id"`
	ProfileVisibility           VisibilityType `json:"profile_visibility"`
	ActivityVisibility          VisibilityType `json:"activity_visibility"`
	DataSharingWithThirdParties bool           `json:"data_sharing_with_third_parties"`
	PersonalizedAds             bool           `json:"personalized_ads"`
}
