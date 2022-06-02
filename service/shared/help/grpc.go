package help

import (
	"context"
	"fmt"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/encoding/protojson"
	"net/http"
)

type Register func(ctx context.Context, mux *runtime.ServeMux, endpoint string, opts []grpc.DialOption) (err error)

type GateWay struct {
	Address string
	Mux     *runtime.ServeMux
}

func (s *GateWay) Register(port string, registerFunc Register) {
	background := context.Background()
	c, cancelFunc := context.WithCancel(background)
	defer cancelFunc()
	registerFunc(c, s.Mux, fmt.Sprintf(":%s", port), []grpc.DialOption{grpc.WithInsecure()})
}

func (s *GateWay) Start() {
	http.ListenAndServe(s.Address, s.Mux)
}

func GetBaseMux() *runtime.ServeMux {
	mux := runtime.NewServeMux(runtime.WithMarshalerOption(runtime.MIMEWildcard, &runtime.JSONPb{
		MarshalOptions: protojson.MarshalOptions{
			UseEnumNumbers: true,
			UseProtoNames:  true,
			AllowPartial:   true,
		},
	}))

	return mux
}
