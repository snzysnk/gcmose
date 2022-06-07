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
