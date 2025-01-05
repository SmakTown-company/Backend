package model

type User struct {
	UserID     string `json:"user_id" db:"user_id"`
	Email      string `json:"email" db:"email"`
	Phone      string `json:"phone" db:"phone"`
	OrderCount int    `json:"order_count" db:"order_count"`
}
