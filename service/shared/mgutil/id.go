package mgutil

import "go.mongodb.org/mongo-driver/bson/primitive"

const IdField = "_id"
const UpdateField = "updateat"
const AccountField = "accountid"

type ObjectIdField struct {
	ID primitive.ObjectID `bson:"_id"`
}

type UpdateAtField struct {
	UpdateAt int64
}

type CarId string

func (c CarId) TransformToMongoId() (primitive.ObjectID, error) {
	return primitive.ObjectIDFromHex(c.String())
}

func (c CarId) String() string {
	return string(c)
}

var NewObjectIdFunc = primitive.NewObjectID
