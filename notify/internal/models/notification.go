package models

const (
	SMS   = "sms"
	EMAIL = "email"
	PUSH  = "push"
)

type NotificationRequest struct {
	UserID   string      `json:"userID,omitempty"`
	Channel  string      `json:"channel,omitempty"`
	Template string      `json:"template,omitempty"`
	Data     interface{} `json:"data,omitempty"`
}
type NotificationResponse struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}
