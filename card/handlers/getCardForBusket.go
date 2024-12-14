package handlers

import (
	"card/database"
	"card/models"
	"context"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
)

// Функция для добавления карты в корзину пользователя
func addCardToBasket(ctx *gin.Context) {
	// Извлекаем user_id из контекста
	userID, exists := ctx.Get("user_id")
	if !exists {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "user_id is required",
		})
		return
	}

	// Преобразуем user_id в строку
	userIDStr, ok := userID.(string)
	if !ok {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid user_id type",
		})
		return
	}

	// Проверка на инициализацию коллекции
	if database.CardCollection == nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": "CardCollection не инициализирована",
		})
		return
	}

	// Извлекаем данные карты из тела запроса
	var card models.Card
	if err := ctx.BindJSON(&card); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": fmt.Sprintf("invalid card data: %v", err),
		})
		return
	}

	// Получаем информацию о карте (например, из базы данных MongoDB)
	cardData := models.BasketItem{
		CardID:      card.Id,
		NameProduct: *card.NameProduct,
		Price:       *card.Price,
		Quantity:    1, // Начинаем с 1 штуки
	}

	// Ищем корзину для этого пользователя
	var existingBasket models.Basket
	err := database.BasketCollection.FindOne(context.Background(), bson.M{"user_id": userIDStr}).Decode(&existingBasket)

	if err != nil {
		if err.Error() == "mongo: no documents in result" {
			// Если корзина не найдена, создаем новую
			newBasket := models.Basket{
				UserID: userIDStr,
				Items:  []models.BasketItem{cardData},
			}

			// Сохраняем корзину в базе данных
			_, err := database.BasketCollection.InsertOne(context.Background(), newBasket)
			if err != nil {
				ctx.JSON(http.StatusInternalServerError, gin.H{
					"error": fmt.Sprintf("failed to create basket: %v", err),
				})
				return
			}
			ctx.JSON(http.StatusOK, gin.H{
				"message": "Карта добавлена в корзину",
				"basket":  newBasket,
			})
			return
		}
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": fmt.Sprintf("failed to find basket: %v", err),
		})
		return
	}

	// Если корзина найдена, обновляем или добавляем новый элемент
	updated := false
	for i, item := range existingBasket.Items {
		if item.CardID == card.Id {
			// Если такая карта уже есть, увеличиваем количество
			existingBasket.Items[i].Quantity++
			updated = true
			break
		}
	}

	if !updated {
		// Если карты еще нет, добавляем новый элемент
		existingBasket.Items = append(existingBasket.Items, cardData)
	}

	// Обновляем корзину в базе данных
	_, err = database.BasketCollection.UpdateOne(
		context.Background(),
		bson.M{"user_id": userIDStr},
		bson.M{"$set": bson.M{"items": existingBasket.Items}},
	)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": fmt.Sprintf("failed to update basket: %v", err),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "Карта успешно добавлена",
		"basket":  existingBasket,
	})
}
