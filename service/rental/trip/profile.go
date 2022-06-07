package trip

import (
	"context"
	"fmt"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"io/ioutil"
	"net/http"
	blobpb "project/service/blob/api"
	trippb "project/service/rental/api"
	"project/service/shared/service"
	"time"
)

type ProfileService struct {
	trippb.UnimplementedProfileServiceServer
	Dao ProfileServiceDao
	Oss blobpb.BlobServiceClient
}

func (p *ProfileService) GetUploadUrl(ctx context.Context, request *trippb.GetUploadUrlRequest) (*trippb.GetUploadUrlResponse, error) {
	accountId, err := service.GetContextAccountId(ctx)
	if err != nil {
		return nil, status.Error(codes.Aborted, "解析用户身份出错")
	}
	profile, err := p.Dao.GetProfile(accountId)
	if err != nil {
		return nil, status.Error(codes.Aborted, "获取用户身份验证信息出错")
	}
	if profile.Profile.Status != trippb.ValidateStatus_WAIT && profile.Profile.Status != trippb.ValidateStatus_FAIL {
		return nil, status.Error(codes.Aborted, "当前状态不能上传图片")
	}
	path := fmt.Sprintf("cool/%s.png", profile.ID.Hex())
	response, err := p.Oss.CreateUrl(ctx, &blobpb.CreateUrlRequest{
		Path:      path,
		Operation: blobpb.Operation_Status_UPLOAD,
	})
	if err != nil {
		return nil, status.Error(codes.Aborted, "不能创建临时上传连接")
	}
	return &trippb.GetUploadUrlResponse{Url: response.Url}, nil
}

func (p *ProfileService) ProfileCheck(ctx context.Context, request *trippb.ProfileCheckRequest) (*trippb.ProfileCheckResponse, error) {
	go func() {
		time.Sleep(time.Second * 3)
		accountId, err := service.GetContextAccountId(ctx)
		if err != nil {
			fmt.Println(err)
		}
		profile, err := p.Dao.GetProfile(accountId)
		if err != nil {
			fmt.Println(err)

		}
		path := fmt.Sprintf("cool/%s.png", profile.ID.Hex())
		grpcResponse, err := p.Oss.CreateUrl(context.Background(), &blobpb.CreateUrlRequest{
			Path:      path,
			Operation: blobpb.Operation_Status_DOWNLOAD,
		})

		if err != nil {
			fmt.Println("can't create singUrl")
		}

		response, err := http.Get(grpcResponse.Url)
		if err != nil {
			fmt.Println(err)
		}

		defer response.Body.Close()

		_, err = ioutil.ReadAll(response.Body)

		if err != nil {
			fmt.Println(err)
		}

		err = p.Dao.UpdateProfile(profile.ID, &trippb.Profile{
			Name:      request.Name,
			Sex:       request.Sex,
			Birth:     request.Birth,
			Path:      path,
			AccountId: accountId.String(),
			Status:    trippb.ValidateStatus_OK,
		})

		if err != nil {
			fmt.Sprintf("updatet profile fail, error %+v", err)
		}

		fmt.Println("执行结束")
	}()

	return &trippb.ProfileCheckResponse{Status: "检查进行中"}, nil
}

func (p *ProfileService) GetProfile(ctx context.Context, request *trippb.GetProfileRequest) (*trippb.GetProfileResponse, error) {
	accountId, err := service.GetContextAccountId(ctx)
	if err != nil {
		return nil, status.Error(codes.Aborted, "解析用户身份出错")
	}
	profile, err := p.Dao.GetProfile(accountId)
	if err != nil {
		return nil, status.Error(codes.Aborted, "获取用户认证信息出错")
	}

	if profile.Profile.Path != "" {
		re, _ := p.Oss.CreateUrl(ctx, &blobpb.CreateUrlRequest{
			Path:      profile.Profile.Path,
			Operation: blobpb.Operation_Status_DOWNLOAD,
		})
		profile.Profile.Path = re.Url
	}
	return &trippb.GetProfileResponse{Profile: profile.Profile}, nil
}
