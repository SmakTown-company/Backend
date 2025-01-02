package main

import (
	"fmt"
	"github.com/SmakTown-company/Backend/notify/internal/config"
	"github.com/SmakTown-company/Backend/notify/internal/handlers"
	"github.com/SmakTown-company/Backend/notify/internal/models"
	"github.com/SmakTown-company/Backend/notify/internal/repository"
	"github.com/SmakTown-company/Backend/notify/internal/services"
	"github.com/SmakTown-company/Backend/notify/pkg/logging"
	"github.com/joho/godotenv"
	"github.com/spf13/viper"
	"os"
)

func main() {
	if err := godotenv.Load(); err != nil {
		fmt.Errorf(logging.MakeLog("Ошибка загрузки .env", err))
	}

	logging.NewLogService(os.Stdout, os.Getenv("LOG_MODE"))
	logging.Logger.Debug(logging.MakeLog("Загрузки конфига", nil))

	if err := config.InitConfig(); err != nil {
		logging.Logger.Warn(logging.MakeLog("Ошибка инициализации конфига", err))
	}
	db, err := repository.NewPostgresDB(repository.Config{
		Host:     os.Getenv("DB_HOST"),
		Port:     viper.GetString("db.port"),
		Username: viper.GetString("db.username"),
		Password: os.Getenv("DB_PASSWORD"),
		DBName:   viper.GetString("db.dbname"),
		SSLMode:  viper.GetString("db.sslmode"),
	})
	if err != nil {
		logging.Logger.Warn(logging.MakeLog("Ошибка инициализации базы данных", err))
	}
	repos := repository.NewRepository(db)
	logging.Logger.Info("Инициализация репозитория")
	services := services.NewService(repos)
	logging.Logger.Info("Инициалиазация сервисов")
	handlers := handlers.NewHandler(services)
	logging.Logger.Info("Определение маршрутов")
	srv := new(models.Server)
	if err := srv.Run(viper.GetString("port"), handlers.InitRoutes()); err != nil {
		logging.Logger.Warn(err.Error())
	}
	logging.Logger.Info("Запущен сервер на порту :" + viper.GetString("port"))

}
