package main

import (
	"google.golang.org/grpc"
	myos "project/service/blob/oss"
	trippb "project/service/rental/api"
	"project/service/rental/trip"
	"project/service/shared/mgutil"
	"project/service/shared/service"
)

const AccessKeyId = "LTAI5t8msVaKSFucLeBwsVC6"
const AccessKeySecret = "4ellrwDdU7vREBXvlErPE5YN5FZwKn"

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

			trippb.RegisterProfileServiceServer(s, &trip.ProfileService{
				Dao: trip.ProfileServiceDao{
					Database: mgutil.NewMongoDatabaseClient("cool"),
				},
				Oss: myos.OssService{
					AccessKeyId:     AccessKeyId,
					AccessKeySecret: AccessKeySecret,
				},
			})
		},
		ValidateToken: true,
	})
}
