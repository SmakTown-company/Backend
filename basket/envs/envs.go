package envs

import (
	"os"
)

// Хранение данных значений ENV
var ServerEnvs Envs

// Структура для хранения значений ENV
type Envs struct {
	MONGO_INITDB_ROOT_PASSWORD_BASKET string
	MONGO_INITDB_ROOT_USERNAME_BASKET string
	MONGO_INITDB_PORT_BASKET          string
	MONGO_INITDB_HOST_BASKET          string
	BASKET_PORT                       string
	JWT_SECRET                        string
}

// Инициализация значений ENV
func LoadEnvs() error {

	// Инициализация значений ENV
	ServerEnvs.BASKET_PORT = os.Getenv("BASKET_PORT")
	ServerEnvs.MONGO_INITDB_ROOT_USERNAME_BASKET = os.Getenv("MONGO_INITDB_ROOT_USERNAME")
	ServerEnvs.MONGO_INITDB_ROOT_PASSWORD_BASKET = os.Getenv("MONGO_INITDB_ROOT_PASSWORD")
	ServerEnvs.MONGO_INITDB_PORT_BASKET = os.Getenv("MONGO_INITDB_PORT")
	ServerEnvs.MONGO_INITDB_HOST_BASKET = os.Getenv("MONGO_INITDB_HOST")
	ServerEnvs.JWT_SECRET = os.Getenv("JWT_SECRET")

	return nil
}
