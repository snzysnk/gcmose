package text

import (
	"fmt"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"golang.org/x/net/context"
	"project/service/auth/dao"
	"testing"
)

func TestFindAndModify(t *testing.T) {
	background := context.Background()
	connect, err := mongo.Connect(background, options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		t.Errorf("cannot connect mongodb %v", err)
	}
	mg := dao.NewMongo(connect.Database("cool"), background, func() primitive.ObjectID {
		return primitive.NewObjectID()
	})
	//fmt.Println(mg)
	id, err := mg.ResolveAccountId("123")

	if err != nil {
		t.Errorf("error %v", err)
	}

	fmt.Println(id)
}
