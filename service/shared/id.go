package shared

import "go.mongodb.org/mongo-driver/bson/primitive"

type AccountId string
type TripId string

func (a AccountId) String() string {
	return string(a)
}

func (t TripId) String() string {
	return string(t)
}

func (t *TripId) ObjectId() (primitive.ObjectID, error) {
	return primitive.ObjectIDFromHex(t.String())
}
