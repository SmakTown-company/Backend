package handlers

import (
	"auth/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

type RefreshTokenRequest struct {
	RefreshToken string `json:"refresh_token"`
}

func refreshToken(ctx *gin.Context) {

	var token RefreshTokenRequest

	// Извлекаем данные из тела запроса и заполняем в структуру
	if err := ctx.ShouldBindJSON(&token); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Ошибка при разборе тела запроса"})
		return
	}
	userId, err := utils.ValidateRefreshToken(token.RefreshToken)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Неверный refresh токен"})
		return
	}

	tokens, err := utils.GenerateTokens(userId)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Не удалось создать токены"})
		return
	}

	ctx.JSON(http.StatusOK, tokens)

}
