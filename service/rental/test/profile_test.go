package test

import (
	"context"
	"fmt"
	myos "project/service/blob/oss"
	trippb "project/service/rental/api"
	"project/service/rental/trip"
	"project/service/shared"
	"project/service/shared/mgutil"
	service2 "project/service/shared/service"
	"testing"
)

const AccessKeyId = "LTAI5t8msVaKSFucLeBwsVC6"
const AccessKeySecret = "4ellrwDdU7vREBXvlErPE5YN5FZwKn"

func TestProfile(t *testing.T) {
	service := trip.ProfileService{
		Dao: trip.ProfileServiceDao{
			Database: mgutil.NewMongoDatabaseClient("cool"),
		},
		Oss: myos.OssService{
			AccessKeyId:     AccessKeyId,
			AccessKeySecret: AccessKeySecret,
		},
	}
	c := context.Background()
	newC := context.WithValue(c, service2.AccountKey{}, shared.AccountId("99698"))
	profile, err := service.GetUploadUrl(newC, &trippb.GetUploadUrlRequest{})
	if err != nil {
		t.Error(err)
	}
	fmt.Println(profile.Url)

	check, err := service.ProfileCheck(newC, &trippb.ProfileCheckRequest{})
	if err != nil {
		t.Error(err)
	}
	fmt.Println(check)

	getProfile, err := service.GetProfile(newC, &trippb.GetProfileRequest{})
	if err != nil {
		t.Error(err)
	}
	fmt.Println(getProfile)
}
