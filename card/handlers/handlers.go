package handlers

import "github.com/gin-gonic/gin"

func GetCardHandler(ctx *gin.Context) {
	getCardHandler(ctx)
}

func AddInBasket(ctx *gin.Context) {
	addCardToBasket(ctx)
}
