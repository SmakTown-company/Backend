package services

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/SmakTown-company/Backend/notify/internal/models"
	"github.com/SmakTown-company/Backend/notify/internal/repository"
	"github.com/SmakTown-company/Backend/notify/pkg/logging"
	"github.com/k3a/html2text"
	"gopkg.in/gomail.v2"
	"html/template"
	"os"
	"strconv"

	"strings"
)

type EmailService struct {
	repo     repository.Notify
	ctx      context.Context
	from     string
	smtpPass string
	smtpUser string
	smtpHost string
	smtpPort string
}

func (e *EmailService) Get(UserID string) (string, error) {
	email, err := e.repo.GetEmail(UserID)
	if err != nil {
		return "", err
	}
	return email, nil
}

func NewEmailService(repo repository.Notify, ctx context.Context) *EmailService {
	return &EmailService{
		repo:     repo,
		from:     os.Getenv("EMAIL_FROM"),
		smtpPass: os.Getenv("SMTP_PASS"),
		smtpUser: os.Getenv("SMTP_USER"),
		smtpHost: os.Getenv("SMTP_HOST"),
		smtpPort: os.Getenv("SMTP_PORT"),
		ctx:      ctx,
	}
}
func (e *EmailService) Send(Data interface{}, To string) error {
	newData, ok := Data.(models.NotificationRequest)
	if !ok {
		return fmt.Errorf("Ошибка преобразования в EmailData")
	}
	jsonData, err := json.Marshal(newData.Data)
	if err != nil {
		return err
	}
	var data models.EmailData
	err = json.Unmarshal(jsonData, &data)
	if err != nil {
		return fmt.Errorf("Ошибка преобразования в EmailData")
	}
	FilledHTML := e.ParseHTML(data)
	m := e.MakeMessage(To, data.Subject, FilledHTML)
	port, err := strconv.Atoi(e.smtpPort)
	if err != nil {
		return fmt.Errorf("invalid SMTP port: %v", err.Error())
	}
	dialer := gomail.NewDialer(e.smtpHost, port, e.smtpUser, e.smtpPass)
	if err := dialer.DialAndSend(m); err != nil {
		return err
	}
	return nil
}
func (e *EmailService) MakeMessage(To string, subject string, message string) *gomail.Message {
	m := gomail.NewMessage()
	m.SetHeader("From", e.from)
	m.SetHeader("To", To)
	m.SetHeader("Subject", subject)
	m.SetBody("text/html", message)
	m.AddAlternative("text/plain", html2text.HTML2Text(message))
	return m
}
func (e *EmailService) ParseHTML(Data models.EmailData) string {
	tmpl, err := template.ParseFiles("assets/email/emailTemplate.html")
	if err != nil {
		logging.Logger.Warn("Ошибка при чтение шаблона")
	}
	var emailBody strings.Builder
	err = tmpl.Execute(&emailBody, Data)
	if err != nil {
		panic(err)
	}
	return emailBody.String()
}
