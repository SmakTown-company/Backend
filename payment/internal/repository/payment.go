package repository

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/SmakTown-company/Backend/payment/internal/model"
	"github.com/redis/go-redis/v9"
	"strconv"
)

type PaymentRedis struct {
	ctx context.Context
	db  *redis.Client
}

func NewPaymentRedis(db *redis.Client, ctx context.Context) *PaymentRedis {
	return &PaymentRedis{db: db, ctx: ctx}
}
func (p *PaymentRedis) GetPayment(paymentID string) (map[string]string, error) {
	key := fmt.Sprintf("payment:%s", paymentID)
	val, err := p.db.HGetAll(p.ctx, key).Result()
	if err != nil {
		return nil, err
	}

	if len(val) == 0 {
		return nil, fmt.Errorf("payment with ID %s not found", paymentID)
	}
	return val, nil
}
func (p *PaymentRedis) WriteUser(user model.User) (model.User, error) {
	key := fmt.Sprintf("user:%s", user.UserID)

	val, err := p.db.HGetAll(p.ctx, key).Result()
	if err != nil {
		return model.User{}, err
	}

	if len(val) == 0 {
		user.OrderCount = 1
		_, err = p.db.HSet(p.ctx, key, "user_id", user.UserID, "email", user.Email, "phone", user.Phone, "order_count", user.OrderCount).Result()
		if err != nil {
			return model.User{}, err
		}
	}

	orderCount, _ := strconv.Atoi(val["order_count"])
	user.OrderCount = orderCount + 1

	_, err = p.db.HSet(p.ctx, key, "order_count", user.OrderCount).Result()
	if err != nil {
		return model.User{}, err
	}

	return user, nil
}

func (p *PaymentRedis) WritePayment(payment model.Payment) error {
	key := fmt.Sprintf("payment:%s", payment.ID)
	product, err := json.Marshal(payment.ProductList)
	if err != nil {
		return err
	}
	amount, err := json.Marshal(payment.Amount)
	if err != nil {
		return err
	}
	status, err := json.Marshal(payment.Status)
	if err != nil {
		return err
	}
	payment_method, err := json.Marshal(payment.PaymentMethod)
	if err != nil {
		return err
	}
	created_at, err := json.Marshal(payment.CreatedAt)
	if err != nil {
		return err
	}
	expires_at, err := json.Marshal(payment.ExpiresAt)
	if err != nil {
		return err
	}
	_, err = p.db.HSet(p.ctx, key, map[string]interface{}{
		"user_id":        payment.UserID,
		"amount":         amount,
		"status":         status,
		"payment_method": payment_method,
		"created_at":     created_at,
		"expires_at":     expires_at,
		"product_list":   product,
	}).Result()
	if err != nil {
		return fmt.Errorf("failed to save payment: %w", err)
	}
	fmt.Println(p.GetPayment(payment.ID))
	return nil

}
