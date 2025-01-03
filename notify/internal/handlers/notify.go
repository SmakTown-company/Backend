package handlers

import (
	"github.com/SmakTown-company/Backend/notify/internal/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

const (
	SMS   = "sms"
	EMAIL = "email"
	PUSH  = "push"
)

func (h *Handler) Notify(c *gin.Context) {
	var input models.NotificationRequest

	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, "Неправильный формат данных")
		return
	}
	switch input.Channel {
	case EMAIL:
		var email string
		switch input.Template {
		case "email_confirmation":
			email = input.Data.(map[string]interface{})["email"].(string)
		default:
			var err error
			email, err = h.service.Email.Get(input.UserID)
			if err != nil {
				newErrorResponse(c, http.StatusBadRequest, "Ошибка получения email адреса")
				return
			}
		}

		err := h.service.Email.Send(input, email)
		if err != nil {
			newErrorResponse(c, http.StatusBadRequest, "Ошибка отправки сообщения")
			return
		}
	case SMS:
		var phone_number string
		switch input.Template {
		case "sms_confirmation":
			phone_number = input.Data.(map[string]interface{})["to"].(string)
		default:
			var err error
			phone_number, err = h.service.SMS.Get(input.UserID)
			if err != nil {
				newErrorResponse(c, http.StatusBadRequest, "Ошибка получения email адреса")
				return
			}
		}
		err := h.service.SMS.Send(input, phone_number)
		if err != nil {
			newErrorResponse(c, http.StatusBadRequest, "Ошибка отправки сообщения")
			return
		}
	case PUSH:
		err := h.service.Push.Send(input, input.UserID)
		if err != nil {
			newErrorResponse(c, http.StatusBadRequest, "Ошибка отправки сообщения")
			return
		}
	}
}
func (h *Handler) GetNotification(c *gin.Context) {
	userID := c.Query("id")
	pushTokens, err := h.service.Push.Get(userID)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	c.AbortWithStatusJSON(http.StatusOK, pushTokens)

}
