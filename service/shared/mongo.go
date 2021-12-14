package shared

import (
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var Row struct {
	Id primitive.ObjectID `bson:"_id"`
}

func Set(set bson.M) bson.M {
	return bson.M{
		"$set": set,
	}
}

func SetOnInsert(set bson.M) bson.M {
	return bson.M{
		"$setOnInsert": set,
	}
}
