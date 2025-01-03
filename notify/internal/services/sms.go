package services

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/SmakTown-company/Backend/notify/internal/models"
	"github.com/SmakTown-company/Backend/notify/internal/repository"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
)

type SmsService struct {
	//Login    string
	//Password string
	ctx    context.Context
	APIKEY string
	repo   repository.Notify
}

func NewSmsService(repo repository.Notify, ctx context.Context) *SmsService {
	S := SmsService{
		//Login:    os.Getenv("SMS_LOGIN"),
		//Password: os.Getenv("SMS_PASSWORD"),
		ctx:    ctx,
		APIKEY: os.Getenv("SMS_APIKEY"),
		repo:   repo,
	}
	return &S
}
func (s SmsService) Get(UserID string) (string, error) {
	//TODO implement me
	panic("implement me")
}

func (s *SmsService) Send(Data interface{}, To string) error {
	newData, ok := Data.(models.NotificationRequest)
	if !ok {
		return fmt.Errorf("Ошибка преобразования в SmsData")
	}
	jsonData, err := json.Marshal(newData.Data)
	if err != nil {
		return err
	}
	var data models.SmsData
	err = json.Unmarshal(jsonData, &data)
	if err != nil {
		return fmt.Errorf("Ошибка преобразования в SmsData")
	}
	req, err := s.MakeURL(data)
	if err != nil {
		return fmt.Errorf("ошибка при создании запроса: %v", err)
	}

	// Отправка запроса
	client := http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("ошибка при отправке запроса: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, _ := ioutil.ReadAll(resp.Body)
		return fmt.Errorf("неожиданный статус ответа: %d, тело: %s", resp.StatusCode, string(body))
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("ошибка при чтении тела ответа: %v", err)
	}

	fmt.Printf("Ответ сервера: %s\n", string(body))
	return nil
	//
	//client := twilio.NewRestClient()
	//params := &api.CreateMessageParams{}
	//params.SetBody(data.BodyText)
	//params.SetFrom(s.From)
	//params.SetTo(data.To)
	//
	//resp, err := client.Api.CreateMessage(params)
	//if err != nil {
	//	return err
	//}
	//if resp.Body != nil {
	//	fmt.Println(*resp.Body)
	//} else {
	//	fmt.Println(resp.Body)
	//}
}

func (s *SmsService) MakeURL(data models.SmsData) (*http.Request, error) {
	params := url.Values{}
	params.Add("apikey", s.APIKEY)
	params.Add("phones", data.To)
	params.Add("mes", data.BodyText)

	fullURL := "https://smsc.ru/sys/send.php" + "?" + params.Encode()

	req, err := http.NewRequest("GET", fullURL, nil)
	if err != nil {
		return nil, fmt.Errorf("ошибка при создании запроса: %v", err)
	}
	return req, nil
}
