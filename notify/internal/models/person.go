package models

type Person struct {
	UserID    string   `json:"userID" db:"user_id"`
	Email     string   `json:"email" db:"email"`
	Number    string   `json:"number" db:"number"`
	PushToken []string `json:"pushToken" db:"push_token"`
}
