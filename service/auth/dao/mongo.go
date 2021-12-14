package dao

import (
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"golang.org/x/net/context"
	"project/service/shared"
)

const OPENID = "openId"

type CreateId func() primitive.ObjectID

type Mg struct {
	col        *mongo.Collection
	background context.Context
	CreateId   CreateId
}

func NewMongo(database *mongo.Database, ctx context.Context, createId CreateId) Mg {
	return Mg{col: database.Collection("user"), background: ctx, CreateId: createId}
}

func (s *Mg) ResolveAccountId(openId string) (string, error) {
	update := s.col.FindOneAndUpdate(s.background, bson.M{OPENID: openId}, shared.SetOnInsert(bson.M{OPENID: openId, "_id": s.CreateId()}), options.FindOneAndUpdate().SetUpsert(true).SetReturnDocument(options.After))

	if err := update.Err(); err != nil {
		return "", fmt.Errorf("cannot findOneAndUpdate %v", err)
	}

	update.Decode(&shared.Row)

	return shared.Row.Id.Hex(), nil
}

func (s *Mg) InsertMany(sql []interface{}) error {
	_, err := s.col.InsertMany(s.background, sql)
	if err != nil {
		return err
	}

	return nil
}
