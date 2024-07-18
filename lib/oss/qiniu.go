package oss

import (
	"context"
	"errors"
	"fmt"
	"github.com/ichaly/go-next/lib/oss/internal"
	"github.com/qiniu/go-sdk/v7/auth"
	"github.com/qiniu/go-sdk/v7/storage"
	"io"
)

type qiniu struct {
	bucket   string
	domain   string
	mac      *auth.Credentials
	uploader *storage.FormUploader
}

func NewQiNiu() Uploader {
	return &qiniu{}
}

func (my *qiniu) Name() string {
	return "七牛"
}

func (my *qiniu) Init(c *internal.OssConfig) error {
	my.bucket = c.Bucket
	my.domain = c.Domain

	my.mac = auth.New(c.AccessKey, c.SecretKey)
	region, ok := storage.GetRegionByID(storage.RegionID(c.Region))
	if !ok {
		return errors.New("地区解析失败")
	}
	cfg := storage.Config{UseHTTPS: true, Region: &region}
	my.uploader = storage.NewFormUploader(&cfg)
	return nil
}

func (my *qiniu) Upload(data io.Reader, name string, opts ...UploadOption) (string, error) {
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

func (my *qiniu) accessToken() string {
	policy := storage.PutPolicy{Scope: my.bucket}
	return policy.UploadToken(my.mac)
}
