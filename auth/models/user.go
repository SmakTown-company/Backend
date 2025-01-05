package models

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	UserName string `gorm:"username" json:"username"`
	Email    string `gorm:"not nul;unique" json:"email"`
	Phone    string `gorm:"not nul;unique" json:"phone"`
	Hash     string `gorm:"hash" json:"-"`
	Verified bool   `gorm:"verified" json:"verified"`
}

type VerificationToken struct {
	gorm.Model
	Email       string    `gorm:"not null"`
	Phone       string    `gorm:"not null"`
	Token       string    `gorm:"not null;unique"`
	ExpiresAt   time.Time `gorm:"not null"`
	ConfirmedAt time.Time
}
