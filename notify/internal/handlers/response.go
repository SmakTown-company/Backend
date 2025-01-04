package handlers

import (
	"fmt"
	"github.com/SmakTown-company/Backend/notify/pkg/logging"
	"github.com/gin-gonic/gin"
)

type Response struct {
	Message string `json:"message,omitempty"`
}

func newResponse(c *gin.Context, statusCode int, message string, err error) {
	if err != nil {
		logging.Logger.Warn(logging.MakeLog(fmt.Sprintf("Код ошибки: %d\t", statusCode), err))
	} else {
		logging.Logger.Info(message)
	}
	c.AbortWithStatusJSON(statusCode, Response{message})
}
