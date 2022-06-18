package dao

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	carpb "project/service/car/api"
	"project/service/shared"
	"project/service/shared/mgutil"
)

type CarRecord struct {
	mgutil.ObjectIdField `bson:"inline"`
	mgutil.UpdateAtField `bson:"inline"`
	Car                  carpb.Car
}

type CarInterface interface {
}

type CarDao struct {
	Database *mongo.Database
}

func (c *CarDao) CreateCar() error {
	CarId := mgutil.NewObjectIdFunc()
	record := CarRecord{
		Car: carpb.Car{
			Status:   carpb.Car_Status_LOCKED,
			TripId:   "",
			Location: nil,
			CarId:    CarId.Hex(),
		},
	}
	record.ID = CarId
	record.UpdateAt = mgutil.UpdateAtFunc()
	_, err := c.Database.Collection("car").InsertOne(context.Background(), record)
	if err != nil {
		return err
	}
	return nil
}

func (c *CarDao) GetCars() ([]CarRecord, error) {
	find, err := c.Database.Collection("car").Find(context.Background(), bson.M{})
	if err != nil {
		return nil, err
	}
	var result CarRecord
	var record []CarRecord
	for find.Next(context.Background()) {
		find.Decode(&result)
		record = append(record, result)
	}
	return record, nil
}

func (c *CarDao) UpdateCar(carId mgutil.CarId, tripId shared.TripId, status carpb.Car_Status, location carpb.Location) error {
	id, err := carId.TransformToMongoId()
	if err != nil {
		return err
	}
	singleResult := c.Database.Collection("car").FindOne(context.Background(), bson.M{
		mgutil.IdField: id,
	})
	if err = singleResult.Err(); err != nil {
		return err
	}
	_, err = c.Database.Collection("car").UpdateOne(context.Background(), bson.M{mgutil.IdField: id}, mgutil.Set(bson.M{
		"car.status":   status,
		"car.location": location,
		"car.tripid":   tripId.String(),
	}))
	if err != nil {
		return err
	}
	return nil
}
