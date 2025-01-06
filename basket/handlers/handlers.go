package handlers

import "github.com/gin-gonic/gin"

func AddInBasketHandler(ctx *gin.Context) {
	addInBasket(ctx)
}

func DeleteFromBasketHandler(ctx *gin.Context) {
	// Функция для удаления товара из корзину
}

func ShowAllProductsInBasketHandler(ctx *gin.Context) {
	// Функция для удаления товара из корзину
}

func PlaceAnOrderButtonHandler(ctx *gin.Context) {
	// Функция кнопки передачи данных на микросервис оформления заказа
}
