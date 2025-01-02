package repository

import (
	"fmt"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file" // Добавлено для поддержки файлового источника
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type Config struct {
	Host     string
	Port     string
	Username string
	Password string
	DBName   string
	SSLMode  string
}

func NewPostgresDB(cfg Config) (*sqlx.DB, error) {
	// Подключение к базе "postgres" для проверки и создания основной БД
	connStr := fmt.Sprintf("host=%s port=%s user=%s dbname=postgres password='%s' sslmode=%s", cfg.Host, cfg.Port, cfg.Username, cfg.Password, cfg.SSLMode)
	initialDB, err := sqlx.Open("postgres", connStr)
	if err != nil {
		return nil, fmt.Errorf("Ошибка подключения к базе данных postgres")
	}
	_, err = initialDB.Exec(fmt.Sprintf("CREATE DATABASE %s", cfg.DBName))
	if err != nil {
		err = nil
	}
	defer func() {
		initialDB.Close()
	}()
	mainConnStr := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password='%s' sslmode=%s", cfg.Host, cfg.Port, cfg.Username, cfg.DBName, cfg.Password, cfg.SSLMode)
	db, err := sqlx.Open("postgres", mainConnStr)
	if err != nil {
		return nil, fmt.Errorf("Ошибка подключения к базе данных %s", cfg.DBName)
	}
	if err = db.Ping(); err != nil {
		return nil, fmt.Errorf("Ошибка пингования базы данных")
	}
	err = Migrate(db, cfg)
	if err != nil {
		return nil, err
	}
	return db, nil
}

func Migrate(db *sqlx.DB, cfg Config) error {
	driver, err := postgres.WithInstance(db.DB, &postgres.Config{})
	if err != nil {
		return fmt.Errorf("Ошибка создания объекта драйвера")
	}
	m, err := migrate.NewWithDatabaseInstance(
		"file://schema",
		cfg.DBName, driver,
	)
	if err != nil {
		return fmt.Errorf("Ошибка создания объекта миграция")
	}

	if err := m.Up(); err != nil && err != migrate.ErrNoChange {
		return fmt.Errorf("Ошибка принятия миграции")
	}
	return nil
}
