package oss

import (
	"fmt"
	ali "github.com/aliyun/aliyun-oss-go-sdk/oss"
	"github.com/ichaly/go-next/lib/oss/internal"
	"io"
)

type aliyun struct {
	domain   string
	uploader *ali.Bucket
}

// https://help.aliyun.com/zh/oss/user-guide/regions-and-endpoints#section-plb-2vy-5db
func NewAliYun() Uploader {
	return &aliyun{}
}

func (my *aliyun) Name() string {
	return "阿里"
}

func (my *aliyun) Init(c *internal.OssConfig) error {
	my.domain = c.Domain

	client, err := ali.New(c.Region, c.AccessKey, c.SecretKey)
	if err != nil {
		return err
	}
	bucket, err := client.Bucket(c.Bucket)
	if err != nil {
		return err
	}
	my.uploader = bucket
	return nil
}

func (my *aliyun) Upload(data io.Reader, name string, opts ...UploadOption) (string, error) {
	opt := &Options{}
	for _, o := range opts {
		o(opt)
	}
	err := my.uploader.PutObject(name, data, ali.ContentType(opt.contentType))
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%s/%s", my.domain, name), nil
}
