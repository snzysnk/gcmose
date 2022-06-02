package service

import (
	"context"
	"fmt"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/encoding/protojson"
	"net"
	"net/http"
	"project/service/shared"
	"project/service/shared/token"
	"strings"
)

type AccountKey struct {
}

type RpcServiceConfig struct {
	Name          string
	Port          int
	RegisterFunc  func(s *grpc.Server)
	ValidateToken bool
}

type RpcGateWayConfig struct {
	Name         string
	Port         int
	RegisterFunc func(ctx context.Context, mux *runtime.ServeMux, endpoint string, opts []grpc.DialOption) (err error)
}

func RegisterRpcService(config RpcServiceConfig) {

	listen, err := net.Listen("tcp", fmt.Sprintf(":%d", config.Port))
	if err != nil {
		panic(err)
	}

	var serverOptions []grpc.ServerOption

	if config.ValidateToken {
		i := func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
			tkn, err := getTokenFromContext(ctx)

			if err != nil {
				return "", err
			}

			subject, err := token.JwTVerify(tkn, "service/auth/auth/public_key.pem")

			if err != nil {
				return "", status.Error(codes.Unauthenticated, "")
			}

			return handler(createContextWithAccountId(ctx, shared.AccountId(subject)), req)
		}
		serverOptions = append(serverOptions, grpc.UnaryInterceptor(i))
	}

	server := grpc.NewServer(serverOptions...)
	config.RegisterFunc(server)
	fmt.Printf("启动 %s 服务Rpc", config.Name)
	server.Serve(listen)
}

func createContextWithAccountId(ctx context.Context, accountId shared.AccountId) context.Context {
	return context.WithValue(ctx, AccountKey{}, accountId)
}

func GetContextAccountId(ctx context.Context) (accountId shared.AccountId, err error) {
	account := ctx.Value(AccountKey{})
	accountId, ok := account.(shared.AccountId)
	if !ok {
		return "", status.Error(codes.Unauthenticated, "")
	}

	return accountId, nil
}

// 从context中获取token
func getTokenFromContext(ctx context.Context) (token string, err error) {
	unauthenticated := status.Error(codes.Unauthenticated, "")
	data, ok := metadata.FromIncomingContext(ctx)

	if !ok {
		return "", unauthenticated
	}
	//获取token
	tkn := ""

	for _, v := range data["authorization"] {
		if strings.HasPrefix(v, "Bearer ") {
			tkn = v[len("Bearer "):]
		}
	}

	if tkn == "" {
		return "", unauthenticated
	}

	return tkn, nil
}

func RegisterRpcGateWay(configs []RpcGateWayConfig) {

	//可以取消的上下文
	//为啥用context.WithCancel,因为这个服务一般在协程中启动,可以在某个协程异常时取消其关联的上下文
	background := context.Background()
	c, cancelFunc := context.WithCancel(background)
	defer cancelFunc()

	//通过runtime.NewServeMux 创建一个http服务器
	//runtime.NewServeMux 只有一个opts 参数 即服务器参数
	//runtime.WithMarshalerOption 返回一个转换关系的服务器参数
	//runtime.MIMEWildcard 即 "*" 对所有数据应用转换关系
	//runtime.JSONPb 对数据转换成json 时的 转换关系
	//MarshalOptions 具体关系参数
	//protojson.MarshalOptions 当proto 转换成 json 时 的规则

	mux := runtime.NewServeMux(runtime.WithMarshalerOption(runtime.MIMEWildcard, &runtime.JSONPb{
		MarshalOptions: protojson.MarshalOptions{
			UseEnumNumbers: true, //转换枚举变量 输出枚举变量值,false 则输出枚举变量名
			UseProtoNames:  true, //响应body 中 使用 proto 中声明的变量名,不然其用_组成的变量名变成大驼峰形式 如 my_logo => myLogo
			AllowPartial:   true, //允许请求报文中缺少字段
		},
	}))

	for _, v := range configs {
		fmt.Printf("绑定网关服务 :%s", v.Name)
		v.RegisterFunc(c, mux, fmt.Sprintf(":%d", v.Port), []grpc.DialOption{grpc.WithInsecure()})
	}

	//通过http.ListenAndServe 在本地 9002 端口 启动http服务器
	http.ListenAndServe("localhost:9002", mux)
}
