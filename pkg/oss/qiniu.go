package oss

import (
	"context"
	"fmt"
	"github.com/ichaly/go-next/pkg/base"
	"github.com/qiniu/go-sdk/v7/auth"
	"github.com/qiniu/go-sdk/v7/storage"
	"io"
)

type QiNiu struct {
	bucket string
	domain string
	mac    *auth.Credentials
}

func NewQiNiu(c *base.Config) Uploader {
	my := &QiNiu{domain: c.Oss.Domain, bucket: c.Oss.Bucket}
	my.mac = auth.New(c.Oss.AccessKey, c.Oss.SecretKey)
	return my
}

func (my *QiNiu) Name() string {
	return "七牛"
}

func (my *QiNiu) AccessToken() string {
	policy := storage.PutPolicy{Scope: my.bucket}
	return policy.UploadToken(my.mac)
}

func (my *QiNiu) Upload(file io.Reader, size int64, name string) (string, error) {
	cfg := storage.Config{UseHTTPS: true, Region: &storage.ZoneHuadong}
	uploader := storage.NewFormUploader(&cfg)
	ret := storage.PutRet{}
	if err := uploader.Put(
		context.Background(), &ret, my.AccessToken(),
		name, file, size, &storage.PutExtra{},
	); err != nil {
		return "", err
	}
	return fmt.Sprintf("%s/%s", my.domain, ret.Key), nil
}
