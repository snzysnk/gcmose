package test

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"project/service/auth/dao"
	"testing"
)

var ApplyURI string

func TestDocker(t *testing.T) {
	background := context.Background()
	connect, err := mongo.Connect(background, options.Client().ApplyURI(ApplyURI))
	if err != nil {
		t.Errorf("cannot connect mongodb %v", err)
	}
	mg := dao.NewMongo(connect.Database("cool"), background, func() primitive.ObjectID {
		hex, err2 := primitive.ObjectIDFromHex("619b49a1d7d544609ca2c25f")
		if err2 != nil {
			panic(err2)
		}
		return hex
	})

	err = mg.InsertMany([]interface{}{bson.M{"_id": "619b49a1d7d544609ca2c25d", "openId": "open1"}, bson.M{"_id": "619b49a1d7d544609ca2c25e", "openId": "open2"}})

	if err != nil {
		t.Errorf("%+v", err)
	}

	data := []struct {
		openId string
		want   string
	}{
		{
			openId: "open1",
			want:   "619b49a1d7d544609ca2c25d",
		},
		{
			openId: "open2",
			want:   "619b49a1d7d544609ca2c25e",
		},
		{
			openId: "open3",
			want:   "619b49a1d7d544609ca2c25f",
		},
	}

	for _, v := range data {
		id, err := mg.ResolveAccountId(v.openId)
		if err != nil {
			t.Error(err)
		}
		if id != v.want {
			t.Errorf("%s not equal %s", id, v.want)
		}
	}

}

//func TestMain(m *testing.M) {
//	//help.RunMongoDbInDocker(&ApplyURI, m)
//	//os.Exit(m.Run())
//}
