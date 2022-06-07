package trip

import (
	"context"
	"fmt"
	"github.com/golang/protobuf/proto"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"hash/fnv"
	trippb "project/service/rental/api"
	"project/service/shared"
	"project/service/shared/mgutil"
)

var Address = []string{
	"浦东新区",
	"长宁区",
	"金山区",
	"宝山区",
	"普陀区",
	"闵行区",
	"嘉定区",
	"奉贤区",
}

type Trip interface {
	CreateTrip(accountId shared.AccountId, cartId int64, start trippb.LocationStatus) (trip *Record, err error)
	FindTrip(account shared.AccountId) (*Record, error)
	UpdateTrip(tripId primitive.ObjectID, updateAt int64, trip *trippb.Trip) error
}

type Help interface {
	Resolve(location *trippb.Location) string
	Calculate(before *trippb.Location, current *trippb.Location, beforeSecond int64, currentSecond int64) (fee float64, km float64, err error)
}

var UnlockError error
var IdentityError error

type Lock interface {
	Unlock(c context.Context, cartId int64) error
}

type Identity interface {
	Verify(c context.Context, id shared.AccountId) error
}

type HelpService struct {
}

func (h *HelpService) Resolve(location *trippb.Location) string {
	marshal, err := proto.Marshal(location)
	if err != nil {
		return "未知地址"
	}
	hs := fnv.New32()
	hs.Write(marshal)
	key := hs.Sum32() % 8
	return Address[key]
}

func (h *HelpService) Calculate(before *trippb.Location, current *trippb.Location, beforeSecond int64, currentSecond int64) (fee float64, km float64, err error) {
	return float64(currentSecond) * 0.02, float64(currentSecond) * 0.01, nil
}

type LockService struct {
}

type IdentityService struct {
}

func (l *LockService) Unlock(c context.Context, cartId int64) error {
	return UnlockError
}

func (i *IdentityService) Verify(c context.Context, id shared.AccountId) error {
	return IdentityError
}

type TripMongoService struct {
}

func (t *TripMongoService) CreateTrip(accountId shared.AccountId, cartId int64, start trippb.LocationStatus) (trip *Record, err error) {
	record := &Record{
		Trip: &trippb.Trip{
			Start:     &start,
			Current:   &start,
			Status:    StartStatus,
			CartId:    cartId,
			AccountId: accountId.String(),
		},
	}

	record.UpdateAt = UpdateAtFunc()
	record.ID = NewObjectIdFunc()
	mg := mgutil.NewMongoDatabaseClient("cool")
	res, err := mg.Collection("trip").InsertOne(context.Background(), record)
	if err != nil || res == nil {
		return nil, status.Error(codes.Unavailable, "can't create trip")
	}

	return record, nil
}

func (t *TripMongoService) FindTrip(accountId shared.AccountId) (*Record, error) {
	mg := mgutil.NewMongoDatabaseClient("cool")
	trip := &Record{}

	result := mg.Collection("trip").FindOne(context.Background(), bson.M{"trip.accountid": accountId.String(), "trip.status": trippb.TripStatus_TRIP_ING})

	if err := result.Err(); err != nil {
		return nil, err
	}

	err := result.Decode(trip)

	if err != nil {
		return nil, err
	}

	return trip, nil
}

func (t *TripMongoService) UpdateTrip(tripId primitive.ObjectID, updateAt int64, trip *trippb.Trip) error {
	mg := mgutil.NewMongoDatabaseClient("cool")

	result, err := mg.Collection("trip").UpdateOne(context.Background(), bson.M{
		mgutil.IdField:     tripId,
		"trip.status":      trippb.TripStatus_TRIP_ING,
		mgutil.UpdateField: updateAt,
	}, mgutil.Set(bson.M{"trip": trip, "updateat": UpdateAtFunc()}))

	if err != nil {
		fmt.Println(err)
		return status.Error(codes.Unknown, "update result fail")
	}

	if result.MatchedCount == 0 {
		return status.Error(codes.Unknown, "no document to update")
	}

	return nil
}

type ProfileRecord struct {
	mgutil.ObjectIdField `bson:"inline"`
	mgutil.UpdateAtField `bson:"inline"`
	Profile              *trippb.Profile
}

type ProfileServiceDao struct {
	Database *mongo.Database
}

func (s *ProfileServiceDao) GetProfile(accountId shared.AccountId) (*ProfileRecord, error) {
	where := bson.M{
		"profile.accountid": accountId.String(),
	}

	res := s.Database.Collection("profile").FindOne(context.Background(), where)

	if res.Err() == mongo.ErrNoDocuments {
		return s.CreateProfile(accountId)
	}

	if err := res.Err(); err != nil {
		return nil, err
	}

	var result ProfileRecord

	res.Decode(&result)

	return &result, nil
}

func (s *ProfileServiceDao) CreateProfile(accountId shared.AccountId) (*ProfileRecord, error) {
	profileRecord := &ProfileRecord{
		Profile: &trippb.Profile{
			Name:      "",
			Sex:       0,
			Birth:     0,
			Path:      "",
			AccountId: accountId.String(),
			Status:    trippb.ValidateStatus_WAIT,
		},
	}
	profileRecord.ID = NewObjectIdFunc()
	profileRecord.UpdateAt = UpdateAtFunc()
	_, err := s.Database.Collection("profile").InsertOne(context.Background(), profileRecord)
	if err != nil {
		return nil, err
	}
	return profileRecord, nil
}

func (s *ProfileServiceDao) UpdateProfile(id primitive.ObjectID, profile *trippb.Profile) error {
	result, err := s.Database.Collection("profile").UpdateOne(context.Background(), bson.M{
		mgutil.IdField: id,
	}, mgutil.Set(bson.M{"profile": profile}))
	if err != nil {
		return err
	}
	if result.MatchedCount == 0 {
		return fmt.Errorf("数据不存在")
	}
	return nil
}
