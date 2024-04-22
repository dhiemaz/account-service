package mongo

import (
	"context"
	"github.com/dhiemaz/AccountService/constant"
	"github.com/dhiemaz/AccountService/infrastructures/logger"
	"github.com/dhiemaz/AccountService/internal/domain/role/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type MongoAdapter struct {
	mgo *mongo.Database
}

// NewRoleMongoAdapter role access with mongo adapter
func NewRoleMongoAdapter(mgo *mongo.Database) *MongoAdapter {
	return &MongoAdapter{
		mgo: mgo,
	}
}

// GetAllRoleData from database
func (r *MongoAdapter) GetAllRoleData() ([]model.Role, error) {
	var datas []model.Role

	collection := r.mgo.Collection(constant.RoleCollection)

	cursor, err := collection.Find(context.TODO(), bson.M{})
	if err != nil {
		logger.WithFields(logger.Fields{"component": "role_repository", "action": "GetAllRoleData"}).
			Errorf("error fetching all roles: %v", err)

		return datas, err
	}

	if err := cursor.All(context.TODO(), &datas); err != nil {
		logger.WithFields(logger.Fields{"component": "role_repository", "action": "GetAllRoleData"}).
			Errorf("error decoding cursor data: %v", err)

		return datas, err
	}

	return datas, nil
}
