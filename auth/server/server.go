package server

import (
	"auth/database"
	"auth/envs"
	"auth/models"
	"log"
)

func InitServer() {
	// Инициализация внешних значений ENV
	errEnvs := envs.LoadEnvs()
	if errEnvs != nil {
		log.Fatal("Ошибка загрузки ENV: ", errEnvs)
	} else {
		log.Println("Успешное получение ENV")
	}
	// Инициализация базы данных
	errDatabase := database.InitDatabase()
	if errDatabase != nil {
		log.Fatal("Ошибка подключения к базе данных: ", errDatabase)
	} else {
		// Автоматическое создание таблицы на основе модели Note, если она не существует
		database.DB.AutoMigrate(&models.User{})

		log.Println("Успешное подключение к базе данных")
	}
}

func StartServer() {
	InitRoutes()
	// Запуск сервера
}
