package service

import (
	"context"
	"github.com/SmakTown-company/Backend/payment/internal/model"
	"github.com/SmakTown-company/Backend/payment/internal/repository"
)

type Payment interface {
	MakePayment(payment model.Payment) (model.Payment, error)
	SavePayment(payment model.Payment) error
	UpdateUser(user model.User) (model.User, error)
}

type Service struct {
	Payment
}

func NewService(paymentMethod, returnURL string, repo repository.Repository, ctx context.Context) *Service {
	return &Service{Payment: NewPaymentService(paymentMethod, returnURL, repo.Payment, ctx)}
}
