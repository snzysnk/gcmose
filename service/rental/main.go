package main

import (
	"fmt"
	"google.golang.org/grpc"
	blobpb "project/service/blob/api"
	trippb "project/service/rental/api"
	"project/service/rental/trip"
	"project/service/shared/mgutil"
	"project/service/shared/service"
)

const AccessKeyId = "xxxxx"
const AccessKeySecret = "xxxxx"

func main() {
	dial, err := grpc.Dial("localhost:9004", grpc.WithInsecure())
	if err != nil {
		fmt.Println("can't to connect oss service")
	}
	client := blobpb.NewBlobServiceClient(dial)
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

			trippb.RegisterProfileServiceServer(s, &trip.ProfileService{
				Dao: trip.ProfileServiceDao{
					Database: mgutil.NewMongoDatabaseClient("cool"),
				},
				Oss: client,
			})
		},
		ValidateToken: true,
	})
}
