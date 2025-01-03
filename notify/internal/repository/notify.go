package repository

import (
	"context"
	"fmt"
	"github.com/SmakTown-company/Backend/notify/internal/models"
	"github.com/jmoiron/sqlx"
)

type NotifyMongo struct {
	ctx context.Context
	db  *sqlx.DB
}

func NewNotifyPostgres(db *sqlx.DB, ctx context.Context) *NotifyMongo {
	return &NotifyMongo{db: db, ctx: ctx}
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
	query := `SELECT phone FROM users WHERE id=$1`
	var number string
	err := n.db.QueryRow(query, UserID).Scan(&number)
	if err != nil {
		return "", err
	}
	return number, nil
}

func (n *NotifyMongo) GetPushToken(UserID string) ([]models.PushToken, error) {
	tx, err := n.db.BeginTx(n.ctx, nil)
	defer tx.Rollback()

	if err != nil {
		return nil, fmt.Errorf("Ошибка запуска транзакции: %v", err)
	}
	query := `SELECT user_id, subject,body_text, created_at  FROM push_tokens WHERE user_id=$1 AND is_looked=false`
	rows, err := tx.QueryContext(n.ctx, query, UserID)
	if err != nil {
		return nil, fmt.Errorf("Ошибка выполнения запроса: %w", err)

	}
	defer rows.Close()
	var pushTokens []models.PushToken
	for rows.Next() {
		var pt models.PushToken
		if err := rows.Scan(&pt.UserID, &pt.Subject, &pt.BodyText, &pt.CreatedAt); err != nil {
			return nil, fmt.Errorf("Ошибка сканирования данных: %w", err)
		}
		pushTokens = append(pushTokens, pt)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("Ошибка в ходе итерации: %w", err)
	}
	update_query := `UPDATE push_tokens SET is_looked = true WHERE user_id = $1 AND is_looked = false`
	_, err = tx.ExecContext(n.ctx, update_query, UserID)
	if err != nil {
		return nil, fmt.Errorf("Ошибка обновления push-уведомлений: %w", err)
	}
	if err := tx.Commit(); err != nil {
		return nil, fmt.Errorf("Ошибка фиксации транзакции: %w", err)
	}
	return pushTokens, nil
}
func (n *NotifyMongo) SendPushToken(UserID string, data models.PushData) error {
	tx, err := n.db.BeginTx(n.ctx, nil)
	defer tx.Rollback() // Откат транзакции в случае ошибки

	if err != nil {
		return fmt.Errorf("Ошибка запуска транзакции: %v", err)

	}
	query := `INSERT INTO push_tokens(user_id, subject, body_text) VALUES ($1, $2, $3)`
	_, err = tx.ExecContext(n.ctx, query, UserID, data.Subject, data.BodyText)
	if err != nil {
		return fmt.Errorf("Ошибка вставки push-уведомления: %w", err)
	}
	if err := tx.Commit(); err != nil {
		return fmt.Errorf("Ошибка фиксации транзакции: %w", err)
	}
	return nil
}
