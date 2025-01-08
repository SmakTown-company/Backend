package handlers

import (
	"card/database"
	"card/models"
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
)

func cardForBasket(ctx *gin.Context) {

	cardID := ctx.Param("id")

	var card models.Card

	filter := bson.M{"id": cardID}
	err := database.CardCollection.FindOne(context.Background(), filter).Decode(&card)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Не удалось найти карту в БД"})
		return
	}

	ctx.JSON(http.StatusOK, card)

}
