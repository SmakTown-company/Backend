package database

import (
	"card/envs"
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

// Объявление переменной MongoClient, хранящей ссылку на экземпляр клиента MongoDB
var MongoClient *mongo.Client
var CardCollection *mongo.Collection

// Инициализация подключения к MongoDB
func InitDatabase() error {
	// Загружаем данные окружения из структуры envs
	env := &envs.ServerEnvs
	// Формируем URI для подключения к MongoDB
	mongoURI := fmt.Sprintf("mongodb://%s:%s@%s:%s", env.MONGO_INITDB_ROOT_USERNAME, env.MONGO_INITDB_ROOT_PASSWORD, env.MONGO_INITDB_HOST, env.MONGO_INITDB_PORT)
	log.Println("URI: " + mongoURI)

	// Создаем новый контекст с таймаутом и предусматриваем его корректное завершение
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Создаем клиента MongoDB и пытаемся подключиться
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(mongoURI))
	if err != nil {
		return fmt.Errorf("Ошибка при подключении к MongoDB: %v", err)
	}

	// Проверяем подключение к базе данных
	if err := client.Ping(ctx, readpref.Primary()); err != nil {
		return fmt.Errorf("Ошибка при подключении к MongoDB: %v", err)
	}

	// Сохраняем клиента в глобальной переменной
	MongoClient = client

	// Инициализируем коллекцию для работы
	CardCollection = MongoClient.Database("yourDatabase").Collection("cards")

	log.Println("Успешное подключение к MongoDB")
	return nil
}

// Закрытие соединения с MongoDB
func CloseDatabase() {
	if err := MongoClient.Disconnect(context.Background()); err != nil {
		log.Fatal("Ошибка при подключении к MongoDB:", err)
	}
	log.Println("MongoDB завершило работу.")
}
