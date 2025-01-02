package services

import "github.com/SmakTown-company/Backend/notify/internal/repository"

type PushService struct {
	repo repository.Notify
}

func (p PushService) Get(UserID string) (string, error) {
	//TODO implement me
	panic("implement me")
}

func (p PushService) Send(Template string, Data interface{}, To string) error {
	//TODO implement me
	panic("implement me")
}

func NewPushService(repo repository.Notify) *PushService {
	return &PushService{repo: repo}
}
