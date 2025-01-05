package handler

import (
	"github.com/SmakTown-company/Backend/payment/internal/model"
	"github.com/gin-gonic/gin"
	yoowebhook "github.com/rvinnie/yookassa-sdk-go/yookassa/webhook"
	"log"
	"net"
	"net/http"
)

func (h *Handler) CreatePayment(c *gin.Context) {
	var input PaymentRequest
	if err := c.ShouldBindJSON(&input); err != nil {
		newResponse(c, http.StatusBadRequest, "Ошибка получения данных платежа", err)
		return
	}
	input.Payment.UserID = input.UserID
	payment, err := h.service.MakePayment(input.Payment)
	if err != nil {
		newResponse(c, http.StatusInternalServerError, "Ошибка создания платежа", err)
		return
	}
	url := payment.URL
	err = h.service.SavePayment(payment)
	if err != nil {
		newResponse(c, http.StatusInternalServerError, "Ошибка сохранения данных платежа", err)
		return
	}
	_, err = h.service.UpdateUser(model.User{
		UserID: input.UserID,
		Email:  input.Email,
		Phone:  input.Phone,
	})
	if err != nil {
		newResponse(c, http.StatusInternalServerError, "Ошибка сохранения данных пользователя", err)
		return
	}
	c.AbortWithStatusJSON(http.StatusOK, gin.H{"url": url})
}
func (h *Handler) IPFilterMiddleware(handler gin.HandlerFunc, allowedCIDRs []string) gin.HandlerFunc {
	return func(c *gin.Context) {
		remoteIP := c.ClientIP() // Gin автоматически выбирает правильный IP-адрес.
		log.Printf("Remote IP: %s", remoteIP)

		if !IsIPAllowed(remoteIP, allowedCIDRs) {
			c.JSON(http.StatusForbidden, gin.H{"error": "Forbidden"})
			c.Abort()
			return
		}

		handler(c)
	}
}
func IsIPAllowed(ip string, allowedCIDRs []string) bool {
	parsedIP := net.ParseIP(ip)
	if parsedIP == nil {
		return false
	}

	for _, cidr := range allowedCIDRs {
		_, allowedNet, err := net.ParseCIDR(cidr)
		if err != nil {
			continue
		}
		if allowedNet.Contains(parsedIP) {
			return true
		}
	}
	return false
}
func (h *Handler) PaymentStatus(c *gin.Context) {
	var webhookEvent yoowebhook.WebhookEvent
	if err := c.ShouldBindJSON(&webhookEvent); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid webhook data"})
		return
	}
	payment := webhookEvent.Object
	p := model.Payment{
		ID:            "",
		UserID:        "",
		Amount:        payment.Amount,
		Status:        payment.Status,
		PaymentMethod: payment.PaymentMethod,
		CreatedAt:     payment.CreatedAt,
		ExpiresAt:     payment.ExpiresAt,
		ProductList:   payment.P,
		URL:           "",
		Currency:      "",
	}
	c.Status(http.StatusOK)
}
