package handlers

import (
	"auth/database"
	"auth/models"
	"auth/utils"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func signIn(ctx *gin.Context) {

	var signInData struct {
		Email    string `json:"email" binding:"required"`
		Password string `json:"password" binding:"required"`
	}

	// Получаем данные из запроса
	if err := ctx.ShouldBindJSON(&signInData); err != nil {
		log.Printf("Введены неверные данные: %v", err) // Логируем ошибку, если данные невалидны
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Неверные данные"})
		return
	}
	log.Printf("Received sign-in data: %+v", signInData) // Логируем полученные данные

	// Ищем пользователя по номеру телефона
	var user models.User
	result := database.DB.Where("email = ?", signInData.Email).First(&user)
	if result.Error != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Неверный телефон или пароль"})
		return
	}

	// Проверяем, совпадает ли пароль с сохраненным хешом
	if !utils.CheckPasswordHash(signInData.Password, user.Hash) {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Неверный телефон или пароль"})
		return
	}

	// Генерируем токены
	tokens, err := utils.GenerateTokens(user.ID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Не удалось создать токен для пользователя"})
		return
	}

	// Создаем анонимную структуру для ответа
	userResponse := struct {
		ID    uint   `json:"id"`
		Email string `json:"email"`
	}{
		ID:    user.ID,
		Email: user.Email,
	}

	// Отправляем успешный ответ
	ctx.JSON(http.StatusOK, gin.H{"tokens": tokens, "user": userResponse})
}
