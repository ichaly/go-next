package auth

import (
	"database/sql"
	"errors"
	"fmt"
	"github.com/eko/gocache/lib/v4/cache"
	"github.com/gin-gonic/gin"
	"github.com/go-oauth2/oauth2/v4"
	"github.com/graphql-go/graphql/gqlerrors"
	"github.com/ichaly/go-next/app/sys"
	"github.com/ichaly/go-next/pkg/base"
	"github.com/ichaly/go-next/pkg/util"
	"gorm.io/gorm"
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
	gt := sys.BindType(c.Request.FormValue("grant_type"))
	if gt != sys.Email && gt != sys.Mobile {
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
	user := sys.User{Username: username, Password: my.config.Oauth.Passkey}
	err := my.db.Table("sys_user").Select("sys_user.id,sys_user.username").
		Joins("left join sys_bind b on b.uid = sys_user.id").
		Where("sys_user.username = @username or b.value = @username", sql.Named("username", username)).
		First(&user).Error
	//如果不存在则自动注册
	if errors.Is(err, gorm.ErrRecordNotFound) {
		my.db.Save(&user)
		bind := sys.Bind{
			Type:  gt,
			Uid:   user.ID,
			Value: username,
		}
		my.db.Save(&bind)
	}
	c.Next()
}

func keyGenerate(key string) string {
	return fmt.Sprintf("%s:%s", CaptchaPrefix, util.MD5(key))
}
