package wechat

import (
	"fmt"
	"github.com/medivhzhan/weapp/v2"
)

type Service struct {
	AppId     string
	AppSecret string
}

func (s *Service) Resolve(code string) (string, error) {
	login, err := weapp.Login(s.AppId, s.AppSecret, code)

	if err != nil {
		return "", fmt.Errorf("login error %v", err)
	}

	if err := login.GetResponseError(); err != nil {
		return "", fmt.Errorf("response error %v", err)
	}

	return login.OpenID, nil
}
