package token

import (
	"crypto/rsa"
	"fmt"
	"github.com/golang-jwt/jwt/v4"
	"time"
)

//const privateKey = `-----BEGIN RSA PRIVATE KEY-----
//MIIEowIBAAKCAQEAkIrKO0FzcAm2qhmSD4UlUqUHwL0f0Qh0w6pPqtVg1WLAiMKs
//A0ncQQdKiRBQO6U5PA9gMfPjraKy/cv+LzrdtQJGrT89LjayF50krylnwgN6Ml02
//13Vu/43pdEfJbuQgbYIC45TfQIyfoWXeqCG64jg7ptIz9ueuqBFpNYh0ifjSpi1I
//6Nppjb6SZ8QQYln6oySF1npU1py1zM7LZXmWvo1MZHoVakO3XgFG/5rEchQgb1Ku
//zSRKjSGdyPCMciJHzmTasu1Yj7q5rlDaC1AB+gpTPijrRhZnjjVzn8Pdf6aAI+SW
//1Hic3gQtKe6DguOFODV7+Z6bCLvjrNWnyyh9HwIDAQABAoIBAGexqfNouRAfZ9QQ
//FCDePDBWBIHZLAIiNuJIRG8iNR3ggo+aWRzi2Agu2uIJugWQOsdDoJc1iJkAp7g7
//qFbpvDQnNd0ECTgmNQ2Lxs2CvtUI+Zk4FifsL31Wqayp83CCaGEnbDsoiiEcwUhg
//dKPZzuCLxv+evXjcLQiVtQyaoEB6EJeCEZ+DI1vrX1Sv9/KMuZeSM3UBc2EtfunP
//X9ubOTVeVT3GS5CFpNEnONRRvUlfFoMlBPTCG1244YQPQOdBvzYXRrJNWb438vL3
//soPI70x9ztpaa/KVnrSm+FAVgmh7ydHgZsCp1ob3/Y8HMKWptAc8i/TuOLK6TqUu
//zf38WsECgYEA57zlJd1c0R6kWSaCsbKcjldkAELhQ1Df/E3zCz5tiSfTzY5/fYF3
///yP5zOik8w1AkDxju7MBkPuBBxrZ4qC6JbISDf47h7TZGv/CrL/zPlGgJtwIU18w
//RSXqga59pBcDYDX2VNPDb9uSQYu33NpHinwumf5oy9AeBPpy2pRwPfsCgYEAn6zZ
//nrO8TkHOrG0frEvJ5u+7wd4O5SoVPL6YcR5+HAwezQvuFaEAvwdNxXjeSXiI/WQM
//qCBTgtC42v+Kdh7tsghWeePvTM5p/6xc7LhUbXekyVPrgOANM0vRQYf1fqLMlLm7
//8Wlg2UiVhwtN9QogZDsWVDrtE+jXtaqaR4DNSC0CgYEApMQs1ThIFzPHPM/h2wCY
//yagfJQRob9lf+F9f8F2pufxEXwHzacQVEqZ8pRHAvkNs5WDatIk5EuGPwwJ3auhG
//kijDJ4ncv8d5GfZBb1xnUabtzNNjZAvpITEtRJlRIctvDggwJe4PJkgi/FGgH/pf
//JXrYE4vM7SUDQxzjbsBd8eMCgYBTO5TMx0AXQvJqqw9hjrgWK00iB5CYIFGlsmxg
//Otbd4+uCUv+uQpjS+wmtMA4fGje+bhKy6PbfL9HMRnx0ija5IOOC8i/hR4+1eLGP
//q55G2aEKWTZfnWsEHDdeA3MSqeYeeJrhgl67pLxdAqWGX9r2pF5FVdoTEMKwrddy
//Ak//mQKBgGqdNy/tIIGjNiUQbho9+i3/sbElcfsZJSh+W26NOOaBAaYNebnOeaEa
//xxNBA0YZzCL9XJJqrg3ClHSAJwhMXethVR/nRLJTbyfVMLwvJlwRP+sGmQw/tcEz
//+swVUQH04DayL60EbEC3+VztYGfeXLz2ljQ5OvVsY1Ga0dTtky5e
//-----END RSA PRIVATE KEY-----`
//
//const publicKey = `-----BEGIN PUBLIC KEY-----
//MIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEAkIrKO0FzcAm2qhmSD4Ul
//UqUHwL0f0Qh0w6pPqtVg1WLAiMKsA0ncQQdKiRBQO6U5PA9gMfPjraKy/cv+Lzrd
//tQJGrT89LjayF50krylnwgN6Ml0213Vu/43pdEfJbuQgbYIC45TfQIyfoWXeqCG6
//4jg7ptIz9ueuqBFpNYh0ifjSpi1I6Nppjb6SZ8QQYln6oySF1npU1py1zM7LZXmW
//vo1MZHoVakO3XgFG/5rEchQgb1KuzSRKjSGdyPCMciJHzmTasu1Yj7q5rlDaC1AB
//+gpTPijrRhZnjjVzn8Pdf6aAI+SW1Hic3gQtKe6DguOFODV7+Z6bCLvjrNWnyyh9
//HwIDAQAB
//-----END PUBLIC KEY-----`

type JWTToken struct {
	NowFunc      func() time.Time
	PrivateKey   *rsa.PrivateKey
	GetPublicKey func() (interface{}, error)
}

func (s *JWTToken) Create(subject string, timeOut time.Duration) (string, error) {

	now := s.NowFunc()

	//使用标准字段声明 payload
	claims := jwt.StandardClaims{
		ExpiresAt: now.Unix() + int64(timeOut.Seconds()), //过期时间
		IssuedAt:  now.Unix(),                            //颁发时间
		Issuer:    "XRX",                                 //发行人
		Subject:   subject,                               //主题
	}

	//设置jwt header(主要设置加密方式) 和 payload
	//创建jwt
	withClaims := jwt.NewWithClaims(jwt.SigningMethodRS512, claims)

	//开始签名
	//如果 使用 rsa 加密方式
	//withClaims.SignedString() 的参数类型必须 是 *rsa.PrivateKey
	//返回完整的string header.payload.signature
	return withClaims.SignedString(s.PrivateKey)
}

func (s *JWTToken) Verify(signed string) (string, error) {

	//签名是否有效，是否过期
	//jwt.ParseWithClaims 有三个参数
	//第一个参数是要验证的token 字符串
	//第二个参数是claims,会往claims写入数据
	//jwt.ParseWithClaims 第三个参数是获取密钥的函数
	//获取密钥为啥是函数呢(只有当读取到header时，才知道加密方式，才知道是否需要密钥,获取密钥有可能是io 操作)
	claims, err := jwt.ParseWithClaims(signed, &jwt.StandardClaims{}, func(token *jwt.Token) (i interface{}, err error) {
		return s.GetPublicKey()
	})

	if err != nil {
		return "", fmt.Errorf("cannot parse token")
	}

	//token 是否有效
	if !claims.Valid {
		return "", fmt.Errorf("token is not valid")
	}

	//验证jwt.ParseWithClaims 是否将&jwt.StandardClaims 正确的赋值到token.Claims 中
	//两者应该是同一个
	standardClaims, ok := claims.Claims.(*jwt.StandardClaims)

	if !ok {
		return "", fmt.Errorf("not eq standardClaims")
	}

	return standardClaims.Subject, nil
}
