package handlers

import (
	"auth/database"
	"auth/models"
	"auth/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

func registerUserHandler(ctx *gin.Context) {
	var user models.User
	var registerData models.RegisterData

	// Парсим JSON тело запроса в структуру RegisterData
	if err := ctx.ShouldBind(&registerData); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Неверные данные"})
		return
	}

	// Хэшируем пароль
	hashedPassword, err := utils.HashPassword(registerData.Password)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка при хэшировании пароля"})
		return
	}
	user.Email = registerData.Email
	user.Hash = hashedPassword
	// Сохраняем пользователя в базу данных
	result := database.DB.Create(&user)
	if result.Error != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Не удалось сохранить пользователя"})
		return
	} else {
		ctx.JSON(http.StatusOK, gin.H{"message": "Регистрация пользователя прошла успешно"})
	}

}
