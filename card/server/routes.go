package server

import (
	"card/handlers"

	"github.com/gin-gonic/gin"
)

func InitRoutes() *gin.Engine {
	router := gin.Default()
	router.GET("/card/:id", handlers.GetCardHandler)

	router.Run(":8080")

	return router
}
