package model

// Access : access data model
type Access struct {
	AccessId   int    `json:"access_id" bson:"access_id"`
	AccessName string `json:"access_name" bson:"access_name"`
}
