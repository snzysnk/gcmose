package main

import (
	wechatpb "project/service/auth/api"
	trippb "project/service/rental/api"
	grpc_help "project/service/shared/grpc"
)

func main() {
	manager := grpc_help.GateWayManager{
		Port: "9002",
		Mux:  grpc_help.CreateBaseMux(),
	}
	manager.RegisterPoint(wechatpb.RegisterLoginServiceHandlerFromEndpoint, "9001")
	manager.RegisterPoint(trippb.RegisterTripServiceHandlerFromEndpoint, "9003")
	manager.Start()
}
