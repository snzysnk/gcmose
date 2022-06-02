package text

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"testing"
)

func TestCreateUniqueIndex(t *testing.T) {
	connect, err := mongo.Connect(context.Background(), options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		t.Error(err)
	}
	collection := connect.Database("cool").Collection("trip")
	/**
	  等价于直接在mongo中创建索引语句
	  该语句保证 trip.accountid + trip.status = 1 的数据唯一
	  db.trip.createIndex({
	    "trip.accountid": 1,
	    "trip.status": 1,
	}, {
	    unique: true,
	    partialFilterExpression: {
	        "trip.status": 1,
	    }
	})
	*/
	res, err := collection.Indexes().CreateOne(context.Background(), mongo.IndexModel{
		Keys: bson.D{
			{Key: "trip.accountid", Value: 1},
			{Key: "trip.status", Value: 1},
		},
		Options: options.Index().SetUnique(true).SetPartialFilterExpression(bson.M{
			"trip.status": 1,
		}),
	})

	if err != nil {
		t.Error(err)
	}

	fmt.Println(res)
}
