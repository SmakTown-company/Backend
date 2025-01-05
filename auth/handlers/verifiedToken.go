package handlers

import (
	"auth/database"
	"auth/models"
	"auth/utils"
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

	// Проверка на действительность токена
	if !utils.CheckTokenExpiration(token) {
		// Генерация нового токена
		newToken := utils.GenerateVerifiedToken()

		// Обновление данных в базе с новым токеном
		verificationToken.Token = newToken
		verificationToken.CreatedAt = time.Now()
		verificationToken.ExpiresAt = time.Now().Add(time.Minute * 10) // Устанавливаем новый срок действия токена

		if err := database.DB.Save(&verificationToken).Error; err != nil {
			ctx.JSON(http.StatusInternalServerError, "Не удалось обновить токен в базе данных")
			return
		}

		ctx.JSON(http.StatusOK, gin.H{"message": "Время жизни токена истекло, новый токен сгенерирован."})
		return
	}

	// Подтверждаем токен
	confirmedAt := time.Now()
	verificationToken.ConfirmedAt = confirmedAt

	// Обновляем запись в базе данных
	if err := database.DB.Save(&verificationToken).Error; err != nil {
		ctx.JSON(http.StatusInternalServerError, "Ошибка при подтверждении токена")
		return
	}
	// Обновляем пользователю статус verified на true в случае удачной верификации
	var user models.User
	if err := database.DB.Where("email = ? OR phone = ?", verificationToken.Email, verificationToken.Phone).First(&user).Error; err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Не удалось найти пользователя для подтверждения верификации"})
		return
	}
	user.Verified = true
	if err := database.DB.Save(&user).Error; err != nil {
		ctx.JSON(http.StatusInternalServerError, "Ошибка при обновлении статуса верификации")
		return
	}

	// Удаляем запись о токене по id физически, игнорируя логику мягкого удаления
	if err := database.DB.Unscoped().Where("id = ?", verificationToken.ID).Delete(&models.VerificationToken{}).Error; err != nil {
		ctx.JSON(http.StatusInternalServerError, "Ошибка при удалении записи о токене")
		return
	}

	ctx.JSON(http.StatusOK, "Код подтверждения успешен, статус обновлен")
}
