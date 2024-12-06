package server

import (
	"auth/envs"
	"auth/handlers"

	"github.com/gin-gonic/gin"
)

func InitRoutes() {
	// Инициализация роута (по умолчанию)
	router := gin.Default()
	// Создание пользователя
	router.PUT("/user", handlers.RegisterUserHandler)
	// Авторизация пользователя
	router.POST("/user", handlers.SignInHandler)
	// Обновление токена
	router.PUT("/refresh", handlers.RefreshTokenHandler)
	// Получение данных пользователя
	router.GET("/user", handlers.GetUserHandler)

	auth := router.Group("/auth")
	auth.Use(handlers.AuthMiddleware())
	{
		// Получение данных от пользователя, если пропустит перехватчик
		auth.GET("/user", handlers.GetUserHandler)
	}
	router.Run(":" + envs.ServerEnvs.AUTH_PORT)
}
