package db

import (
	"context"
	"fmt"
	"go-identity-service/config"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

var (
	MongoClient *mongo.Client
)

func Connect() *mongo.Client {
	uri := fmt.Sprintf("%s%s:%s%s", config.Env.DBConnectionPrefix, config.Env.DBUser, config.Env.DBPassword, config.Env.DBConnectionSuffix)

	MongoClient, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(uri))
	if err != nil {
		panic(err)
	}

	if err := MongoClient.Ping(context.TODO(), readpref.Primary()); err != nil {
		panic(err)
	}

	return MongoClient
}
