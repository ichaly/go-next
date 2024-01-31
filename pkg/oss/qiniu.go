package oss

import (
	"context"
	"errors"
	"fmt"
	"github.com/ichaly/go-next/pkg/base"
	"github.com/qiniu/go-sdk/v7/auth"
	"github.com/qiniu/go-sdk/v7/storage"
	"io"
)

type QiNiu struct {
	bucket    string
	domain    string
	region    string
	accessKey string
	secretKey string

	mac      *auth.Credentials
	uploader *storage.FormUploader
}

func NewQiNiu(c *base.Config) Uploader {
	return &QiNiu{
		bucket:    c.Oss.Bucket,
		domain:    c.Oss.Domain,
		region:    c.Oss.Region,
		accessKey: c.Oss.AccessKey,
		secretKey: c.Oss.SecretKey,
	}
}

func (my *QiNiu) Name() string {
	return "七牛"
}

func (my *QiNiu) Init() error {
	my.mac = auth.New(my.accessKey, my.secretKey)
	region, ok := storage.GetRegionByID(storage.RegionID(my.region))
	if !ok {
		return errors.New("地区解析失败")
	}
	cfg := storage.Config{UseHTTPS: true, Region: &region}
	my.uploader = storage.NewFormUploader(&cfg)
	return nil
}

func (my *QiNiu) Upload(data io.Reader, name string, opts ...UploadOption) (string, error) {
	opt := &Options{}
	for _, o := range opts {
		o(opt)
	}
	ret := storage.PutRet{}
	if err := my.uploader.Put(
		context.Background(), &ret, my.accessToken(),
		name, data, opt.size, &storage.PutExtra{},
	); err != nil {
		return "", err
	}
	return fmt.Sprintf("%s/%s", my.domain, ret.Key), nil
}

func (my *QiNiu) accessToken() string {
	policy := storage.PutPolicy{Scope: my.bucket}
	return policy.UploadToken(my.mac)
}
