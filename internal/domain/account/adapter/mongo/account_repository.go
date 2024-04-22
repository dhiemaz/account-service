package mongo

import (
	"context"
	"errors"
	"github.com/dhiemaz/AccountService/constant"
	"github.com/dhiemaz/AccountService/infrastructures/logger"
	"github.com/dhiemaz/AccountService/internal/domain/account/model"
	. "github.com/gobeam/mongo-go-pagination"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"golang.org/x/crypto/bcrypt"
	"log"
	"reflect"
	"strings"
	"time"
)

type MongoAdapter struct {
	mgo *mongo.Database
}

// NewAccountMongoAdapter account data with mongo adapter
func NewAccountMongoAdapter(mgo *mongo.Database) *MongoAdapter {
	return &MongoAdapter{
		mgo: mgo,
	}
}

// InsertAccount to database
func (r *MongoAdapter) InsertAccount(data model.Account) (model.Account, error) {

	account, _ := r.FindOneAccountByUsername(data.Username)
	if account.Username != "" {
		return data, errors.New("username is already taken, please choose another one")
	}

	// validate office (if not exist in master data then return error).
	// TODO : refactor this code
	collection := r.mgo.Collection(constant.OfficeCollection)

	for _, office := range data.Office {
		filter := bson.D{{"branch_id", office.BranchId}, {"branch_name", office.BranchName}}
		err := collection.FindOne(context.TODO(), filter).Err()
		if err != nil {
			if errors.Is(err, mongo.ErrNoDocuments) {
				return data, errors.New("invalid branch_id or branch_name")
			}
		}
	}

	// validate role (if not exist in master data then return error).
	// TODO : refactor this code
	collection = r.mgo.Collection(constant.RoleCollection)

	for _, role := range data.Role {
		filter := bson.D{{"role_id", role.RoleId}, {"role_name", role.RoleName}}
		err := collection.FindOne(context.TODO(), filter).Err()
		if err != nil {
			if errors.Is(err, mongo.ErrNoDocuments) {
				return data, errors.New("invalid role_id or role_name")
			}
		}
	}

	// validate access (if not exist in master data then return error).
	// TODO : refactor this code
	collection = r.mgo.Collection(constant.AccessCollection)

	for _, access := range data.Access {
		filter := bson.D{{"access_id", access.AccessId}, {"access_name", access.AccessName}}
		err := collection.FindOne(context.TODO(), filter).Err()
		if err != nil {
			if errors.Is(err, mongo.ErrNoDocuments) {
				return data, errors.New("invalid access_id or access_name")
			}
		}
	}

	t := time.Now().Unix()
	data.RegisterDate = primitive.DateTime(t)
	data.Password = hashAndSalt([]byte(data.Password))

	collection = r.mgo.Collection(constant.AccountCollection)

	keys := bson.D{}

	typeData := reflect.TypeOf(data)
	values := reflect.ValueOf(data)

	// starting from index 1 to exclude the ID field
	for i := 1; i < typeData.NumField(); i++ {
		field := typeData.Field(i)   // get the field from the struct definition
		val := values.Field(i)       // get the value from the specified field position
		tag := field.Tag.Get("json") // from the field, get the json struct tag

		// we want to avoid zero values, as the omitted fields from model.Account
		// corresponds to their zero values, and we only want provided fields
		if !isZeroType(val) {
			var key bson.E
			if strings.Contains(tag, "id") {
				key = bson.E{Key: tag, Value: 1}
			} else {
				key = bson.E{Key: tag, Value: "text"}
			}
			keys = append(keys, key)
		}
	}

	// Create index (if not already created)
	_, err := collection.Indexes().CreateOne(
		context.TODO(),
		mongo.IndexModel{
			Keys:    keys,                                       // Index key(s)
			Options: options.Index().SetName("f_account_index"), // Index options
		},
	)

	if err != nil {
		return data, err
	}

	data.Id = primitive.NewObjectID().Hex()
	if _, err := collection.InsertOne(context.TODO(), data); err != nil {
		return data, err
	}

	return data, nil
}

