package auth

import (
	"github.com/gin-gonic/gin"
	"github.com/graphql-go/graphql/gqlerrors"
	"github.com/ichaly/go-next/pkg/base"
	"net/http"

	"github.com/go-oauth2/oauth2/v4/server"
)

type Oauth struct {
	Session *Session
	Oauth   *server.Server
}

func NewOauth(o *server.Server, s *Session) base.Plugin {
	return &Oauth{Oauth: o, Session: s}
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

func (my *Oauth) loginHandler(c *gin.Context) {
	data := map[string]string{"error": ""}
	if c.Request.Method == http.MethodPost {
		// 对提交的信息进行处理
		username, password := c.PostForm("username"), c.PostForm("password")
		uid, err := my.Oauth.PasswordAuthorizationHandler(c, "", username, password)
		if err != nil {
			data["error"] = err.Error()
		} else {
			my.Session.SaveUserSession(c, uid)
			c.Redirect(http.StatusFound, "/oauth/authorize?"+c.Request.URL.RawQuery)
			return
		}
	}
	c.HTML(200, "login.html", data)
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
	c.Redirect(301, c.Query("redirect_uri"))
}
