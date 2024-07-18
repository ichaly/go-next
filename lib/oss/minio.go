package oss

import (
	"context"
	"fmt"
	"github.com/ichaly/go-next/lib/oss/internal"
	mio "github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
	"io"
)

type minio struct {
	bucket string
	domain string
	client *mio.Client
}

func NewMinio() Uploader {
	return &minio{}
}

func (my *minio) Name() string {
	return "MINIO"
}

func (my *minio) Init(c *internal.OssConfig) error {
	my.bucket = c.Bucket
	my.domain = c.Domain

	client, err := mio.New(c.Region, &mio.Options{
		Creds:  credentials.NewStaticV4(c.AccessKey, c.SecretKey, ""),
		Secure: false,
	})
	if err != nil {
		return err
	}
	my.client = client
	return nil
}

func (my *minio) Upload(data io.Reader, name string, opts ...UploadOption) (string, error) {
	opt := &Options{}
	for _, o := range opts {
		o(opt)
	}
	res, err := my.client.PutObject(
		context.Background(), my.bucket, name, data, opt.size,
		mio.PutObjectOptions{ContentType: opt.contentType},
	)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%s/%s/%s", my.domain, res.Bucket, res.Key), nil
}
