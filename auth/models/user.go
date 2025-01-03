package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	UserName string `gorm:"username" json:"username"`
	Email    string `gorm:"not nul;unique" json:"email"`
	Phone    string `gorm:"not nul;unique" json:"phone"`
	Hash     string `gorm:"hash" json:"-"`
}
