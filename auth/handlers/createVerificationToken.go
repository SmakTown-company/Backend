package handlers

import (
	"auth/database"
	"auth/models"
	"auth/utils"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

// Обработчик для создания записи с токеном подтверждения
func createVerificationToken(ctx *gin.Context) {
	var input struct {
		Email string `json:"email"`
		Phone string `json:"phone"`
	}

	// Разбор входных данных
	if err := ctx.BindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, "Неправильный формат данных")
		return
	}

	// Проверяем, существует ли пользователь с таким email или phone
	var user models.User
	if err := database.DB.Where("email = ? OR phone = ?", input.Email, input.Phone).First(&user).Error; err != nil {
		ctx.JSON(http.StatusBadRequest, "Пользователь с таким email или телефоном не найден")
		return
	}

	var existingToken models.VerificationToken
	err := database.DB.Where("email = ? OR phone = ?", input.Email, input.Phone).Where("confirmed_at IS NULL").First(&existingToken).Error

	if err == nil {
		if time.Now().After(existingToken.ExpiresAt) {
			if err := database.DB.Delete(&existingToken).Error; err != nil {
				ctx.JSON(http.StatusInternalServerError, "Ошибка при удалении старого токена")
				return
			}
		} else {
			ctx.JSON(http.StatusBadRequest, "Для данного пользователя уже существует неподтвержденный токен")
		}
	}

	// Генерация токена
	token := utils.GenerateVerifiedToken()

	// Время истечения токена
	expiration := time.Now().Add(10 * time.Minute)

	// Создание записи в таблице verification_tokens
	verificationToken := models.VerificationToken{
		Email:     input.Email,
		Phone:     input.Phone,
		Token:     token,
		ExpiresAt: expiration,
	}

	// Сохранение токена в базе данных
	if err := database.DB.Create(&verificationToken).Error; err != nil {
		ctx.JSON(http.StatusInternalServerError, "Ошибка при сохранении токена в базе данных")
		return
	}

	ctx.JSON(http.StatusOK, "Код подтверждения отправлен")
}
