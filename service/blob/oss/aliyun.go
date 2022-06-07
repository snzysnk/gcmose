package myos

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/services/sts"
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
)

type OssInterface interface {
	GetCredentials(AccessKeyId string, AccessKeySecret string) (*sts.Credentials, error)
	CreateSingUrl(path string, operation int64) (string, error)
}

type OssService struct {
	OssInterface
	AccessKeyId     string
	AccessKeySecret string
}

func (o *OssService) CreateSingUrl(path string, operation int64) (string, error) {
	credentials, err := o.GetCredentials(o.AccessKeyId, o.AccessKeySecret)
	if err != nil {
		return "", err
	}

	/**
	  创建客户端
	  客户端需一个节点信息，这里选择了杭州阿里云节点
	  需要的临时key，secret，token从 sts.Credentials 中获取
	*/
	client, err := oss.New("http://oss-cn-hangzhou.aliyuncs.com", credentials.AccessKeyId, credentials.AccessKeySecret, oss.SecurityToken(credentials.SecurityToken))
	if err != nil {
		return "", err
	}

	bucketName := "study-golang"
	bucket, err := client.Bucket(bucketName)

	/**
	获取预签名url
	指定了该url存储位置，提交方式，有效时间，附加参数
	oss.HTTPGet 代表生成的url可以用来下载
	oss.HTTPPut 代表生成的url可以用来上传
	*/
	var options []oss.Option
	operationAction := oss.HTTPGet
	if operation == 2 {
		/**
		  这里以png图片为例，故此设置为 image/png
		*/
		options = append(options, oss.ContentType("image/png"))
		operationAction = oss.HTTPPut
	}

	signedURL, err := bucket.SignURL(path, operationAction, 6000, options...)

	return signedURL, err
}

func (o *OssService) GetCredentials(AccessKeyId string, AccessKeySecret string) (*sts.Credentials, error) {
	client, err := sts.NewClientWithAccessKey("cn-hangzhou", AccessKeyId, AccessKeySecret)

	if err != nil {
		return nil, err
	}

	request := sts.CreateAssumeRoleRequest()
	request.Scheme = "https"

	/**
	  访问 https://ram.console.aliyun.com/roles 可以看到
	  要保证该角色有权限操作oss
	  RoleArn 即 该角色的Arn
	  RoleSessionName 标识名称
	*/
	request.RoleArn = "acs:ram::1677026861345899:role/osstest"
	request.RoleSessionName = "ossTest"

	response, err := client.AssumeRole(request)
	if err != nil {
		return nil, err
	}

	/**
	  返回临时身份信息
	*/
	return &response.Credentials, nil
}
