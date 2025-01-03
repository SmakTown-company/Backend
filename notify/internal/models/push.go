package models

import "time"

type PushData struct {
	Subject  string `json:"subject,omitempty" db:"subject"`
	BodyText string `json:"body_text,omitempty" db:"body_text"`
	UserID   string `json:"userID,omitempty" db:"user_id"`
}
type PushToken struct {
	UserID    int       `json:"user_id"`
	Subject   string    `json:"subject"`
	BodyText  string    `json:"body_text"`
	CreatedAt time.Time `json:"created_at"`
}
