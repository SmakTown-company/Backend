package models

type SmsData struct {
	Subject  string `json:"subject,omitempty"`
	BodyText string `json:"body_text,omitempty"`
	Code     string `json:"code,omitempty"`
	To       string `json:"to,omitempty"`
}
