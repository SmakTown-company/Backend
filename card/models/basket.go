package models

type BasketItem struct {
	CardID      string `json:"card_id" bson:"card_id"`
	NameProduct string `json:"nameproduct" bson:"nameproduct"`
	Price       string `json:"price" bson:"price"`
	Quantity    int    `json:"quantity" bson:"quantity"`
}

type Basket struct {
	UserID string       `json:"user_id" bson:"user_id"`
	Items  []BasketItem `json:"items" bson:"items"`
}
