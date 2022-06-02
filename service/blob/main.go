package main

import (
	"context"
	"google.golang.org/grpc"
	profilepb "project/service/blob/api"
	"project/service/shared/service"
)

type Service struct {
	profilepb.UnimplementedProfileServiceServer
}

func (s *Service) GetProfile(ctx context.Context, in *profilepb.Request) (*profilepb.Response, error) {
	return &profilepb.Response{Code: "200"}, nil
}

func main() {
	service.RegisterRpcService(service.RpcServiceConfig{
		Name: "图片服务",
		Port: 9006,
		RegisterFunc: func(s *grpc.Server) {
			profilepb.RegisterProfileServiceServer(s, &Service{})
		},
	})
}
