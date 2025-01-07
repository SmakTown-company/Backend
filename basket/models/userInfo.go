package models

type UserInfo struct {
	ID       uint `json:"user_id"`
	Verified bool `json:"verified"`
}
