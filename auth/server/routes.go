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
	router.POST("/register", handlers.RegisterUserHandler)
	// Авторизация пользователя
	router.POST("SmakTown/API/signIn", handlers.SignInHandler)
	// Обновление токена
	router.PUT("SmakTown/API/refresh", handlers.RefreshTokenHandler)
	// Получение данных пользователя
	router.GET("SmakTown/API/user", handlers.GetUserHandler)

	auth := router.Group("/auth")
	auth.Use(handlers.AuthMiddleware())
	{
		// Получение данных от пользователя, если пропустит перехватчик
		auth.GET("SmakTown/API/user", handlers.GetUserHandler)
	}
	router.Run(":" + envs.ServerEnvs.AUTH_PORT)
}
