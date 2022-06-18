package main

import (
	"context"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	carpb "project/service/car/api"
	"project/service/car/dao"
	"project/service/shared"
	"project/service/shared/mgutil"
	"project/service/shared/service"
)

type CarService struct {
	carpb.UnimplementedCarServiceServer
	Dao dao.CarDao
	Log *zap.Logger
}

func (c *CarService) GetCars(ctx context.Context, in *carpb.GetCarsRequest) (*carpb.GetCarsResponse, error) {
	cars, err := c.Dao.GetCars()
	if err != nil {
		c.Log.Error("GetCars error", zap.Error(err))
		return nil, err
	}
	var carResult []*carpb.Car
	for _, c := range cars {
		carResult = append(carResult, &c.Car)
	}

	return &carpb.GetCarsResponse{Cars: carResult}, nil
}

func (c *CarService) UpdateCar(ctx context.Context, request *carpb.UpdateCarRequest) (*carpb.UpdateCarResponse, error) {
	err := c.Dao.UpdateCar(mgutil.CarId(request.CarId), shared.TripId(request.TripId), request.Status, *request.Location)
	if err != nil {
		c.Log.Error("GetCars error", zap.Error(err))
		return nil, err
	}
	return &carpb.UpdateCarResponse{}, nil
}

func main() {
	service.RegisterRpcService(service.RpcServiceConfig{
		Name: "汽车服务",
		Port: 9005,
		RegisterFunc: func(s *grpc.Server) {
			development, err := zap.NewDevelopment()
			if err != nil {
				panic(err)
			}
			carpb.RegisterCarServiceServer(s, &CarService{
				Dao: dao.CarDao{Database: mgutil.NewMongoDatabaseClient("cool")},
				Log: development,
			})
		},
	})
}
