package models

type Card struct {
	Id             string  `json:"id,omitempty" bson:"id,omitempty"`
	Image          string  `json:"img" bson:"img"`
	NameProduct    *string `json:"nameproduct,omitempty" bson:"nameproduct,omitempty"`
	Content        *string `json:"content,omitempty" bson:"content,omitempty"`
	Price          *string `json:"price,omitempty" bson:"price,omitempty"`
	ShopID         string  `json:"shop_id,omitempty" bson:"shop_id,omitempty"`
	DiscountStatus bool    `json:"discountStatus,omitempty" bson:"discountStatus,omitempty"`
}
