package database

import (
	"context"
	"errors"
	"fmt"
	"github.com/dhiemaz/AccountService/config"
	"github.com/dhiemaz/AccountService/infrastructures/logger"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"os"
	"time"
)

var mgoDB *mongo.Database

// MongoMustConnect connection to MongoDB
func MongoMustConnect(cfg *config.Config) {
	// uri := "mongodb://admin:password@localhost:27017/db"
	uri := fmt.Sprintf("mongodb://%s:%s@%s:%s", cfg.MongoUsername, cfg.MongoPassword, cfg.MongoHost, cfg.MongoPort)

	logger.WithFields(logger.Fields{"component": "database", "action": "MongoMustConnect"}).
		Infof("connecting to mongodb, uri: %s", uri)

	conn, err := connect(uri, cfg.MongoDatabase)
	if err != nil {
		logger.WithFields(logger.Fields{"component": "database", "action": "MongoMustConnect"}).
			Errorf("failed connect to mongodb, error: %v", err)

		os.Exit(0)
	}

	mgoDB = conn
}

func GetMongoConnection() *mongo.Database {
	return mgoDB
}

func connect(uri, database string) (*mongo.Database, error) {
	var (
		opt                   = make([]*options.ClientOptions, 0)
		defaultConnectTimeout = 10 * time.Second
		defaultPingTimeout    = 2 * time.Second
	)

	opt = append(opt, options.Client().SetAppName("account-service").ApplyURI(uri))
	ctx, _ := context.WithTimeout(context.Background(), defaultConnectTimeout)
	mgo, err := mongo.Connect(ctx, opt...)
	if err != nil {
		err = errors.New("failed to create mongodb client")
		return nil, err
	}

	ctx, _ = context.WithTimeout(context.Background(), defaultPingTimeout)

	if err = mgo.Ping(ctx, readpref.Primary()); err != nil {
		err = errors.New("failed to establish connection to mongodb server")
		return nil, err
	}

	return mgo.Database(database), err
}
