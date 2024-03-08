package auth

import (
	"github.com/gin-gonic/gin"
	"github.com/gorilla/sessions"
	"net/http"
)

type Session struct {
	store *sessions.CookieStore
}

func NewSession() *Session {
	s := sessions.NewCookieStore([]byte("123456"))
	s.Options.Path = "/"
	s.Options.MaxAge = 0 // 关闭浏览器就清除 session
	return &Session{store: s}
}

// SaveUserSession 保存当前用户的信息
func (my *Session) SaveUserSession(c *gin.Context, userID string) {
	s, err := my.store.Get(c.Request, "LoginUser")
	if err != nil {
		panic(err.Error())
	}
	s.Values["userID"] = userID
	err = s.Save(c.Request, c.Writer) // save 保存
	if err != nil {
		panic(err.Error())
	}
}

// GetUserSession 获取用户信息
func (my *Session) GetUserSession(r *http.Request) string {
	if s, err := my.store.Get(r, "LoginUser"); err == nil {
		if s.Values["userID"] != nil {
			return s.Values["userID"].(string)
		}
	}
	return ""
}

// DeleteUserSession oauth2 服务端使用： 删除当前session
func (my *Session) DeleteUserSession(c *gin.Context) {
	s, err := my.store.Get(c.Request, "LoginUser")
	if err == nil {
		s.Options.MaxAge = -1             // 清除 session
		err = s.Save(c.Request, c.Writer) // 保存操作
		if err != nil {
			panic(err.Error())
		}
	} else {
		panic(err.Error())
	}
}
