package totp

import (
	"errors"
	"github.com/eko/gocache/lib/v4/cache"
	"github.com/eko/gocache/lib/v4/store"
	"github.com/gin-gonic/gin"
	"github.com/graphql-go/graphql/gqlerrors"
	"github.com/ichaly/go-next/pkg/base"
	"github.com/ichaly/go-next/pkg/util"
	"net/http"
)

type captcha struct {
	config *base.Config
	cache  *cache.Cache[string]
}

func NewCaptcha(config *base.Config, cache *cache.Cache[string]) base.Plugin {
	return &captcha{config: config, cache: cache}
}

func (my *captcha) Base() string {
	return "/oauth"
}

func (my *captcha) Init(r gin.IRouter) {
	//快捷登录验证码
	r.Match([]string{http.MethodGet, http.MethodPost}, "/captcha", my.captchaHandler)
}

func (my *captcha) captchaHandler(c *gin.Context) {
	defer func() {
		if err := recover(); err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"errors": gqlerrors.FormatErrors(err.(error))})
		}
	}()
	req := struct {
		Username    string `form:"username" json:"username,omitempty"`
		CaptchaType string `form:"captcha_type" json:"captcha_type,omitempty"`
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
		val = util.RandomCode(my.config.Captcha.Length)
	}
	err = my.cache.Set(c.Request.Context(), key, val, store.WithExpiration(my.config.Captcha.Expired))
	if err != nil {
		panic(err)
	}
	c.JSON(http.StatusOK, gin.H{"msg": "操作成功"})
}
