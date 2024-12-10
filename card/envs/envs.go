package envs

import (
	"os"
)

// Хранение данных значений ENV
var ServerEnvs Envs

// Структура для хранения значений ENV
type Envs struct {
	MONGO_INITDB_ROOT_PASSWORD string
	MONGO_INITDB_ROOT_USERNAME string
	MONGO_INITDB_PORT          string
	MONGO_INITDB_HOST          string
	CARD_PORT                  string
	JWT_SECRET                 string
}

// Инициализация значений ENV
func LoadEnvs() error {

	// Инициализация значений ENV
	ServerEnvs.CARD_PORT = os.Getenv("CARD_PORT")
	ServerEnvs.MONGO_INITDB_ROOT_USERNAME = os.Getenv("MONGO_INITDB_ROOT_USERNAME")
	ServerEnvs.MONGO_INITDB_ROOT_PASSWORD = os.Getenv("MONGO_INITDB_ROOT_PASSWORD")
	ServerEnvs.MONGO_INITDB_PORT = os.Getenv("MONGO_INITDB_PORT")
	ServerEnvs.MONGO_INITDB_HOST = os.Getenv("MONGO_INITDB_HOST")
	ServerEnvs.JWT_SECRET = os.Getenv("JWT_SECRET")

	return nil
}
