package handlers

import (
	"github.com/gin-gonic/gin"
)

func SignInHandler(ctx *gin.Context) {
	signIn(ctx)
}
func RefreshTokenHandler(ctx *gin.Context) {
	refreshToken(ctx)
}
func GetUserHandler(ctx *gin.Context) {
	getUser(ctx)
}

func RegisterUserHandler(ctx *gin.Context) {
	registerUserHandler(ctx)
}
