package model

// Office : office data model
type Office struct {
	BranchId   int    `json:"branch_id" bson:"branch_id"`
	BranchName string `json:"branch_name" bson:"branch_name"`
}
