package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Account : account data model
type Account struct {
	Id           string             `json:"id" bson:"_id"`
	Username     string             `json:"username" bson:"username"`
	Password     string             `json:"password" bson:"password"`
	Fullname     string             `json:"fullname" bson:"fullname"`
	Address      string             `json:"address" bson:"address"`
	zipcode      string             `json:"zipcode" bson:"zipcode"`
	province     string             `json:"province" bson:"province"`
	RegisterDate primitive.DateTime `json:"register_date" bson:"register_date"`
	Office       []Office           `json:"office" bson:"office"`
	Role         []Role             `json:"role" bson:"role"`
	Access       []Access           `json:"access" bson:"access"`
}
