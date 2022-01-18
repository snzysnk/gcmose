package auth

import (
	"context"
	"crypto/rsa"
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

type Interceptor struct {
	PublicKeyPath string
}

func (i *Interceptor) Handler(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
	token, err := getTokenByHeader(ctx)
	if err != nil {
		return "", err
	}

	validator := token2.JwtVerify{GetPublicKey: func() (pem interface{}, err error) {
		return getPublicKey(i.PublicKeyPath)
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
	incomingContext, _ := metadata.FromIncomingContext(c)
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
