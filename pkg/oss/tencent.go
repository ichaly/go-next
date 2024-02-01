package oss

import (
	"context"
	"fmt"
	"github.com/ichaly/go-next/pkg/base"
	"github.com/tencentyun/cos-go-sdk-v5"
	"io"
	"net/http"
	"net/url"
	"time"
)

type Tencent struct {
	bucket    string
	domain    string
	region    string
	accessKey string
	secretKey string

	client *cos.Client
}

func NewTencent(c *base.Config) Uploader {
	return &Tencent{
		bucket:    c.Oss.Bucket,
		domain:    c.Oss.Domain,
		region:    c.Oss.Region,
		accessKey: c.Oss.AccessKey,
		secretKey: c.Oss.SecretKey,
	}
}

func (my *Tencent) Name() string {
	return "腾讯"
}

func (my *Tencent) Init() error {
	u, _ := url.Parse(fmt.Sprintf("https://%s.cos.%s.myqcloud.com", my.bucket, my.region))
	b := &cos.BaseURL{BucketURL: u}
	my.client = cos.NewClient(b, &http.Client{
		//设置超时时间
		Timeout: 100 * time.Second,
		Transport: &cos.AuthorizationTransport{
			//如实填写账号和密钥，也可以设置为环境变量
			SecretID:  my.accessKey,
			SecretKey: my.secretKey,
		},
	})
	return nil
}

func (my *Tencent) Upload(data io.Reader, name string, opts ...UploadOption) (string, error) {
	opt := &Options{}
	for _, o := range opts {
		o(opt)
	}
	_, err := my.client.Object.Put(context.Background(), name, data, &cos.ObjectPutOptions{
		ObjectPutHeaderOptions: &cos.ObjectPutHeaderOptions{
			ContentType: opt.contentType,
		},
	})
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%s/%s", my.domain, name), nil
}
