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
	handler := s.Interceptor
	server := grpc.NewServer(grpc.UnaryInterceptor(handler))
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
