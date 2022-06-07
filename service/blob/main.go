package main

import (
	"context"
	"google.golang.org/grpc"
	"project/service/blob/api"
	myos "project/service/blob/oss"
)

const AccessKeyId = "LTAI5t8msVaKSFucLeBwsVC6"
const AccessKeySecret = "4ellrwDdU7vREBXvlErPE5YN5FZwKn"

type Service struct {
	blobpb.UnimplementedBlobServiceServer
	Client myos.OssInterface
}

func (s *Service) GetUploadUrl(ctx context.Context, in *blobpb.GetUploadUrlRequest, opts ...grpc.CallOption) (*blobpb.GetUploadUrlResponse, error) {
	panic("implement me")
}

func (s *Service) GetUploadData(ctx context.Context, in *blobpb.GetUploadDataRequest, opts ...grpc.CallOption) (*blobpb.GetUploadDataResponse, error) {
	panic("implement me")
}

func (s *Service) GetFileUrl(ctx context.Context, in *blobpb.GetFileUrlRequest, opts ...grpc.CallOption) (*blobpb.GetFileUrlResponse, error) {
	panic("implement me")
}

func main() {
	service := Service{
		Client: &myos.OssService{
			AccessKeyId:     AccessKeyId,
			AccessKeySecret: AccessKeySecret,
		},
	}

	service.GetFileUrl(context.Background(), &blobpb.GetFileUrlRequest{Path: "/cool/b.png"})
}
