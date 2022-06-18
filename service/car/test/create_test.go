package test

import (
	"fmt"
	carpb "project/service/car/api"
	"project/service/car/dao"
	"project/service/shared/mgutil"
	"testing"
)

func TestCreate(t *testing.T) {
	carDao := dao.CarDao{Database: mgutil.NewMongoDatabaseClient("cool")}
	err := carDao.CreateCar()
	if err != nil {
		t.Error(err)
	}
}

func TestGets(t *testing.T) {
	carDao := dao.CarDao{Database: mgutil.NewMongoDatabaseClient("cool")}
	cars, err := carDao.GetCars()
	if err != nil {
		t.Error(err)
	}
	fmt.Println(cars)
}

func TestUpdate(t *testing.T) {
	carDao := dao.CarDao{Database: mgutil.NewMongoDatabaseClient("cool")}
	err := carDao.UpdateCar(mgutil.CarId("62aca081dc8aa63390a07ea4"), carpb.Car_Status_UNLOCKING, carpb.Location{
		Longitude: 120.52,
		Latitude:  31.53,
	})
	if err != nil {
		t.Error(err)
	}
}
