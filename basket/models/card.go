package models

type Card struct {
	ID             uint    `json:"id" bson:"id"`
	Image          string  `json:"img" bson:"img"`
	NameProduct    *string `json:"nameproduct" bson:"nameproduct"`
	Content        *string `json:"content" bson:"content"`
	Price          *string `json:"price" bson:"price"`
	ShopID         string  `json:"shop_id" bson:"shop_id"`
	DiscountStatus bool    `json:"discountStatus,omitempty" bson:"discountStatus,omitempty"`
}
