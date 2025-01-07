package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func cheackCookie(ctx *gin.Context) {
	userID, err := ctx.Cookie("user_id")
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, "Не удалось узнать id пользователя из coockie")
	}
	verified, err := ctx.Cookie("verified")
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, "Не удалось узнать verified пользователя из cookie")
	}

	ctx.JSON(http.StatusOK, gin.H{"user_id": userID, "verified": verified})

}
