package main

import (
	"go.uber.org/zap"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	trippb "project/service/rental/api"
	srpc "project/service/shared/grpc"
	"strconv"
)

type TripService struct {
	trippb.UnimplementedTripServiceServer
}

func (s *TripService) GetTrip(c context.Context, request *trippb.GetTripRequest) (*trippb.GetTripResponse, error) {
	development, _ := zap.NewDevelopment()
	development.Info("get request id", zap.String("id", strconv.Itoa(int(request.Id))))
	return &trippb.GetTripResponse{
		StatusCode: strconv.Itoa(int(request.Id)),
	}, nil
}

func main() {
	go createGrpcService()
	createGrpcGateWay()
}

func createGrpcService() {
	manger := srpc.ServiceManger{
		NewWork: "tcp",
		Port:    "9003",
		CallBack: func(server grpc.ServiceRegistrar) {
			trippb.RegisterTripServiceServer(server, &TripService{})
		},
	}

	manger.Start()
}

func createGrpcGateWay() {
	manager := srpc.GateWayManager{Port: "9002", Mux: srpc.CreateBaseMux()}
	manager.RegisterPoint(trippb.RegisterTripServiceHandlerFromEndpoint, "9003")
	manager.Start()
}
