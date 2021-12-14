package test

import (
	"fmt"
	"github.com/golang-jwt/jwt/v4"
	"project/service/shared/token"
	"testing"
	"time"
)

const privateKey = `-----BEGIN RSA PRIVATE KEY-----
MIIEowIBAAKCAQEAkIrKO0FzcAm2qhmSD4UlUqUHwL0f0Qh0w6pPqtVg1WLAiMKs
A0ncQQdKiRBQO6U5PA9gMfPjraKy/cv+LzrdtQJGrT89LjayF50krylnwgN6Ml02
13Vu/43pdEfJbuQgbYIC45TfQIyfoWXeqCG64jg7ptIz9ueuqBFpNYh0ifjSpi1I
6Nppjb6SZ8QQYln6oySF1npU1py1zM7LZXmWvo1MZHoVakO3XgFG/5rEchQgb1Ku
zSRKjSGdyPCMciJHzmTasu1Yj7q5rlDaC1AB+gpTPijrRhZnjjVzn8Pdf6aAI+SW
1Hic3gQtKe6DguOFODV7+Z6bCLvjrNWnyyh9HwIDAQABAoIBAGexqfNouRAfZ9QQ
FCDePDBWBIHZLAIiNuJIRG8iNR3ggo+aWRzi2Agu2uIJugWQOsdDoJc1iJkAp7g7
qFbpvDQnNd0ECTgmNQ2Lxs2CvtUI+Zk4FifsL31Wqayp83CCaGEnbDsoiiEcwUhg
dKPZzuCLxv+evXjcLQiVtQyaoEB6EJeCEZ+DI1vrX1Sv9/KMuZeSM3UBc2EtfunP
X9ubOTVeVT3GS5CFpNEnONRRvUlfFoMlBPTCG1244YQPQOdBvzYXRrJNWb438vL3
soPI70x9ztpaa/KVnrSm+FAVgmh7ydHgZsCp1ob3/Y8HMKWptAc8i/TuOLK6TqUu
zf38WsECgYEA57zlJd1c0R6kWSaCsbKcjldkAELhQ1Df/E3zCz5tiSfTzY5/fYF3
/yP5zOik8w1AkDxju7MBkPuBBxrZ4qC6JbISDf47h7TZGv/CrL/zPlGgJtwIU18w
RSXqga59pBcDYDX2VNPDb9uSQYu33NpHinwumf5oy9AeBPpy2pRwPfsCgYEAn6zZ
nrO8TkHOrG0frEvJ5u+7wd4O5SoVPL6YcR5+HAwezQvuFaEAvwdNxXjeSXiI/WQM
qCBTgtC42v+Kdh7tsghWeePvTM5p/6xc7LhUbXekyVPrgOANM0vRQYf1fqLMlLm7
8Wlg2UiVhwtN9QogZDsWVDrtE+jXtaqaR4DNSC0CgYEApMQs1ThIFzPHPM/h2wCY
yagfJQRob9lf+F9f8F2pufxEXwHzacQVEqZ8pRHAvkNs5WDatIk5EuGPwwJ3auhG
kijDJ4ncv8d5GfZBb1xnUabtzNNjZAvpITEtRJlRIctvDggwJe4PJkgi/FGgH/pf
JXrYE4vM7SUDQxzjbsBd8eMCgYBTO5TMx0AXQvJqqw9hjrgWK00iB5CYIFGlsmxg
Otbd4+uCUv+uQpjS+wmtMA4fGje+bhKy6PbfL9HMRnx0ija5IOOC8i/hR4+1eLGP
q55G2aEKWTZfnWsEHDdeA3MSqeYeeJrhgl67pLxdAqWGX9r2pF5FVdoTEMKwrddy
Ak//mQKBgGqdNy/tIIGjNiUQbho9+i3/sbElcfsZJSh+W26NOOaBAaYNebnOeaEa
xxNBA0YZzCL9XJJqrg3ClHSAJwhMXethVR/nRLJTbyfVMLwvJlwRP+sGmQw/tcEz
+swVUQH04DayL60EbEC3+VztYGfeXLz2ljQ5OvVsY1Ga0dTtky5e
-----END RSA PRIVATE KEY-----`

const publicKey = `-----BEGIN PUBLIC KEY-----
MIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEAkIrKO0FzcAm2qhmSD4Ul
UqUHwL0f0Qh0w6pPqtVg1WLAiMKsA0ncQQdKiRBQO6U5PA9gMfPjraKy/cv+Lzrd
tQJGrT89LjayF50krylnwgN6Ml0213Vu/43pdEfJbuQgbYIC45TfQIyfoWXeqCG6
4jg7ptIz9ueuqBFpNYh0ifjSpi1I6Nppjb6SZ8QQYln6oySF1npU1py1zM7LZXmW
vo1MZHoVakO3XgFG/5rEchQgb1KuzSRKjSGdyPCMciJHzmTasu1Yj7q5rlDaC1AB
+gpTPijrRhZnjjVzn8Pdf6aAI+SW1Hic3gQtKe6DguOFODV7+Z6bCLvjrNWnyyh9
HwIDAQAB
-----END PUBLIC KEY-----`

func TestCreateJwt(t *testing.T) {
	//将私钥转换成 *rsa.PrivateKey 类型
	pem, err2 := jwt.ParseRSAPrivateKeyFromPEM([]byte(privateKey))
	if err2 != nil {
		t.Error("cannot parse private key")
	}

	timeFunc := func() time.Time {
		return time.Now()
	}

	jwtClient := token.JWTToken{
		timeFunc,
		pem,
		nil,
	}

	token, err := jwtClient.Create("19940113", time.Hour*2)

	if err != nil {
		t.Errorf("cannot create token by jwt %+v", err)
	}

	fmt.Println(token)
}

func TestVerifyJwt(t *testing.T) {
	want := "eyJhbGciOiJSUzUxMiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2MzkzODgzNTgsImlhdCI6MTYzOTM4MTE1OCwiaXNzIjoiWFJYIiwic3ViIjoiMTk5NDAxMTMifQ.FENgt91Tyjwx94dC2hiBkuDetRCsBD6WEwf26gy59bakurqjb6yk550P22Bjl5Zam8godEdYo8mo_Qkph-24J_Jk-ATFlbBjwDo21BRZp4eKMW9ppN8zu92OElP-_jGIbWuQ-USUQp36qC5wWLUCAB5yfJw3Musjg1qgxYetzuEKetwmwXNEQbtyVOCT2FhajTlBxm8nh0skd8Ud_0BPMJPYQsAM-0Q1vsOUV7GW2mPYmbCW3BM9lvQOHqm8xuQagwjnXuFam5TJBx_J6GBnxSJkmK0KVaSX0YF7B1y0U3QZSKTWBOBkKf7DbQm0-VWa5qdC9WLh6mk9ZjZlECpgsQ"
	jwtToken := token.JWTToken{
		NowFunc:    nil,
		PrivateKey: nil,
		GetPublicKey: func() (i interface{}, err error) {
			return jwt.ParseRSAPublicKeyFromPEM([]byte(publicKey))
		},
	}

	accountId, err := jwtToken.Verify(want)
	if err != nil {
		t.Errorf("%v", err)
	}

	fmt.Println(accountId)
}
