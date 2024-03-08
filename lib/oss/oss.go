package oss

import (
	"bytes"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/graphql-go/graphql/gqlerrors"
	"github.com/h2non/filetype"
	"github.com/ichaly/go-next/lib/base"
	"github.com/ichaly/go-next/lib/util"
	"go.uber.org/fx"
	"golang.org/x/sync/errgroup"
	"io"
	"mime/multipart"
	"net/http"
	"path"
	"strconv"
	"strings"
)

const (
	KEY_FILE   = "file"
	KEY_PATH   = "path"
	KEY_RENAME = "rename"
)

type Uploader interface {
	Name() string
	Init() error
	Upload(data io.Reader, name string, opts ...UploadOption) (string, error)
}

type UploaderGroup struct {
	fx.In
	All []Uploader `group:"oss"`
}

type oss struct {
	uploader Uploader
}

func NewOss(c *base.Config, g UploaderGroup) (base.Plugin, error) {
	var u Uploader
	for _, v := range g.All {
		if v.Name() == c.Oss.Vendor {
			err := v.Init()
			if err != nil {
				return nil, err
			}
			u = v
			break
		}
	}
	if u == nil {
		return nil, errors.New("存储服务初始化失败")
	}
	return &oss{uploader: u}, nil
}

func (my *oss) Base() string {
	return "/oss"
}

func (my *oss) Init(r gin.IRouter) {
	r.POST("/upload", my.uploadHandler)
}

func (my *oss) uploadHandler(c *gin.Context) {
	defer func() {
		if err := recover(); err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"errors": gqlerrors.FormatErrors(err.(error))})
		}
	}()
	form, err := c.MultipartForm()
	if err != nil {
		panic(err)
	}
	files := form.File[KEY_FILE]
	if len(files) == 0 {
		panic(errors.New("上传的文件不能为空!"))
	}
	folder := c.PostForm(KEY_PATH)
	rename, err := strconv.ParseBool(c.PostForm(KEY_RENAME))
	//默认是自动重命名
	if err != nil {
		rename = true
	}
	var urls []string
	var eg errgroup.Group
	for _, f := range files {
		f := f // https://golang.org/doc/faq#closures_and_goroutines
		eg.Go(func() error {
			url, err := my.doUpload(f, folder, rename)
			if err != nil {
				return err
			}
			urls = append(urls, url)
			return nil
		})
	}
	err = eg.Wait()
	if err != nil {
		panic(err)
	}
	c.JSON(http.StatusOK, gin.H{"msg": "操作成功", "urls": urls})
}

func (my *oss) doUpload(header *multipart.FileHeader, folder string, rename bool) (string, error) {
	file, err := header.Open()
	if err != nil {
		return "", err
	}
	defer func(file multipart.File) {
		_ = file.Close()
	}(file)

	size := header.Size
	name := header.Filename
	var data bytes.Buffer
	tee := io.TeeReader(file, &data)
	key := util.Md5File(tee)
	kind, _ := filetype.Match(data.Bytes())
	if kind == filetype.Unknown {
		panic(errors.New("unknown file type"))
	}
	if rename {
		name = fmt.Sprintf("%s.%s", key, kind.Extension)
	}
	//拼接路径
	name = path.Join(folder, name)
	//移除多余的斜杠
	name = strings.TrimPrefix(name, "/")
	//执行文件上传
	url, err := my.uploader.Upload(&data, name, WithSize(size), WithContentType(kind.MIME.Value))
	if err != nil {
		return "", err
	}
	return url, nil
}
