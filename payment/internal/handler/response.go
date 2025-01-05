package handler

import (
	"fmt"
	"github.com/SmakTown-company/Backend/payment/internal/model"
	"github.com/SmakTown-company/Backend/payment/pkg/logging"
	"github.com/gin-gonic/gin"
)

type Response struct {
	Message string `json:"message,omitempty"`
}
type PaymentRequest struct {
	UserID  string        `json:"user_id"`
	Phone   string        `json:"phone"`
	Email   string        `json:"email"`
	Payment model.Payment `json:"payment"`
}

func newResponse(c *gin.Context, statusCode int, message string, err error) {
	if err != nil {
		logging.Logger.Warn(logging.MakeLog(fmt.Sprintf("Код ошибки: %d\t", statusCode), err))
	} else {
		logging.Logger.Info(message)
	}
	c.AbortWithStatusJSON(statusCode, Response{message})
}
