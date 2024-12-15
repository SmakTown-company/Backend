package server

import (
	"card/handlers"

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

	router.GET("SmakTown/API/getAllCards", handlers.GetCardHandler)
	router.GET("SmakTown/API/getCardDiscount", handlers.GetCardDiscountHandler)
	router.POST("SmakTown/API/addInBasket", handlers.AddInBasket)

	router.Run(":8080")

	return router
}
