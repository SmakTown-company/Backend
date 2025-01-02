package handlers

import (
	"github.com/SmakTown-company/Backend/notify/pkg/logging"
	"github.com/gin-gonic/gin"
)

type Response struct {
	Message string `json:"text,omitempty" example:"Аккаунт подтвержден"`
}

func newErrorResponse(c *gin.Context, statusCode int, message string) {
	logging.Logger.Warn(message)
	c.AbortWithStatusJSON(statusCode, Response{message})
}
func newResultResponse(c *gin.Context, text string) {
	c.AbortWithStatusJSON(200, Response{Message: text})
}
