package repository

import (
	"github.com/jmoiron/sqlx"
)

type Notify interface {
	GetEmail(UserID string) (string, error)
	GetNumber(UserID string) (string, error)
	GetPushToken(UserID string) ([]string, error)
}
type Repository struct {
	Notify
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{}
}
