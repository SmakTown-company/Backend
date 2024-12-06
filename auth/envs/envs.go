package envs

import "os"

// Хранение данных значений ENV
var ServerEnvs Envs

// Структура для хранения значений ENV
type Envs struct {
	POSTGRES_PASSWORD string
	POSTGRES_USER     string
	POSTGRES_PORT     string
	JWT_SECRET        string
	AUTH_PORT         string
	POSTGRES_NAME     string
	POSTGRES_HOST     string
	POSTGRES_USE_SSL  string
}

// / Инициализация значений ENV
func LoadEnvs() error {
	ServerEnvs.JWT_SECRET = os.Getenv("JWT_SECRET")

	ServerEnvs.POSTGRES_USER = os.Getenv("POSTGRES_USER")
	ServerEnvs.POSTGRES_PASSWORD = os.Getenv("POSTGRES_PASSWORD")
	ServerEnvs.POSTGRES_PORT = os.Getenv("POSTGRES_PORT")
	ServerEnvs.POSTGRES_NAME = os.Getenv("POSTGRES_NAME")
	ServerEnvs.POSTGRES_HOST = os.Getenv("POSTGRES_HOST")
	ServerEnvs.POSTGRES_USE_SSL = os.Getenv("POSTGRES_USE_SSL")
	ServerEnvs.AUTH_PORT = os.Getenv("AUTH_PORT")

	return nil
}
