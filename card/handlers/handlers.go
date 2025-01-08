package handlers

import "github.com/gin-gonic/gin"

func GetCardHandler(ctx *gin.Context) {
	getCardHandler(ctx)
}

func GetCardDiscountHandler(ctx *gin.Context) {
	getCardDiscountHandler(ctx)
}

func CardForBasket(ctx *gin.Context) {
	cardForBasket(ctx)
}
