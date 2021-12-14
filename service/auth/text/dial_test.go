package text

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	wechatpb "project/service/auth/api"
	"testing"
)

func TestDial(t *testing.T) {

	//拨号与grpc服务器进行连接
	dial, err := grpc.Dial(":9001", grpc.WithInsecure())
	background := context.Background()
	if err != nil {
		panic(err)
	}

	//在连接上创建客户端
	client := wechatpb.NewLoginServiceClient(dial)

	//客户端上调用服务方法获取响应
	info, err := client.GetUserInfo(background, &wechatpb.LoginRequest{Code: "666"})
	if err != nil {
		panic(info)
	}

	fmt.Printf("%+v", info)
}
