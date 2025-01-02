package services

import "github.com/SmakTown-company/Backend/notify/internal/repository"

type Service struct {
	Email NotifyService
	SMS   NotifyService
	Push  NotifyService
}

type NotifyService interface {
	Send(Template string, Data interface{}, To string) error
	Get(UserID string) (string, error)
}

func NewService(repo *repository.Repository) *Service {
	return &Service{
		Email: NewEmailService(repo.Notify),
		SMS:   NewSmsService(repo.Notify),
		Push:  NewPushService(repo.Notify),
	}
}
