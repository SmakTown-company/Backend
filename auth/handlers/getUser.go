package handlers

import (
	"auth/database"
	"auth/models"
	"auth/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

func getUser(ctx *gin.Context) {
	userID, err := utils.ExtractUserID(ctx.GetHeader("Authorization"))
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Неверный токен:" + err.Error()})
	}
	var user models.User
	result := database.DB.Where("ID = ?", userID).First(&user)
	if result.Error != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Пользователь не найден"})
		return
	}
	// Создаем анонимную структуру с id и email
	userResponse := struct {
		ID    uint   `json:"id"`
		Email string `json:"email"`
	}{
		ID:    userID,
		Email: user.Email,
	}
	ctx.JSON(http.StatusOK, gin.H{"user": userResponse})
}
