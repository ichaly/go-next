package auth

import (
	"github.com/gin-gonic/gin"
	"github.com/go-oauth2/oauth2/v4/server"
	"github.com/graphql-go/graphql/gqlerrors"
	"github.com/ichaly/go-next/pkg/base"
	"net/http"
)

type Oauth struct {
	session *Session
	config  *base.Config
	oauth   *server.Server
}

func NewOauth(c *base.Config, o *server.Server, s *Session) base.Plugin {
	return &Oauth{config: c, oauth: o, session: s}
}

func (my *Oauth) Base() string {
	return "/oauth"
}

func (my *Oauth) Init(r gin.IRouter) {
	//登录退出
	r.POST("/login", my.loginHandler)
	r.Match([]string{http.MethodGet, http.MethodPost}, "/logout", my.logoutHandler)

	//授权路由
	r.Match([]string{http.MethodGet, http.MethodPost}, "/token", my.tokenHandler)
	r.Match([]string{http.MethodGet, http.MethodPost}, "/authorize", my.authorizeHandler)

	//登录回调
	r.Any("/callback/:name", my.callbackHandler)
}

func (my *Oauth) tokenHandler(c *gin.Context) {
	if err := my.oauth.HandleTokenRequest(c.Writer, c.Request); err != nil {
		c.JSON(http.StatusInternalServerError, gqlerrors.FormatErrors(err))
	}
}

func (my *Oauth) authorizeHandler(c *gin.Context) {
	if err := my.oauth.HandleAuthorizeRequest(c.Writer, c.Request); err != nil {
		c.JSON(http.StatusInternalServerError, gqlerrors.FormatErrors(err))
	}
}

func (my *Oauth) loginHandler(ctx *gin.Context) {
	c, u, p := ctx.Query("client_id"), ctx.PostForm("username"), ctx.PostForm("password")
	if uid, err := my.oauth.PasswordAuthorizationHandler(ctx, c, u, p); err == nil {
		my.session.SaveUserSession(ctx, uid)
		ctx.JSON(http.StatusOK, gin.H{"redirect": "/oauth/authorize?" + ctx.Request.URL.RawQuery})
	} else {
		ctx.JSON(http.StatusInternalServerError, gin.H{"errors": gqlerrors.FormatErrors(err)})
	}
}

func (my *Oauth) logoutHandler(c *gin.Context) {
	defer func() {
		if err := recover(); err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"errors": gqlerrors.FormatErrors(err.(error))})
		}
	}()
	val := c.DefaultQuery("access_token", c.PostForm("access_token"))
	token, err := my.oauth.Manager.LoadAccessToken(c, val)
	if err != nil {
		panic(err)
	}
	err = my.oauth.Manager.RemoveAccessToken(c, token.GetAccess())
	if err != nil {
		panic(err)
	}
	my.session.DeleteUserSession(c)
	uri := c.DefaultQuery("redirect_uri", c.PostForm("redirect_uri"))
	c.JSON(http.StatusOK, gin.H{"redirect": uri})
}

func (my *Oauth) callbackHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"code": c.Query("code"), "name": c.Param("name")})
}
