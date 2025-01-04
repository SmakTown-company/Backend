package handlers

import (
	"auth/database"
	"auth/models"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

// Обработчик для подтверждения токена
func verifyToken(ctx *gin.Context) {
	token := ctx.Query("token") // Получаем токен из параметров запроса

	// Ищем токен в базе данных
	var verificationToken models.VerificationToken
	if err := database.DB.Where("token = ?", token).First(&verificationToken).Error; err != nil {
		ctx.JSON(http.StatusBadRequest, "Неверный или несуществующий токен")
		return
	}

	// Проверяем, не истек ли срок действия токена
	if time.Now().After(verificationToken.ExpiresAt) {
		ctx.JSON(http.StatusBadRequest, "Срок действия токена истек")
		return
	}

	// Подтверждаем токен
	confirmedAt := time.Now().Unix()
	verificationToken.ConfirmedAt = &confirmedAt

	// Обновляем запись в базе данных
	if err := database.DB.Save(&verificationToken).Error; err != nil {
		ctx.JSON(http.StatusInternalServerError, "Ошибка при подтверждении токена")
		return
	}
	ctx.JSON(http.StatusOK, "Код подтверждения успешен")
}
