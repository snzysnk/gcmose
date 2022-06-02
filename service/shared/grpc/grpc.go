package grpc_help

import (
	"context"
	"fmt"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/encoding/protojson"
	"net"
	"net/http"
)

type ServiceManger struct {
	NewWork     string
	Port        string
	CallBack    func(server grpc.ServiceRegistrar)
	Interceptor grpc.UnaryServerInterceptor
}

func (s *ServiceManger) Start() {
	listen, err := net.Listen(s.NewWork, fmt.Sprintf(":%s", s.Port))
	if err != nil {
		panic(err)
	}
	i := s.Interceptor

	//拦截器在创建grpc服务时通过传参加入
	//type UnaryServerInterceptor func(ctx context.Context, req interface{}, info *UnaryServerInfo, handler UnaryHandler) (resp interface{}, err error)
	//i 是 UnaryServerInterceptor 函数类型
	//其中ctx 是 context
	//req 是请求信息
	//info 是一些其它信息
	//handler 是下一个要执行的流程,return handler(ctx,req) 代表执行下一个流程，如果输出了的error非nil,代表拦截中断流程
	server := grpc.NewServer(grpc.UnaryInterceptor(i))
	s.CallBack(server)
	server.Serve(listen)
}

type registerFunc func(ctx context.Context, mux *runtime.ServeMux, endpoint string, opts []grpc.DialOption) (err error)

type GateWayManager struct {
	Port string
	Mux  *runtime.ServeMux
}

func (s *GateWayManager) RegisterPoint(f registerFunc, port string) {
	err := f(context.Background(), s.Mux, fmt.Sprintf(":%s", port), []grpc.DialOption{grpc.WithInsecure()})
	if err != nil {
		panic(err)
	}
}

func (s *GateWayManager) Start() {
	http.ListenAndServe(fmt.Sprintf("localhost:%s", s.Port), s.Mux)
}

func CreateBaseMux() *runtime.ServeMux {
	background := context.Background()
	_, cancelFunc := context.WithCancel(background)
	defer cancelFunc()

	mux := runtime.NewServeMux(runtime.WithMarshalerOption(runtime.MIMEWildcard, &runtime.JSONPb{
		MarshalOptions: protojson.MarshalOptions{
			UseEnumNumbers: true,
			UseProtoNames:  true,
			AllowPartial:   true,
		},
	}))

	return mux
}
