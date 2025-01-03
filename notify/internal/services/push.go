package services

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/SmakTown-company/Backend/notify/internal/models"
	"github.com/SmakTown-company/Backend/notify/internal/repository"
)

type PushService struct {
	ctx  context.Context
	repo repository.Notify
}

func (p PushService) Get(UserID string) ([]models.PushToken, error) {
	tokens, err := p.repo.GetPushToken(UserID)
	if err != nil {
		return nil, err
	}
	return tokens, nil
}

func (p PushService) Send(Data interface{}, To string) error {
	newData, ok := Data.(models.NotificationRequest)
	if !ok {
		return fmt.Errorf("Ошибка преобразования в PushData")
	}
	jsonData, err := json.Marshal(newData.Data)
	if err != nil {
		return err
	}
	var data models.PushData
	err = json.Unmarshal(jsonData, &data)
	if err != nil {
		return fmt.Errorf("Ошибка преобразования в PushData")
	}
	err = p.repo.SendPushToken(To, data)
	if err != nil {
		return err
	}
	return nil
}

func NewPushService(repo repository.Notify, ctx context.Context) *PushService {
	return &PushService{repo: repo}
}
