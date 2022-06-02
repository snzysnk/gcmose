package mgutil

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func Set(set bson.M) bson.M {
	return bson.M{
		"$set": set,
	}
}

func NewMongoDatabaseClient(database string) *mongo.Database {
	connect, err := mongo.Connect(context.Background(), options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		panic("can't connect to mongo")
	}

	return connect.Database(database)
}
