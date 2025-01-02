package services

import "github.com/SmakTown-company/Backend/notify/internal/repository"

type SmsService struct {
	repo repository.Notify
}

func (s SmsService) Get(UserID string) (string, error) {
	//TODO implement me
	panic("implement me")
}

func (s SmsService) Send(Template string, Data interface{}, To string) error {
	//TODO implement me
	panic("implement me")
}

func NewSmsService(repo repository.Notify) *SmsService {
	return &SmsService{
		repo: repo,
	}
}
