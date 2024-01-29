package oss

import (
	"github.com/gin-gonic/gin"
	"github.com/graphql-go/graphql/gqlerrors"
	"github.com/ichaly/go-next/pkg/base"
	"go.uber.org/fx"
	"mime/multipart"
	"net/http"
)

type Uploader interface {
	Name() string
	AccessToken() string
	Upload(file multipart.File, size int64, name string) (string, error)
}

type UploaderGroup struct {
	fx.In
	All []Uploader `group:"oss"`
}

type oss struct {
	uploader Uploader
}

func NewOss(c *base.Config, g UploaderGroup) base.Plugin {
	var u Uploader
	for _, v := range g.All {
		if v.Name() == c.Oss.Vendor {
			u = v
			break
		}
	}
	return &oss{uploader: u}
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
	file, header, err := c.Request.FormFile("file")
	if err != nil {
		panic(err)
	}
	key, err := my.uploader.Upload(file, header.Size, header.Filename)
	if err != nil {
		panic(err)
	}
	c.JSON(http.StatusOK, gin.H{"msg": "操作成功", "key": key})
}
