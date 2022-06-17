package main

import (
	"context"
	"crypto/rsa"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/encoding/protojson"
	"io/ioutil"
	"net/http"
	wechatpb "project/service/auth/api"
	"project/service/auth/dao"
	"project/service/auth/wechat"
	"project/service/shared/service"
	"project/service/shared/token"
	"time"
)

//需要实现服务接口
type Service struct {
	//需要写这个,这个是兼容用的，高版本的要写这个,不然会报错的
	wechatpb.UnimplementedLoginServiceServer
	Ws wechat.Service
	mg dao.Mg
	jt token.JWTToken
}

const privateKeyPath = "./service/auth/auth/private_key.pem"

func (s *Service) GetUserInfo(c context.Context, request *wechatpb.LoginRequest) (*wechatpb.LoginResponse, error) {
	development, err := zap.NewDevelopment()
	if err != nil {
		panic(err)
	}
	development.Info("get request", zap.String("code", request.Code))

	openID, err := s.Ws.Resolve(request.Code)

	if err != nil {
		return &wechatpb.LoginResponse{}, status.Error(codes.Unavailable, "你可以继续试试哈")
	}

	accountId, err := s.mg.ResolveAccountId(openID)

	if err != nil {
		return &wechatpb.LoginResponse{}, status.Error(codes.Unavailable, "openId insert mongodb fail")
	}

	development.Info("get response", zap.String("accountId", accountId))

	again := 2 * time.Hour

	token, err := s.jt.Create(accountId, again)

	if err != nil {
		return &wechatpb.LoginResponse{}, status.Errorf(codes.Unavailable, "can't create token")
	}

	development.Info("create token", zap.String("token", token))

	return &wechatpb.LoginResponse{Token: token, Aging: int64(again.Seconds())}, nil
}

func ReadPem(p string) *rsa.PrivateKey {
	file, err := ioutil.ReadFile(p)
	if err != nil {
		panic(err)
	}
	pem, err := jwt.ParseRSAPrivateKeyFromPEM(file)
	if err != nil {
		fmt.Println("can't parse privateKey")
	}
	return pem
}

//remark
func main() {
	service.RegisterRpcService(service.RpcServiceConfig{
		Name: "登录服务",
		Port: 9001,
		RegisterFunc: func(s *grpc.Server) {
			connect, err := mongo.Connect(context.Background(), options.Client().ApplyURI("mongodb://localhost:27017"))
			if err != nil {
				panic(err)
			}

			wechatpb.RegisterLoginServiceServer(s, &Service{Ws: wechat.Service{
				AppId:     "wxace898a3c3893f74",
				AppSecret: "6961b9bd55ca8b4ff3aed79af448d7c3",
			}, mg: dao.NewMongo(connect.Database("cool"), context.Background(), primitive.NewObjectID), jt: token.JWTToken{
				NowFunc: func() time.Time {
					return time.Now()
				},
				PrivateKey:   ReadPem(privateKeyPath),
				GetPublicKey: nil,
			}})
		},
	})
}

func createGateway() {

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

	//将grpc 服务器 与 grpc-gateway 关联
	//RegisterLoginServiceHandlerFromEndpoint  LoginService 是在proto 中 声明的服务名称 如果你声明的叫Xxx 那么这个方法现在就叫RegisterXxxHandlerFromEndpoint
	//第一个参数是传递的上下文
	//第二个参数是 http服务器
	//第三个参数是 grpc 服务端的端口
	//第四个参数是 拨号参数 这里使用不安全的连接方式
	//http 服务器与grpc 9001端口进行关联。通过拨号的方式进行通信
	wechatpb.RegisterLoginServiceHandlerFromEndpoint(c, mux, ":9001", []grpc.DialOption{grpc.WithInsecure()})

	//通过http.ListenAndServe 在本地 9002 端口 启动http服务器
	http.ListenAndServe("localhost:9002", mux)
}
