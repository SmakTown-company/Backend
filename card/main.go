package main

import (
	"card/server"
	"log"
)

func main() {

	// Инициализация и запуск маршрутов
	router := server.InitRoutes()
	if err := router.Run(":8080"); err != nil {
		log.Fatal("Failed to start server:", err)
	}
}
