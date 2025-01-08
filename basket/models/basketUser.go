package models

type BasketUser struct {
	UserInfo `json:"user" bson:"user"`
	Card     `json:"card" bson:"card"`
}
