package auth

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/eko/gocache/lib/v4/cache"
	"github.com/gin-gonic/gin"
	"github.com/graphql-go/graphql/gqlerrors"
	"github.com/ichaly/go-next/app/sys"
	"github.com/ichaly/go-next/pkg/base"
	"github.com/ichaly/go-next/pkg/util"
	"gorm.io/gorm"
	"io"
	"net/http"
)

const CaptchaPrefix = "captcha"

type captcha struct {
	db     *gorm.DB
	config *base.Config
	cache  *cache.Cache[string]
}

func NewCaptcha(db *gorm.DB, config *base.Config, cache *cache.Cache[string]) base.Plugin {
	return &captcha{db: db, config: config, cache: cache}
}

func (my *captcha) Base() string {
	return "/oauth"
}

func (my *captcha) Init(r gin.IRouter) {
	r.Use(my.verifyHandler)
}

func (my *captcha) verifyHandler(c *gin.Context) {
	if c.Request.RequestURI != "/oauth/token" {
		c.Next()
		return
	}
	defer func() {
		if err := recover(); err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"errors": gqlerrors.FormatErrors(err.(error))})
		}
	}()
	req := struct {
		Code         string `json:"code,omitempty"`
		Password     string `json:"password,omitempty"`
		Username     string `json:"username,omitempty"`
		ClientId     string `json:"client_id,omitempty"`
		ClientSecret string `json:"client_secret,omitempty"`
		RedirectUri  string `json:"redirect_uri,omitempty"`
		GrantType    string `json:"grant_type,omitempty"`
		CaptchaType  string `json:"captcha_type,omitempty"`
	}{}
	body, _ := c.GetRawData()
	err := json.Unmarshal(body, &req)
	if err != nil {
		panic(err)
	}
	ct := sys.BindType(req.CaptchaType)
	if ct != sys.Email && ct != sys.Mobile {
		c.Next()
		return
	}
	key := keyGenerate(req.Password)
	val, err := my.cache.Get(c.Request.Context(), key)
	if err != nil {
		panic(err)
	}
	if val != req.Password {
		panic(errors.New("验证码错误"))
	}
	_ = my.cache.Delete(c.Request.Context(), key)
	req.Password = my.config.Oauth.Passkey
	req.ClientSecret = my.config.Oauth.Passkey
	req.CaptchaType = ""
	//自动注册
	user := sys.User{}
	my.db.Table("sys_user").
		Select("sys_user.id as id,sys_user.username as username").
		Joins("left join sys_bind on sys_bind.uid = sys_user.id and sys_user.username = ? or sys_bind.value = ?", req.Username, req.Username).
		First(&user)
	if user.ID == 0 {
		user.Username = req.Username
		user.Password = req.Password
		my.db.Save(user)
		bind := sys.Bind{
			Uid:   user.ID,
			Value: user.Username,
			Type:  sys.BindType(req.GrantType),
		}
		my.db.Save(bind)
	}
	//替换参数重写请求体
	res, _ := json.Marshal(req)
	c.Request.Body = io.NopCloser(bytes.NewBuffer(res))
	c.Next()
}

func keyGenerate(key string) string {
	return fmt.Sprintf("%s:%s", CaptchaPrefix, util.MD5(key))
}
