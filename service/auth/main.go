package main

import (
	"context"
	"crypto/rsa"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"io/ioutil"
	wechatpb "project/service/auth/api"
	"project/service/auth/dao"
	"project/service/auth/wechat"
	grpc_help "project/service/shared/grpc"
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

	token, err := s.jt.Create(accountId, 2*time.Hour)

	if err != nil {
		return &wechatpb.LoginResponse{}, status.Errorf(codes.Unavailable, "can't create token")
	}

	development.Info("create token", zap.String("token", token))

	return &wechatpb.LoginResponse{Token: token}, nil
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

func main() {
	go createRpcService()
	createGateway()
}

func createRpcService() {

	manger := grpc_help.ServiceManger{
		NewWork: "tcp",
		Port:    "9001",
		CallBack: func(server grpc.ServiceRegistrar) {
			connect, err := mongo.Connect(context.Background(), options.Client().ApplyURI("mongodb://localhost:27017"))
			if err != nil {
				panic(err)
			}

			wechatpb.RegisterLoginServiceServer(server, &Service{Ws: wechat.Service{
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
	}

	manger.Start()
}

func createGateway() {
	manager := grpc_help.GateWayManager{
		Port: "9002",
		Mux:  grpc_help.CreateBaseMux(),
	}
	manager.RegisterPoint(wechatpb.RegisterLoginServiceHandlerFromEndpoint, "9001")
	manager.Start()
}
