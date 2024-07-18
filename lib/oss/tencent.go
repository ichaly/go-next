package oss

import (
	"context"
	"fmt"
	"github.com/ichaly/go-next/lib/oss/internal"
	"github.com/tencentyun/cos-go-sdk-v5"
	"io"
	"net/http"
	"net/url"
	"time"
)

type tencent struct {
	domain string
	client *cos.Client
}

func NewTencent() Uploader {
	return &tencent{}
}

func (my *tencent) Name() string {
	return "腾讯"
}

func (my *tencent) Init(c *internal.OssConfig) error {
	my.domain = c.Domain

	u, _ := url.Parse(fmt.Sprintf("https://%s.cos.%s.myqcloud.com", c.Bucket, c.Region))
	b := &cos.BaseURL{BucketURL: u}
	my.client = cos.NewClient(b, &http.Client{
		//设置超时时间
		Timeout: 100 * time.Second,
		Transport: &cos.AuthorizationTransport{
			//如实填写账号和密钥，也可以设置为环境变量
			SecretID:  c.AccessKey,
			SecretKey: c.SecretKey,
		},
	})
	return nil
}

func (my *tencent) Upload(data io.Reader, name string, opts ...UploadOption) (string, error) {
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
