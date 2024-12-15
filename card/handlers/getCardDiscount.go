package handlers

import (
	"card/database"
	"card/models"
	"fmt"
	"log" // Для логирования
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
)

// Функция для получения карт со скидкой
func getCardDiscountHandler(c *gin.Context) {
	log.Println("Получен запрос на получение карт со скидкой") // Логируем начало запроса

	// Получаем коллекцию из глобального клиента MongoDB
	collection := database.CardCollection

	// Проверка на nil для коллекции
	if collection == nil {
		log.Println("CardCollection не инициализирована") // Логируем ошибку
		c.JSON(http.StatusInternalServerError, gin.H{"error": "CardCollection не инициализирована"})
		return
	}

	// Получаем контекст из запроса
	ctx := c.Request.Context()

	// Фильтр для получения только тех карт, у которых DiscountStatus = true
	filter := bson.D{
		{"discountStatus", true}, // Добавляем фильтрацию по полю DiscountStatus
	}

	// Получаем курсор из коллекции
	cursor, err := collection.Find(ctx, filter)
	if err != nil {
		log.Printf("Ошибка при извлечении данных из базы данных: %v", err) // Логируем ошибку
		c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Ошибка при извлечении данных из базы данных: %v", err)})
		return
	}
	defer cursor.Close(ctx)

	// Создаем слайс для хранения карточек
	var cards []models.Card
	for cursor.Next(ctx) {
		var card models.Card
		if err := cursor.Decode(&card); err != nil {
			log.Printf("Ошибка декодирования данных: %v", err) // Логируем ошибку
			c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Ошибка декодирования данных: %v", err)})
			return
		}
		cards = append(cards, card)
	}

	// Проверяем на ошибки при итерации по курсу
	if err := cursor.Err(); err != nil {
		log.Printf("Ошибка в курсоре: %v", err) // Логируем ошибку
		c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Ошибка в курсоре: %v", err)})
		return
	}

	log.Printf("Получено %d карт со скидкой", len(cards)) // Логируем количество карт со скидкой

	// Отправляем данные в формате JSON
	c.JSON(http.StatusOK, cards)
}
