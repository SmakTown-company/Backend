package repository

import (
	"context"
	"github.com/SmakTown-company/Backend/notify/internal/models"
	"github.com/jmoiron/sqlx"
)

type Notify interface {
	GetEmail(UserID string) (string, error)
	GetNumber(UserID string) (string, error)
	GetPushToken(UserID string) ([]models.PushToken, error)
	SendPushToken(UserID string, data models.PushData) error
}
type Repository struct {
	Notify
}

func NewRepository(db *sqlx.DB, ctx context.Context) *Repository {
	return &Repository{
		Notify: NewNotifyPostgres(db, ctx),
	}
}
