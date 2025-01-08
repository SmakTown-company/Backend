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

// Объявление коллекций
var CardCollection *mongo.Collection

// Инициализация подключения к MongoDB
func InitDatabase() error {
	env := &envs.ServerEnvs

	// Формируем URI для подключения к MongoDB
	mongoURI := fmt.Sprintf("mongodb://%s:%s@%s:%s", env.MONGO_INITDB_ROOT_USERNAME_CARD, env.MONGO_INITDB_ROOT_PASSWORD_CARD, env.MONGO_INITDB_HOST_CARD, env.MONGO_INITDB_PORT_CARD)
	log.Println("URI: " + mongoURI)

	// Увеличиваем таймаут подключения
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	// Подключаемся к MongoDB
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(mongoURI))
	if err != nil {
		return fmt.Errorf("Ошибка при подключении к MongoDB: %v", err)
	}

	// Проверяем подключение
	if err := client.Ping(ctx, readpref.Primary()); err != nil {
		return fmt.Errorf("Ошибка при подключении к MongoDB: %v", err)
	}

	// Сохраняем клиента и инициализируем коллекции
	MongoClient = client
	CardCollection = MongoClient.Database("card_db").Collection("card")

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