// UpdateAccount in database
func (r *MongoAdapter) UpdateAccount(data model.Account) (model.Account, error) {

	account, _ := r.FindOneAccountByUsername(data.Username)
	if account.Username == "" {
		return data, errors.New("username is not found")
	}

	// validate office (if not exist in master data then return error).
	// TODO : refactor this code
	collection := r.mgo.Collection(constant.OfficeCollection)

	for _, office := range data.Office {
		filter := bson.D{{"branch_id", office.BranchId}, {"branch_name", office.BranchName}}
		err := collection.FindOne(context.TODO(), filter).Err()
		if err != nil {
			if errors.Is(err, mongo.ErrNoDocuments) {
				return data, errors.New("invalid branch_id or branch_name")
			}
		}
	}

	// validate role (if not exist in master data then return error).
	// TODO : refactor this code
	collection = r.mgo.Collection(constant.RoleCollection)

	for _, role := range data.Role {
		filter := bson.D{{"role_id", role.RoleId}, {"role_name", role.RoleName}}
		err := collection.FindOne(context.TODO(), filter).Err()
		if err != nil {
			if errors.Is(err, mongo.ErrNoDocuments) {
				return data, errors.New("invalid role_id or role_name")
			}
		}
	}

	// validate access (if not exist in master data then return error).
	// TODO : refactor this code
	collection = r.mgo.Collection(constant.AccessCollection)

	for _, access := range data.Access {
		filter := bson.D{{"access_id", access.AccessId}, {"access_name", access.AccessName}}
		err := collection.FindOne(context.TODO(), filter).Err()
		if err != nil {
			if errors.Is(err, mongo.ErrNoDocuments) {
				return data, errors.New("invalid access_id or access_name")
			}
		}
	}

	collection = r.mgo.Collection(constant.AccountCollection)

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

		// we want to avoid zero values, as the omitted fields from model.Account
		// corresponds to their zero values, and we only want provided fields
		if !isZeroType(val) {
			update := bson.E{Key: tag, Value: val.Interface()}
			updates = append(updates, update)
		}
	}

	filter := bson.D{{"username", data.Username}}
	updateFilter := bson.D{{"$set", updates}}

	if _, err := collection.UpdateOne(context.TODO(), filter, updateFilter); err != nil {
		return data, err
	}

	return data, nil
}

// FindOneAccountByUsername from database
func (r *MongoAdapter) FindOneAccountByUsername(username string) (model.Account, error) {
	var data model.Account

	collection := r.mgo.Collection(constant.AccountCollection)

	filter := bson.D{{"username", username}}

	err := collection.FindOne(context.TODO(), filter).Decode(&data)
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
func (r *MongoAdapter) GetAllAccounts(searchFilter string, page, limit int) ([]model.Account, PaginationData, error) {
	collection := r.mgo.Collection(constant.AccountCollection)
	filter := bson.D{}

	if searchFilter != "" {
		filter = append(filter, bson.E{Key: "$text", Value: bson.M{"$search": searchFilter}})
	}

	convertedLimitInt := int64(limit)
	convertedPageInt := int64(page)

	var accounts []model.Account

	paginatedData, err := New(collection).Limit(convertedLimitInt).Page(convertedPageInt).Decode(&accounts).Filter(filter).Find()
	if err != nil {
		logger.WithFields(logger.Fields{"component": "account_repository", "action": "GetAllAccounts"}).
			Errorf("error fetching all accounts: %v", err)
	}

	return accounts, paginatedData.Pagination, nil
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

func hashAndSalt(pwd []byte) string {

	// Use GenerateFromPassword to hash & salt pwd.
	// MinCost is just an integer constant provided by the bcrypt
	// package along with DefaultCost & MaxCost.
	// The cost can be any value you want provided it isn't lower
	// than the MinCost (4)
	hash, err := bcrypt.GenerateFromPassword(pwd, bcrypt.MinCost)
	if err != nil {
		log.Println(err)
	}

	return string(hash)
}
