package main

import (
	wechatpb "project/service/auth/api"
	profilepb "project/service/blob/api"
	trippb "project/service/rental/api"
	"project/service/shared/service"
)

func main() {
	configs := []service.RpcGateWayConfig{

		service.RpcGateWayConfig{
			Name:         "登录服务",
			Port:         9001,
			RegisterFunc: wechatpb.RegisterLoginServiceHandlerFromEndpoint,
		},
		service.RpcGateWayConfig{
			Name:         "行程服务",
			Port:         9003,
			RegisterFunc: trippb.RegisterTripServiceHandlerFromEndpoint,
		},
		service.RpcGateWayConfig{
			Name:         "图片服务",
			Port:         9006,
			RegisterFunc: profilepb.RegisterProfileServiceHandlerFromEndpoint,
		},
	}

	service.RegisterRpcGateWay(configs)
}
