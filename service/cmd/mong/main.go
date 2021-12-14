package main

import (
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"golang.org/x/net/context"
)

var row struct {
	Name string `bson:"name"`
	Age  int32  `bson:"age"`
	//Money bson.M `bson:"des.money"`
}

func main() {
	s := "mongodb://localhost:27017"
	background := context.Background()
	connect, err := mongo.Connect(background, options.Client().ApplyURI(s))
	if err != nil {
		panic(err)
	}
	collection := connect.Database("app").Collection("student")
	//insertMany(background, collection)
	//findOne(background, collection)
	findMany(background, collection)
}

func findMany(ctx context.Context, collection *mongo.Collection) {
	find, err := collection.Find(ctx, bson.M{})
	if err != nil {
		panic(err)
	}
	for find.Next(ctx) {
		find.Decode(&row)
		fmt.Printf("%+v \n", row)
	}
}

func findOne(ctx context.Context, collection *mongo.Collection) {
	one := collection.FindOne(ctx, bson.M{})

	var row struct {
		Name string `bson:"name"`
	}

	err := one.Decode(&row)
	if err != nil {
		panic(err)
	}

	fmt.Printf("%+v", row.Name)
}

func insertMany(ctx context.Context, collection *mongo.Collection) {
	many, err := collection.InsertMany(ctx, []interface{}{bson.M{"name": "age"}, bson.M{"name": "天王盖地虎"}})
	if err != nil {
		panic(err)
	}
	fmt.Printf("%+v", many)
}

func insertOne(ctx context.Context, collection *mongo.Collection) {
	one, err := collection.InsertOne(ctx, bson.M{"name": "浪子", "age": 27})
	if err != nil {
		panic(err)
	}
	fmt.Printf("%+v", one)
}
