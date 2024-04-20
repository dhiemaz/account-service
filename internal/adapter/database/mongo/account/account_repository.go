package account

import (
	"context"
	"errors"
	"github.com/dhiemaz/AccountService/infrastructures/logger"
	"github.com/dhiemaz/AccountService/internal/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"reflect"
	"time"
)

type MongoAdapter struct {
	mgo *mongo.Collection
}

// NewAccountMongoAdapter account data with mongo adapter
func NewAccountMongoAdapter(mgo *mongo.Collection) *MongoAdapter {
	return &MongoAdapter{
		mgo: mgo,
	}
}

// InsertAccount to database
func (r *MongoAdapter) InsertAccount(data model.Account) (model.Account, error) {
	t := time.Now().Unix()
	data.RegisterDate = primitive.DateTime(t)

	data.Id = primitive.NewObjectID().Hex()
	if _, err := r.mgo.InsertOne(context.TODO(), data); err != nil {
		return data, err
	}

	return data, nil
}

// UpdateAccount in database
func (r *MongoAdapter) UpdateAccount(data model.Account) (model.Account, error) {

	filter := bson.D{{"username", data.Username}}
	updates := bson.D{}

	// get the type of struct data
	typeData := reflect.TypeOf(data)

	// get the values from the provided object
	values := reflect.ValueOf(data)

	// starting from index 1 to exclude the ID field
	for i := 1; i < typeData.NumField(); i++ {
		field := typeData.Field(i)   // get the field from the struct definition
		val := values.Field(i)       // get the value from the specified field position
		tag := field.Tag.Get("json") // from the field, get the json struct tag

		// we want to avoid zero values, as the omitted fields from newBook
		// corresponds to their zero values, and we only want provided fields
		if !isZeroType(val) {
			update := bson.E{Key: tag, Value: val.Interface()}
			updates = append(updates, update)
		}
	}

	updateFilter := bson.D{{"username", data.Username}}

	if _, err := r.mgo.UpdateOne(context.TODO(), filter, updateFilter); err != nil {
		return data, err
	}

	return data, nil
}

// FindOneAccountByUsername from database
func (r *MongoAdapter) FindOneAccountByUsername(username string) (model.Account, error) {
	filter := bson.D{{"username", username}}

	var data model.Account

	err := r.mgo.FindOne(context.TODO(), filter).Decode(&data)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			logger.WithFields(logger.Fields{"component": "account_repository", "action": "GetAccountByUsername"}).
				Errorf("no data found with username: %s\n", username)

			return data, err
		}

		logger.WithFields(logger.Fields{"component": "account_repository", "action": "GetAccountByUsername"}).
			Errorf("error retrieving account: %v", err)
	}

	return data, nil
}

// GetAllAccounts from database
func (r *MongoAdapter) GetAllAccounts() ([]model.Account, error) {
	var datas []model.Account

	cursor, err := r.mgo.Find(context.TODO(), bson.M{})
	if err != nil {
		logger.WithFields(logger.Fields{"component": "account_repository", "action": "GetAllAccounts"}).
			Errorf("error fetching all accounts: %v", err)

		return datas, err
	}

	if err := cursor.All(context.TODO(), &datas); err != nil {
		logger.WithFields(logger.Fields{"component": "account_repository", "action": "GetAllAccounts"}).
			Errorf("error decoding cursor data: %v", err)

		return datas, err
	}

	return datas, nil
}

func isZeroType(value reflect.Value) bool {
	zero := reflect.Zero(value.Type()).Interface()

	switch value.Kind() {
	case reflect.Slice, reflect.Array, reflect.Chan, reflect.Map:
		return value.Len() == 0
	default:
		return reflect.DeepEqual(zero, value.Interface())
	}
}
