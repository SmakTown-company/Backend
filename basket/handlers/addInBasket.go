package handlers

import (
	"basket/database"
	"basket/models"
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func addInBasket(ctx *gin.Context) {

	urlforCookie := "http://localhost:9103/SmakTown/API/checkCookie"

	responseForUser, errForUser := http.Get(urlforCookie)
	if errForUser != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Не удалось получить данные из coockie"})
		return
	}

	var userData struct {
		ID       uint `json:"user_id"`
		Verified bool `json:"verified"`
	}

	errForUser = json.NewDecoder(responseForUser.Body).Decode(&userData)
	if errForUser != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Не удалось изъять информацию о пользователе"})
		return
	}

	basketUser := models.UserInfo{
		ID:       userData.ID,
		Verified: userData.Verified,
	}

	var requestData struct {
		ProductID uint `json:"product_id"`
	}
	if err := ctx.ShouldBindJSON(&requestData); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Неверный формат запроса"})
		return
	}

	urlforProduct := fmt.Sprintf("http://localhost:9101/SmakTown/API/cardForBasket/%d", requestData.ProductID)
	responseforProduct, errForProduct := http.Get(urlforProduct)
	if errForProduct != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Не удалось получить данные из каталога товаров"})
		return
	}

	var productData struct {
		ID             uint    `json:"id" bson:"id"`
		Image          string  `json:"img" bson:"img"`
		NameProduct    *string `json:"nameproduct" bson:"nameproduct"`
		Content        *string `json:"content" bson:"content"`
		Price          *string `json:"price" bson:"price"`
		ShopID         string  `json:"shop_id" bson:"shop_id"`
		DiscountStatus bool    `json:"discountStatus,omitempty" bson:"discountStatus,omitempty"`
	}

	errForProduct = json.NewDecoder(responseforProduct.Body).Decode(&productData)
	if errForProduct != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Не удалось изъять информацию о товаре"})
		return
	}

	basketProduct := models.Card{
		ID:             productData.ID,
		Image:          productData.Image,
		NameProduct:    productData.NameProduct,
		Content:        productData.Content,
		Price:          productData.Price,
		ShopID:         productData.ShopID,
		DiscountStatus: productData.DiscountStatus,
	}

	basket := models.BasketUser{
		UserInfo: basketUser,
		Card:     basketProduct,
	}

	_, err := database.BasketCollection.InsertOne(context.Background(), basket)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Не удалось добавить в корзину"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Товар успешно добавлен в корзину", "basket": basket})

}
