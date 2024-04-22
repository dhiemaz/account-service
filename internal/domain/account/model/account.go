package model

import (
	accessModel "github.com/dhiemaz/AccountService/internal/domain/access/model"
	officeModel "github.com/dhiemaz/AccountService/internal/domain/office/model"
	roleModel "github.com/dhiemaz/AccountService/internal/domain/role/model"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Account : account data model
type Account struct {
	Id           string               `json:"id" bson:"_id"`
	Username     string               `json:"username" bson:"username"`
	Password     string               `json:"password" bson:"password"`
	Fullname     string               `json:"fullname" bson:"fullname"`
	Address      string               `json:"address" bson:"address"`
	Zipcode      string               `json:"zipcode" bson:"zipcode"`
	Province     string               `json:"province" bson:"province"`
	RegisterDate primitive.DateTime   `json:"register_date" bson:"register_date"`
	Office       []officeModel.Office `json:"office" bson:"office"`
	Role         []roleModel.Role     `json:"role" bson:"role"`
	Access       []accessModel.Access `json:"access" bson:"access"`
}
