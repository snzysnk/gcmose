package trip

import (
	"context"
	"fmt"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"io/ioutil"
	"net/http"
	myos "project/service/blob/oss"
	trippb "project/service/rental/api"
	"project/service/shared/service"
	"time"
)

type ProfileService struct {
	trippb.UnimplementedProfileServiceServer
	Dao ProfileServiceDao
	Oss myos.OssService
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
	url, err := p.Oss.CreateSingUrl(path, 2)
	if err != nil {
		return nil, status.Error(codes.Aborted, "不能创建临时上传连接")
	}
	return &trippb.GetUploadUrlResponse{Url: url}, nil
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
		url, err := p.Oss.CreateSingUrl(path, 1)

		response, err := http.Get(url)
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
		url, _ := p.Oss.CreateSingUrl(profile.Profile.Path, 1)
		profile.Profile.Path = url
	}
	return &trippb.GetProfileResponse{Profile: profile.Profile}, nil
}
