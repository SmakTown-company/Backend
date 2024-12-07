package handlers

import (
	"auth/database"
	"auth/models"
	"auth/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

func signIn(ctx *gin.Context) {

	var registerData models.RegisterData
	if err := ctx.ShouldBindJSON(&registerData); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Неверные данные"})
	}

	var user models.User
	result := database.DB.Where("phone = ?", registerData.Phone).First(&user)
	if result.Error != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Неверный телефон или пароль"})
		return
	}
	if !utils.CheckPasswordHash(registerData.Password, user.Hash) {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Неверный телефон или пароль"})
		return
	}
	tokens, err := utils.GenerateTokens(user.ID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Неудалось создать токен для пользователя"})
		return
	}

	// Создаем анонимную функцию с только id и phone
	userResponse := struct {
		ID    uint   `json:"id"`
		Phone string `json:"phone"`
	}{
		ID:    user.ID,
		Phone: user.Phone,
	}
	ctx.JSON(http.StatusOK, gin.H{"tokens": tokens, "user": userResponse})

}
