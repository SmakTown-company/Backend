package server

import (
	"basket/handlers"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func InitRoutes() *gin.Engine {
	router := gin.Default()

	// Настройка CORS (должна быть до определения маршрутов)
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:5173"},                   // Разрешаем доступ только с этого домена
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"}, // Разрешаем методы
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"}, // Разрешаем заголовки
		AllowCredentials: true,                                                // Разрешаем использование cookies и других учётных данных
		MaxAge:           12 * 3600,                                           // Максимальное время кеширования CORS в секундах (12 часов)
	}))

	router.POST("SmakTown/API/addInBasket", handlers.AddInBasketHandler)

	router.Run(":8081")

	return router
}
