package handlers

import (
	"card/database"
	"card/models"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
)

func getCardHandler(c *gin.Context) {
	// Получаем коллекцию из глобального клиента MongoDB
	collection := database.CardCollection
	ctx := c.Request.Context()

	// фильтр, например, получаем все записи
	filter := bson.D{{}} // Можете указать конкретные фильтры

	// Получаем курсор из коллекции
	cursor, err := collection.Find(ctx, filter)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Ошибка при извлечении данных из базы данных: %v", err)})
		return
	}
	defer cursor.Close(ctx)

	// Создаем слайс для хранения карточек
	var cards []models.Card
	for cursor.Next(ctx) {
		var card models.Card
		if err := cursor.Decode(&card); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Ошибка декодирования данных: %v", err)})
			return
		}
		cards = append(cards, card)
	}

	// Проверяем на ошибки при итерации по курсу
	if err := cursor.Err(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Ошибка в курсоре: %v", err)})
		return
	}

	// Отправляем данные в формате JSON
	c.JSON(http.StatusOK, cards)
}
