package repository

import (
	"context"
	"github.com/SmakTown-company/Backend/payment/internal/model"
	"github.com/redis/go-redis/v9"
)

type Payment interface {
	WritePayment(payment model.Payment) error
	WriteUser(user model.User) (model.User, error)
}

func NewRepository(db *redis.Client, ctx context.Context) *Repository {
	return &Repository{Payment: NewPaymentRedis(db, ctx)}
}

type Repository struct {
	Payment
}
