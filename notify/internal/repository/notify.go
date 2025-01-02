package repository

import (
	"github.com/jmoiron/sqlx"
)

type NotifyMongo struct {
	db *sqlx.DB
}

func NewNotifyPostgres(mongoDB *sqlx.DB) *NotifyMongo {
	return &NotifyMongo{db: mongoDB}
}
func (n *NotifyMongo) GetEmail(UserID string) (string, error) {
	query := `SELECT email FROM users WHERE id=$1`
	var email string
	err := n.db.QueryRow(query, UserID).Scan(&email)
	if err != nil {
		return "", err
	}
	return email, nil
}
func (n *NotifyMongo) GetNumber(UserID string) (string, error) {
	return "", nil
}
func (n *NotifyMongo) GetPushToken(UserID string) ([]string, error) {
	return nil, nil

}
