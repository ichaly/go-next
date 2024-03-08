package oss

import (
	"fmt"
	"github.com/ichaly/go-next/lib/base"
	"io"
	"os"
	"path"
	"strings"
)

type Local struct {
	bucket string
	domain string
}

func NewLocal(c *base.Config) Uploader {
	return &Local{
		bucket: c.Oss.Bucket,
		domain: c.Oss.Domain,
	}
}
func (my *Local) Name() string {
	return "本地"
}

func (my *Local) Init() error {
	return os.MkdirAll(my.bucket, os.ModePerm)
}

func (my *Local) Upload(data io.Reader, name string, opts ...UploadOption) (string, error) {
	full := path.Join(my.bucket, name)
	//保存前的路径检查创建
	tmp := strings.SplitAfter(full, "/")
	err := os.MkdirAll(path.Join(tmp[0:len(tmp)-1]...), os.ModePerm)
	if err != nil {
		return "", err
	}
	//保存文件
	dst, err := os.OpenFile(full, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, os.ModePerm)
	if err != nil {
		return "", err
	}
	defer func(dst *os.File) {
		_ = dst.Close()
	}(dst)
	_, err = io.Copy(dst, data)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%s/%s", my.domain, name), nil
}
