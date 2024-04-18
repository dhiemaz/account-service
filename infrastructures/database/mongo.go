package database

import (
	"context"
	"errors"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"time"
)

// MongoConnect connect to mongodb
func MongoConnect(uri string) (conn *mongo.Client, err error) {
	var (
		opt                   = make([]*options.ClientOptions, 0)
		defaultConnectTimeout = 10 * time.Second
		defaultPingTimeout    = 2 * time.Second
	)

	opt = append(opt, options.Client().SetAppName("account-service").ApplyURI(uri))
	ctx, _ := context.WithTimeout(context.Background(), defaultConnectTimeout)
	conn, err = mongo.Connect(ctx, opt...)
	if err != nil {
		err = errors.New("failed to create mongodb client")
		return
	}

	ctx, _ = context.WithTimeout(context.Background(), defaultPingTimeout)

	if err = conn.Ping(ctx, readpref.Primary()); err != nil {
		err = errors.New("failed to establish connection to mongodb server")
	}

	return nil, err
}

// MongoMustConnect connection to MongoDB
func MongoMustConnect(uri string) *mongo.Client {
	conn, err := MongoConnect(uri)
	if err != nil {
		panic(err)
	}

	return conn
}
