package test

import (
	"fmt"
	myos "project/service/blob/oss"
	"testing"
)

const AccessKeyId = "LTAI5t8msVaKSFucLeBwsVC6"
const AccessKeySecret = "4ellrwDdU7vREBXvlErPE5YN5FZwKn"

func TestAliYun(t *testing.T) {
	service := myos.OssService{}
	SingUrl, err := service.CreateSingUrl(AccessKeyId, AccessKeySecret, "cool/a.png", 1)
	if err != nil {
		t.Error(err)
	}

	fmt.Println(SingUrl)
}
