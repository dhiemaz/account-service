package access

import (
	"context"
	"github.com/dhiemaz/AccountService/infrastructures/logger"
	"github.com/dhiemaz/AccountService/internal/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type MongoAdapter struct {
	mgo *mongo.Collection
}

// NewAccessMongoAdapter access data with mongo adapter
func NewAccessMongoAdapter(mgo *mongo.Collection) *MongoAdapter {
	return &MongoAdapter{
		mgo: mgo,
	}
}

// GetAllAccessData from database
func (r *MongoAdapter) GetAllAccessData() ([]model.Access, error) {
	var datas []model.Access

	cursor, err := r.mgo.Find(context.TODO(), bson.M{})
	if err != nil {
		logger.WithFields(logger.Fields{"component": "access_repository", "action": "GetAllAccessData"}).
			Errorf("error fetching all access: %v", err)

		return datas, err
	}

	if err := cursor.All(context.TODO(), &datas); err != nil {
		logger.WithFields(logger.Fields{"component": "access_repository", "action": "GetAllAccessData"}).
			Errorf("error decoding cursor data: %v", err)

		return datas, err
	}

	return datas, nil
}
