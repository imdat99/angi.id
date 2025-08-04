package models

import (
	"time"

	"github.com/google/uuid"
)

type SecurityQuestion struct {
	Id           uuid.UUID `json:"id"`
	SettingID    uuid.UUID `json:"setting_id"`
	QuestionText string    `json:"question_text"`
	AnswerHash   string    `json:"answer_hash"`
	CreatedAt    time.Time `json:"created_at"`
}
