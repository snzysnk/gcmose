package main

import (
	"context"
	"google.golang.org/grpc"
	blobpb "project/service/blob/api"
	myos "project/service/blob/oss"
	"project/service/shared/service"
)

const AccessKeyId = "LTAI5t8msVaKSFucLeBwsVC6"
const AccessKeySecret = "4ellrwDdU7vREBXvlErPE5YN5FZwKn"

type BlobService struct {
	blobpb.UnimplementedBlobServiceServer
	Oss myos.OssInterface
}

func (s *BlobService) CreateUrl(ctx context.Context, request *blobpb.CreateUrlRequest) (*blobpb.CreateUrlResponse, error) {
	url, err := s.Oss.CreateSingUrl(request.Path, request.Operation)
	if err != nil {
		return nil, err
	}
	return &blobpb.CreateUrlResponse{Url: url}, nil
}

func main() {
	service.RegisterRpcService(service.RpcServiceConfig{
		Name: "图片云服务",
		Port: 9004,
		RegisterFunc: func(s *grpc.Server) {
			blobpb.RegisterBlobServiceServer(s, &BlobService{
				Oss: &myos.OssService{
					AccessKeyId:     AccessKeyId,
					AccessKeySecret: AccessKeySecret,
				},
			})
		},
		ValidateToken: false,
	})
}
