package mongo

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"golang-web-project-framework/src/config"
	"time"
)

type Mgo struct {
	Database   string
	Collection string
}

func NewMgo(database, collection string) *Mgo {
	return &Mgo{
		database,
		collection,
	}
}

func (m *Mgo) InsertOne(document interface{}, opts ...*options.InsertOneOptions) (*mongo.InsertOneResult, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(config.Config.MongoDb.OperateCtxTimeout)*time.Second)
	defer cancel()
	return MongoDb.MongoClient.Database(m.Database).Collection(m.Collection).InsertOne(ctx, document, opts...)
}

func (m *Mgo) DeleteOne(filter interface{}, opts ...*options.DeleteOptions) (*mongo.DeleteResult, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(config.Config.MongoDb.OperateCtxTimeout)*time.Second)
	defer cancel()
	return MongoDb.MongoClient.Database(m.Database).Collection(m.Collection).DeleteOne(ctx, filter, opts...)
}

func (m *Mgo) UpdateOne(filter interface{}, update interface{}, opts ...*options.UpdateOptions) (*mongo.UpdateResult, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(config.Config.MongoDb.OperateCtxTimeout)*time.Second)
	defer cancel()
	return MongoDb.MongoClient.Database(m.Database).Collection(m.Collection).UpdateOne(ctx, filter, update, opts...)
}

func (m *Mgo) Find(filter interface{}, opts ...*options.FindOptions) (*mongo.Cursor, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(config.Config.MongoDb.OperateCtxTimeout)*time.Second)
	defer cancel()
	return MongoDb.MongoClient.Database(m.Database).Collection(m.Collection).Find(ctx, filter, opts...)
}

func (m *Mgo) CountDocuments(filter interface{}, opts ...*options.CountOptions) (int64, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(config.Config.MongoDb.OperateCtxTimeout)*time.Second)
	defer cancel()
	return MongoDb.MongoClient.Database(m.Database).Collection(m.Collection).CountDocuments(ctx, filter, opts...)
}

func (m *Mgo) FindOne(filter interface{}, opts ...*options.FindOneOptions) *mongo.SingleResult {
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(config.Config.MongoDb.OperateCtxTimeout)*time.Second)
	defer cancel()
	return MongoDb.MongoClient.Database(m.Database).Collection(m.Collection).FindOne(ctx, filter, opts...)
}
