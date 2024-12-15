package handlers

import "github.com/gin-gonic/gin"

func GetCardHandler(ctx *gin.Context) {
	getCardHandler(ctx)
}

func GetCardDiscountHandler(ctx *gin.Context) {
	getCardDiscountHandler(ctx)
}

func AddInBasket(ctx *gin.Context) {
	// функция для добавления в корзину
}
