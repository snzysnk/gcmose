package text

import (
	"context"
	"fmt"
	"github.com/google/go-cmp/cmp"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"google.golang.org/protobuf/testing/protocmp"
	trippb "project/service/rental/api"
	"project/service/rental/trip"
	"project/service/shared"
	"testing"
)

func TestCreateTrip(t *testing.T) {
	service := trip.TripMongoService{}
	result, err := service.CreateTrip(shared.AccountId("abc123"), 2021, trippb.LocationStatus{
		Location: &trippb.Location{
			Longitude: 110.11,
			Latitude:  120.12,
		},
		Name: "虹桥火车站",
		Fee:  0,
		Km:   0,
	})

	if err != nil {
		t.Error(err)
	}

	record, err := service.FindTrip(shared.AccountId("abc123"))

	if err != nil {
		t.Error(err)
	}

	diff := cmp.Diff(result, record, protocmp.Transform())

	if diff != "" {
		t.Error("存入和取出数据不相同")
	}
}

func createUniqueIndex() error {
	connect, err := mongo.Connect(context.Background(), options.Client().ApplyURI("mongodb://localhost:27017"))

	collection := connect.Database("cool").Collection("trip")

	if err != nil {
		return err
	}

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
		return err
	}

	if res == "" {
		return fmt.Errorf("create unique result is empty")
	}

	return nil
}

func TestCreateTrips(t *testing.T) {
	i := []struct {
		Name      string
		Status    trippb.TripStatus
		AccountId shared.AccountId
		CartId    int64
		wantError bool
	}{
		{
			Name:      "test 001",
			Status:    trippb.TripStatus_TRIP_END,
			CartId:    100,
			AccountId: shared.AccountId("1001"),
			wantError: false,
		},
		{
			Name:      "test 002",
			Status:    trippb.TripStatus_TRIP_ING,
			CartId:    100,
			AccountId: shared.AccountId("1001"),
			wantError: false,
		},
		{
			Name:      "test 003",
			Status:    trippb.TripStatus_TRIP_ING,
			CartId:    100,
			AccountId: shared.AccountId("1001"),
			wantError: true,
		},
		{
			Name:      "test 004",
			Status:    trippb.TripStatus_TRIP_ING,
			CartId:    101,
			AccountId: shared.AccountId("1002"),
			wantError: false,
		},
	}

	service := trip.TripMongoService{}
	err := createUniqueIndex()
	if err != nil {
		t.Error("can't create unique index")
	}

	for _, v := range i {
		trip.StartStatus = v.Status
		result, err := service.CreateTrip(v.AccountId, v.CartId, trippb.LocationStatus{
			Location: &trippb.Location{
				Longitude: 110.11,
				Latitude:  120.12,
			},
			Name: "虹桥火车站",
			Fee:  0,
			Km:   0,
		})

		if v.wantError && err != nil {
			continue
		}

		if err != nil {
			t.Errorf("test 出现错误 %s", v.Name)
		}

		record, err := service.FindTrip(v.AccountId)

		if err != nil {
			t.Error(err)
		}

		diff := cmp.Diff(result, record, protocmp.Transform())

		if diff != "" {
			t.Error("存入和取出数据不相同")
		}
	}

}

func TestUpdateTrip(t *testing.T) {
	tid, err := primitive.ObjectIDFromHex("67875ce54ce71f1073d79749")
	if err != nil {
		panic(err)
	}
	accountId := shared.AccountId("1997")

	testData := []struct {
		Name      string
		wantError bool
		updateAt  int64
		now       int64
		tripId    primitive.ObjectID
	}{
		{
			Name:      "test01",
			wantError: false,
			updateAt:  1000,
			now:       2000,
			tripId:    tid,
		},
		{
			Name:      "test02",
			wantError: true,
			updateAt:  1000,
			now:       2000,
			tripId:    tid,
		},
		{
			Name:      "test03",
			wantError: false,
			updateAt:  2000,
			now:       3000,
			tripId:    tid,
		},
	}

	service := trip.TripMongoService{}
	trip.NewObjectIdFunc = func() primitive.ObjectID {
		return tid
	}
	trip.UpdateAtFunc = func() int64 {
		return 1000
	}

	tr, _ := service.CreateTrip(accountId, 2022, trippb.LocationStatus{
		Location: &trippb.Location{
			Longitude: 110.11,
			Latitude:  120.12,
		},
		Name: "虹桥火车站",
		Fee:  0,
		Km:   0,
	})

	for _, v := range testData {
		trip.UpdateAtFunc = func() int64 {
			return v.now
		}
		err := service.UpdateTrip(tid, v.updateAt, tr.Trip)

		if v.wantError {
			if err == nil {
				t.Errorf("测试%s应该出现异常才对", v.Name)
			}
			continue
		}

		if err != nil {
			t.Errorf("测试%s出现异常%v", v.Name, err)
		}
	}

}
