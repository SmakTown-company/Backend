package server

import (
	"auth/envs"
	"auth/handlers"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func InitRoutes() {
	// Инициализация роута (по умолчанию)
	router := gin.Default()

	// Настройка CORS (должна быть до определения маршрутов)
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:5173"},                   // Разрешаем доступ только с этого домена
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"}, // Разрешаем методы
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"}, // Разрешаем заголовки
		AllowCredentials: true,                                                // Разрешаем использование cookies и других учётных данных
		MaxAge:           12 * 3600,                                           // Максимальное время кеширования CORS в секундах (12 часов)
	}))

	// Роуты
	router.POST("SmakTown/API/register", handlers.RegisterUserHandler)
	router.POST("SmakTown/API/signIn", handlers.SignInHandler)
	router.PUT("SmakTown/API/refresh", handlers.RefreshTokenHandler)
	router.GET("SmakTown/API/user", handlers.GetUserHandler)
	router.GET("SmakTown/API/getUserBasket/:user_id", handlers.GetUserForBasket)
	router.PUT("SmakTown/API/renameUser/:email", handlers.RenameUserHandler)

	auth := router.Group("/auth")
	auth.Use(handlers.AuthMiddleware())
	{
		auth.GET("SmakTown/API/user", handlers.GetUserHandler)
	}

	// Запуск сервера с указанным портом из переменных окружения
	err := router.Run(":" + envs.ServerEnvs.AUTH_PORT)
	if err != nil {
		panic("Ошибка запуска сервера: " + err.Error())
	}
}
