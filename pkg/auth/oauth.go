package auth

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/graphql-go/graphql/gqlerrors"
	"github.com/ichaly/go-next/pkg/base"
	"net/http"
	"strings"

	"github.com/go-oauth2/oauth2/v4/server"
)

type Oauth struct {
	Session *Session
	Config  *base.Config
	Oauth   *server.Server
}

func NewOauth(c *base.Config, o *server.Server, s *Session) base.Plugin {
	return &Oauth{Config: c, Oauth: o, Session: s}
}

func (my *Oauth) Base() string {
	return "/oauth"
}

func (my *Oauth) Init(r gin.IRouter) {
	//登录退出
	r.GET("/logout", my.logoutHandler)
	r.Match([]string{http.MethodGet, http.MethodPost}, "/login", my.loginHandler)

	//授权路由
	r.Match([]string{http.MethodGet, http.MethodPost}, "/token", my.tokenHandler)
	r.Match([]string{http.MethodGet, http.MethodPost}, "/authorize", my.authorizeHandler)
}

func (my *Oauth) tokenHandler(c *gin.Context) {
	if err := my.Oauth.HandleTokenRequest(c.Writer, c.Request); err != nil {
		c.JSON(http.StatusInternalServerError, gqlerrors.FormatErrors(err))
	}
}

func (my *Oauth) authorizeHandler(c *gin.Context) {
	if err := my.Oauth.HandleAuthorizeRequest(c.Writer, c.Request); err != nil {
		c.JSON(http.StatusInternalServerError, gqlerrors.FormatErrors(err))
	}
}

func (my *Oauth) loginHandler(ctx *gin.Context) {
	var err error
	var uid string
	// 如果是post请求则走登录校验
	if ctx.Request.Method == http.MethodPost {
		c, u, p := ctx.PostForm("client_id"), ctx.PostForm("username"), ctx.PostForm("password")
		if uid, err = my.Oauth.PasswordAuthorizationHandler(ctx, c, u, p); err == nil {
			my.Session.SaveUserSession(ctx, uid)
		}
	}
	//如果是ajax请求则返回json
	if my.isAjaxRequest(ctx) {
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"errors": gqlerrors.FormatErrors(err)})
		} else {
			ctx.JSON(http.StatusOK, gin.H{"redirect": "/oauth/authorize?" + ctx.Request.URL.RawQuery})
		}
		return
	}
	if len(uid) > 0 {
		ctx.Redirect(http.StatusFound, fmt.Sprintf("/oauth/authorize?%s", ctx.Request.URL.RawQuery))
		return
	}
	res := map[string]string{"error": ""}
	if err != nil {
		res["error"] = err.Error()
	}
	ctx.HTML(200, "login.html", res)
}

func (my *Oauth) logoutHandler(c *gin.Context) {
	defer func() {
		if err := recover(); err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"errors": gqlerrors.FormatErrors(err.(error))})
		}
	}()
	token, err := my.Oauth.Manager.LoadAccessToken(c, c.Query("access_token"))
	if err != nil {
		panic(err)
	}
	my.Session.DeleteUserSession(c)
	err = my.Oauth.Manager.RemoveAccessToken(c, token.GetAccess())
	if err != nil {
		panic(err)
	}
	redirect := c.Query("redirect_uri")
	if my.isAjaxRequest(c) {
		c.JSON(http.StatusOK, gin.H{"redirect": redirect})
	} else {
		c.Redirect(http.StatusMovedPermanently, redirect)
	}
}

// 判断请求是否是ajax请求
func (my *Oauth) isAjaxRequest(c *gin.Context) bool {
	accept := c.GetHeader("accept")
	if accept != "" && strings.Contains(accept, "application/json") {
		return true
	}
	return c.GetHeader("X-Requested-With") == "XMLHttpRequest"
}
