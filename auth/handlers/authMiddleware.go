package handlers

import (
	"net/http"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

// AuthMiddleware для проверки токена в заголовке Authorization
func AuthMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		tokenValue := strings.Split(ctx.GetHeader("Authorization"), " ")
		// Проверка наличия заголовка Authorization
		if len(tokenValue) != 2 {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Неверный формат заголовка Authorized"})
			return
		}
		accessToken, err := jwt.Parse(tokenValue[1],
			func(token *jwt.Token) (interface{}, error) {
				return []byte(os.Getenv("JWT_SECRET")), nil
			})
		if err != nil || accessToken == nil || !accessToken.Valid {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Неверный токен"})
		}
		ctx.Next()

	}
}
