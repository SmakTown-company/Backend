package database

import (
	"auth/envs"
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// Объявление переменной DB, хранящей ссылку на экземпляр базы данных
var DB *gorm.DB

// Функция инициализации базы данных
func InitDatabase() error {
	env := envs.ServerEnvs
	// Формирование URI для подключения к PostgreSQL
	// Формируется с использованием значений ENV
	fmt.Print(env)
	uri := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s", env.POSTGRES_HOST, env.POSTGRES_USER, env.POSTGRES_PASSWORD, env.POSTGRES_NAME, env.POSTGRES_PORT, env.POSTGRES_USE_SSL)
	// Инициализация базы данных, в случае ошибки возвращается ошибка, иначе возвращается ссылка на экземпляр базы данных
	db, err := gorm.Open(postgres.Open(uri), &gorm.Config{})
	if err != nil {
		return err
	} else {
		DB = db
		return nil
	}
}
