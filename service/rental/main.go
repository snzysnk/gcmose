package main

import (
	"google.golang.org/grpc"
	trippb "project/service/rental/api"
	"project/service/rental/trip"
	"project/service/shared/service"
)

func main() {
	service.RegisterRpcService(service.RpcServiceConfig{
		Name: "行程服务",
		Port: 9003,
		RegisterFunc: func(s *grpc.Server) {
			trippb.RegisterTripServiceServer(s, &trip.TripService{
				MgService: &trip.TripMongoService{},
				Identity:  &trip.IdentityService{},
				Lock:      &trip.LockService{},
				Help:      &trip.HelpService{},
			})
		},
		ValidateToken: true,
	})
}
