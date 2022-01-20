package auth_help

import (
	"context"
	"crypto/rsa"
	"fmt"
	"github.com/golang-jwt/jwt/v4"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
	"io/ioutil"
	token2 "project/service/shared/token"
	"strings"
)

const TOKEN_PREFIX = "Bearer "

type AccountId string
type AccountKey struct {
}

type interceptor struct {
	PublicKeyPath *rsa.PublicKey
}

func Interceptor(path string) grpc.UnaryServerInterceptor {
	key, err := getPublicKey(path)
	if err != nil {
		fmt.Println(err)
	}
	i := interceptor{PublicKeyPath: key}
	return i.Handler
}

func (i *interceptor) Handler(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
	token, err := getTokenByHeader(ctx)
	if err != nil {
		return "", err
	}

	validator := token2.JwtVerify{GetPublicKey: func() (pem interface{}, err error) {
		return i.PublicKeyPath, nil
	}}

	accountId, err := validator.Verify(token)

	if err != nil {
		return nil, err
	}

	context.WithValue(ctx, AccountKey{}, AccountId(accountId))

	return handler(ctx, req)
}

func getPublicKey(p string) (*rsa.PublicKey, error) {
	file, err := ioutil.ReadFile(p)
	if err != nil {
		return nil, status.Errorf(codes.Unauthenticated, "cannot find public key pem")
	}

	pem, err := jwt.ParseRSAPublicKeyFromPEM(file)
	if err != nil {
		return nil, status.Errorf(codes.Unauthenticated, "cannot parse public key pem")
	}
	return pem, nil
}

func getTokenByHeader(c context.Context) (string, error) {
	//使用metadata.FromIncomingContext 获取headers信息 返回map[string][]string
	//google.golang.org/grpc/metadata
	//type MD map[string][]string
	//func FromIncomingContext(ctx context.Context) (MD, bool)
	//incomingContext 格式大概如:map[:authority:[localhost:9003] authorization:[Bearer 666666] content-type:[application/grpc] grpcgateway-accept:[*/*] grpcgateway-authorization:[Bearer 666666] grpcgateway-content-type:[text/plain] grpcgateway-user-agent:[PostmanRuntime/7.29.0] user-agent:[grpc-go/1.40.0] x-forwarded-for:[127.0.0.1] x-forwarded-host:[localhost:9002]]
	incomingContext, _ := metadata.FromIncomingContext(c)
	fmt.Println(incomingContext)
	for k, v := range incomingContext {
		if k == "authorization" {
			for _, s := range v {
				if strings.HasPrefix(s, TOKEN_PREFIX) {
					return s[len(TOKEN_PREFIX):], nil
				}
			}
		}
	}
	return "", status.Errorf(codes.Unauthenticated, "header Authorization must has token")
}
