package main

import (
	"card/database"
	"card/server"
	"log"
)

func main() {
	// Инициализация подключения к базе данных
	if err := database.InitDatabase(); err != nil {
		log.Fatal("Failed to initialize database:", err)
	}
	defer database.CloseDatabase() // Закрываем соединение при завершении работы

	// Инициализация и запуск маршрутов
	router := server.InitRoutes()
	if err := router.Run(":8080"); err != nil {
		log.Fatal("Failed to start server:", err)
	}
}
