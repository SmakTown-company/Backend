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

		err := h.service.Email.Send(input.Template, input, email)
		if err != nil {
			newErrorResponse(c, http.StatusBadRequest, "Ошибка отправки сообщения")
			return
		}
	}
}
