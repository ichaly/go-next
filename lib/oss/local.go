package oss

import (
	"fmt"
	"github.com/ichaly/go-next/lib/oss/internal"
	"io"
	"os"
	"path"
	"strings"
)

type local struct {
	bucket string
	domain string
}

func NewLocal() Uploader {
	return &local{}
}
func (my *local) Name() string {
	return "本地"
}

func (my *local) Init(c *internal.OssConfig) error {
	my.bucket = c.Bucket
	my.domain = c.Domain

	return os.MkdirAll(my.bucket, os.ModePerm)
}

func (my *local) Upload(data io.Reader, name string, opts ...UploadOption) (string, error) {
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
