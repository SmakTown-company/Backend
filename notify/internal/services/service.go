package services

import (
	"context"
	"github.com/SmakTown-company/Backend/notify/internal/repository"
)

type Service struct {
	Email *EmailService
	SMS   *SmsService
	Push  *PushService
}

type NotifyService interface {
	Send(Template string, Data interface{}, To string) error
	Get(UserID string) ([]string, error)
}

func NewService(repo *repository.Repository, ctx context.Context) *Service {
	return &Service{
		Email: NewEmailService(repo.Notify, ctx),
		SMS:   NewSmsService(repo.Notify, ctx),
		Push:  NewPushService(repo.Notify, ctx),
	}
}
