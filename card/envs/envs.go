package envs

import (
	"os"
)

// Хранение данных значений ENV
var ServerEnvs Envs

// Структура для хранения значений ENV
type Envs struct {
	MONGO_INITDB_ROOT_PASSWORD_CARD string
	MONGO_INITDB_ROOT_USERNAME_CARD string
	MONGO_INITDB_PORT_CARD          string
	MONGO_INITDB_HOST_CARD          string
	CARD_PORT                       string
	JWT_SECRET                      string
}

// Инициализация значений ENV
func LoadEnvs() error {

	// Инициализация значений ENV
	ServerEnvs.CARD_PORT = os.Getenv("CARD_PORT")
	ServerEnvs.MONGO_INITDB_ROOT_USERNAME_CARD = os.Getenv("MONGO_INITDB_ROOT_USERNAME")
	ServerEnvs.MONGO_INITDB_ROOT_PASSWORD_CARD = os.Getenv("MONGO_INITDB_ROOT_PASSWORD")
	ServerEnvs.MONGO_INITDB_PORT_CARD = os.Getenv("MONGO_INITDB_PORT")
	ServerEnvs.MONGO_INITDB_HOST_CARD = os.Getenv("MONGO_INITDB_HOST")
	ServerEnvs.JWT_SECRET = os.Getenv("JWT_SECRET")

	return nil
}
