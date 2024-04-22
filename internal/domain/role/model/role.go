package model

// Role : role data model
type Role struct {
	RoleId   int    `json:"json_id" bson:"role_id"`
	RoleName string `json:"role_name" bson:"role_name"`
}
