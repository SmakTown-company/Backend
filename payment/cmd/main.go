package main

import (
	"context"
	"fmt"
	"github.com/SmakTown-company/Backend/payment/internal/config"
	"github.com/SmakTown-company/Backend/payment/internal/handler"
	"github.com/SmakTown-company/Backend/payment/internal/model"
	"github.com/SmakTown-company/Backend/payment/internal/repository"
	"github.com/SmakTown-company/Backend/payment/internal/service"
	"github.com/SmakTown-company/Backend/payment/pkg/logging"
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
	db, err := repository.NewRedis(repository.Config{
		Port:     viper.GetString("db.port"),
		Host:     viper.GetString("db.host"),
		Password: os.Getenv("DB_PASSWORD"),
		DB:       viper.GetInt("db.db"),
		Protocol: viper.GetInt("db.protocol"),
	})
	if err != nil {
		logging.Logger.Warn(logging.MakeLog("Ошибка инициализации базы данных", err))
	}
	ctx := context.Background()
	repos := repository.NewRepository(db, ctx)
	logging.Logger.Info("Инициализация репозитория")
	services := service.NewService(viper.GetString("payment.method"), viper.GetString("payment.return_url"), *repos, ctx)
	logging.Logger.Info("Инициалиазация сервисов")
	handlers := handler.NewHandler(services)
	logging.Logger.Info("Определение маршрутов")
	srv := new(model.Server)
	if err := srv.Run(viper.GetString("port"), handlers.InitRoutes()); err != nil {
		logging.Logger.Warn(err.Error())
	}
	logging.Logger.Info("Запущен сервер на порту :" + viper.GetString("port"))

}
