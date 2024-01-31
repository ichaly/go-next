package oss

import (
	"fmt"
	ali "github.com/aliyun/aliyun-oss-go-sdk/oss"
	"github.com/ichaly/go-next/pkg/base"
	"io"
)

type AliYun struct {
	bucket    string
	domain    string
	region    string
	accessKey string
	secretKey string

	uploader *ali.Bucket
}

// https://help.aliyun.com/zh/oss/user-guide/regions-and-endpoints#section-plb-2vy-5db
func NewAliYun(c *base.Config) Uploader {
	return &AliYun{
		bucket:    c.Oss.Bucket,
		domain:    c.Oss.Domain,
		region:    c.Oss.Region,
		accessKey: c.Oss.AccessKey,
		secretKey: c.Oss.SecretKey,
	}
}

func (my *AliYun) Name() string {
	return "阿里"
}

func (my *AliYun) Init() error {
	client, err := ali.New(
		my.region, my.accessKey, my.secretKey,
		//ali.UseCname(true),
	)
	if err != nil {
		return err
	}
	bucket, err := client.Bucket(my.bucket)
	if err != nil {
		return err
	}
	my.uploader = bucket
	return nil
}

func (my *AliYun) Upload(data io.Reader, size int64, name string) (string, error) {
	err := my.uploader.PutObject(name, data)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%s/%s", my.domain, name), nil
}
