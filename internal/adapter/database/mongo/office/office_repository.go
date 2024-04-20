package office

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

// NewOfficeMongoAdapter office data with mongo adapter
func NewOfficeMongoAdapter(mgo *mongo.Collection) *MongoAdapter {
	return &MongoAdapter{
		mgo: mgo,
	}
}

// GetAllOfficeData from database
func (r *MongoAdapter) GetAllOfficeData() ([]model.Office, error) {
	var datas []model.Office

	cursor, err := r.mgo.Find(context.TODO(), bson.M{})
	if err != nil {
		logger.WithFields(logger.Fields{"component": "office_repository", "action": "GetAllOfficeData"}).
			Errorf("error fetching all offices: %v", err)

		return datas, err
	}

	if err := cursor.All(context.TODO(), &datas); err != nil {
		logger.WithFields(logger.Fields{"component": "office_repository", "action": "GetAllOfficeData"}).
			Errorf("error decoding cursor data: %v", err)

		return datas, err
	}

	return datas, nil
}
