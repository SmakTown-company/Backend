package handlers

import (
	"auth/database"
	"auth/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// Функция для получения пользователя для корзины
func getUserForBasket(ctx *gin.Context) {
	// Извлекаем user_id из параметров URL
	userID := ctx.Param("user_id")

	if userID == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "user_id is required",
		})
		return
	}

	// Ищем пользователя в базе данных по user_id
	var user models.User
	if err := database.DB.Where("id = ?", userID).First(&user).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			ctx.JSON(http.StatusNotFound, gin.H{
				"error": "Пользователь не найден",
			})
		} else {
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"error": "не удалось выполнить запрос к базе данных",
			})
		}
		return
	}

	// Возвращаем успешный ответ с найденным пользователем
	ctx.JSON(http.StatusOK, gin.H{
		"user_id": user.ID,
	})
}
