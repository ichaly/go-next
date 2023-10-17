package auth

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/go-oauth2/oauth2/v4/server"
	"github.com/ichaly/go-next/pkg/base"
)

type verify struct {
	Oauth *server.Server
}

func NewOauthVerify(s *server.Server) base.Plugin {
	return &verify{Oauth: s}
}

func (my *verify) Base() string {
	return ""
}

func (my *verify) Init(r gin.IRouter) {
	//使用中间件鉴权
	r.Use(my.verifyHandler)
}

func (my *verify) verifyHandler(c *gin.Context) {
	if token, err := my.Oauth.ValidationBearerToken(c.Request); err == nil {
		c.Request.WithContext(context.WithValue(c.Request.Context(), base.UserContextKey, token.GetUserID()))
	}
	c.Next()
}
