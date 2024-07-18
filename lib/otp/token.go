package otp

import (
	"errors"
	"github.com/eko/gocache/lib/v4/cache"
	"github.com/eko/gocache/lib/v4/store"
	"github.com/gin-gonic/gin"
	"github.com/graphql-go/graphql/gqlerrors"
	"github.com/ichaly/go-next/lib/base"
	"github.com/ichaly/go-next/lib/otp/internal"
	"github.com/ichaly/go-next/lib/util"
	"github.com/spf13/viper"
	"net/http"
)

type token struct {
	senders []Sender
	config  *internal.CaptchaConfig
	cache   *cache.Cache[string]
}

func NewToken(v *viper.Viper, c *cache.Cache[string], g SenderGroup) (base.Plugin, error) {
	config := &internal.CaptchaConfig{}
	if err := v.Sub("token").Unmarshal(config); err != nil {
		return nil, err
	}
	return &token{config: config, cache: c, senders: g.All}, nil
}

func (my *token) Base() string {
	return "/oauth"
}

func (my *token) Init(r gin.IRouter) {
	//快捷登录验证码
	r.Match([]string{http.MethodGet, http.MethodPost}, "/token", my.captchaHandler)
}

func (my *token) captchaHandler(c *gin.Context) {
	defer func() {
		if err := recover(); err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"errors": gqlerrors.FormatErrors(err.(error))})
		}
	}()
	req := struct {
		Username string `form:"username" json:"username,omitempty"`
		Category string `form:"category" json:"category,omitempty"`
	}{}
	err := c.ShouldBind(&req)
	if req.Username == "" {
		panic(errors.New("手机号/邮箱不能为空"))
	}
	if err != nil {
		panic(err)
	}
	key := keyGenerate(req.Username)
	val, err := my.cache.Get(c.Request.Context(), key)
	if err != nil || val == "" {
		val = util.RandomCode(my.config.Length)
	}
	err = my.cache.Set(c.Request.Context(), key, val, store.WithExpiration(my.config.Expired))
	if err != nil {
		panic(err)
	}
	//发送验证码
	for _, d := range my.senders {
		if d.Support(req.Category) {
			if err = d.Execute(req.Username, val); err != nil {
				panic(err)
			}
			c.JSON(http.StatusOK, gin.H{"msg": "操作成功"})
			return
		}
	}
	panic(errors.New("操作失败"))
}
