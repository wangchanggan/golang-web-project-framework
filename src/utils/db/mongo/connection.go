package mongo

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"golang-web-project-framework/src/config"
	"time"
)

type MongoDatabase struct {
	MongoClient *mongo.Client
}

var MongoDb *MongoDatabase

func init() {
	MongoDb = &MongoDatabase{
		MongoClient: NewClient(),
	}
}

func NewClient() *mongo.Client {
	ctx, cancel := context.WithTimeout(context.Background(),  time.Duration(config.Config.MongoDb.ConnectCtxTimeout)*time.Second)
	defer cancel()

	uri := "mongodb://" + config.Config.MongoDb.Username + ":" + config.Config.MongoDb.Password + "@" + config.Config.MongoDb.Host + config.Config.MongoDb.Port
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(uri).SetMaxPoolSize(config.Config.MongoDb.MaxPoolSize))
	if err != nil {
		panic(err)
	}

	if err := client.Ping(ctx, readpref.Primary()); err != nil {
		panic(err)
	}

	return client
}
