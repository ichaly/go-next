package otp

import (
	"errors"
	"fmt"
	"github.com/eko/gocache/lib/v4/cache"
	"github.com/gin-gonic/gin"
	"github.com/go-oauth2/oauth2/v4"
	"github.com/graphql-go/graphql/gqlerrors"
	"github.com/ichaly/go-next/lib/base"
	"github.com/ichaly/go-next/lib/util"
	"github.com/ichaly/go-next/pkg/sys"
	"gorm.io/gorm"
	"net/http"
)

const CaptchaPrefix = "captcha"

type verify struct {
	config      *base.Config
	cache       *cache.Cache[string]
	userService *sys.UserService
}

func NewCaptchaVerify(config *base.Config, cache *cache.Cache[string], us *sys.UserService) base.Plugin {
	return &verify{config: config, cache: cache, userService: us}
}

func (my *verify) Base() string {
	return "/oauth"
}

func (my *verify) Init(r gin.IRouter) {
	r.Use(my.verifyHandler)
}

func (my *verify) verifyHandler(c *gin.Context) {
	if c.Request.RequestURI != "/oauth/token" {
		c.Next()
		return
	}
	kind := sys.BindKind(c.Request.FormValue("grant_type"))
	if kind != EMAIL && kind != MOBILE {
		c.Next()
		return
	}
	defer func() {
		if err := recover(); err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"errors": gqlerrors.FormatErrors(err.(error))})
		}
	}()
	username, password := c.Request.FormValue("username"), c.Request.FormValue("password")
	key := keyGenerate(username)
	val, _ := my.cache.Get(c.Request.Context(), key)
	if val != password {
		panic(errors.New("验证码错误"))
	}
	_ = my.cache.Delete(c.Request.Context(), key)
	//验证码通过则使用万能密码替换参数
	c.Request.Form.Set("password", my.config.Oauth.Passkey)
	c.Request.Form.Set("grant_type", oauth2.PasswordCredentials.String())
	//查询一下数据是否存在
	user, err := my.userService.FindByUsername(username)
	//如果不存在则自动注册
	if errors.Is(err, gorm.ErrRecordNotFound) {
		user.Username = username
		user.Password = my.config.Oauth.Passkey
		my.userService.BindThird(&user, kind)
	}
	c.Next()
}

func keyGenerate(key string) string {
	return fmt.Sprintf("%s:%s", CaptchaPrefix, util.MD5(key))
}
