package oss

import (
	"context"
	"fmt"
	"github.com/ichaly/go-next/pkg/base"
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
	"io"
)

type Minio struct {
	bucket    string
	region    string
	domain    string
	accessKey string
	secretKey string

	client *minio.Client
}

func NewMinio(c *base.Config) Uploader {
	return &Minio{
		bucket:    c.Oss.Bucket,
		domain:    c.Oss.Domain,
		region:    c.Oss.Region,
		accessKey: c.Oss.AccessKey,
		secretKey: c.Oss.SecretKey,
	}
}

func (my *Minio) Name() string {
	return "MINIO"
}

func (my *Minio) Init() error {
	client, err := minio.New(my.region, &minio.Options{
		Creds:  credentials.NewStaticV4(my.accessKey, my.secretKey, ""),
		Secure: false,
	})
	if err != nil {
		return err
	}
	my.client = client
	return nil
}

func (my *Minio) Upload(data io.Reader, name string, opts ...UploadOption) (string, error) {
	opt := &Options{}
	for _, o := range opts {
		o(opt)
	}
	res, err := my.client.PutObject(
		context.Background(), my.bucket, name, data, opt.size,
		minio.PutObjectOptions{ContentType: opt.contentType},
	)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%s/%s/%s", my.domain, res.Bucket, res.Key), nil
}
