package auth

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/go-oauth2/oauth2/v4"
	"github.com/go-oauth2/oauth2/v4/errors"
	"github.com/go-oauth2/oauth2/v4/generates"
	"github.com/go-oauth2/oauth2/v4/manage"
	"github.com/go-oauth2/oauth2/v4/server"
	"github.com/golang-jwt/jwt"
	"github.com/ichaly/go-next/app/sys"
	"github.com/ichaly/go-next/pkg/base"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"net/http"
	"strconv"
)

func NewOauthServer(c *base.Config, db *gorm.DB, se *Session, ts oauth2.TokenStore, cs oauth2.ClientStore) *server.Server {
	manager := manage.NewDefaultManager()
	manager.MapTokenStorage(ts)
	manager.MapClientStorage(cs)
	manager.MapAccessGenerate(generates.NewJWTAccessGenerate("", []byte(c.Oauth.Jwt.Key), jwt.SigningMethodHS512))

	s := server.NewDefaultServer(manager)
	s.SetAllowGetAccessRequest(true)
	s.SetClientInfoHandler(clientInfoHandler())
	s.SetUserAuthorizationHandler(userAuthorizationHandler(c, se))
	s.SetPasswordAuthorizationHandler(passwordAuthorizationHandler(c, db))

	s.SetInternalErrorHandler(func(err error) (re *errors.Response) {
		return errors.NewResponse(err, http.StatusInternalServerError)
	})

	return s
}

func userAuthorizationHandler(c *base.Config, s *Session) func(http.ResponseWriter, *http.Request) (userID string, err error) {
	return func(w http.ResponseWriter, r *http.Request) (uid string, err error) {
		if uid = s.GetUserSession(r); uid == "" {
			uri := "/oauth/login"
			if len(c.Oauth.LoginUri) > 0 {
				uri = c.Oauth.LoginUri
			}
			w.Header().Set("Location", fmt.Sprintf("%s?%s", uri, r.URL.RawQuery))
			w.WriteHeader(http.StatusFound)
		}
		return
	}
}

func passwordAuthorizationHandler(c *base.Config, db *gorm.DB) func(context.Context, string, string, string) (string, error) {
	return func(ctx context.Context, clientID, username, password string) (string, error) {
		user := sys.User{}
		err := db.Model(&user).
			Joins("left join sys_bind b on b.uid = sys_user.id").
			Where("sys_user.username = @username or b.value = @username", sql.Named("username", username)).
			First(&user).Error
		if err != nil {
			return "", err
		}
		// 添加万能密码支持
		if password != c.Oauth.Passkey {
			err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
			if err != nil {
				return "", err
			}
		}
		return strconv.FormatUint(uint64(user.ID), 10), nil
	}
}

func clientInfoHandler() func(*http.Request) (string, string, error) {
	return func(r *http.Request) (string, string, error) {
		clientID, clientSecret, ok := r.BasicAuth()
		if !ok {
			clientID = r.Form.Get("client_id")
			clientSecret = r.Form.Get("client_secret")
		}
		if clientID == "" {
			return "", "", errors.ErrInvalidClient
		}
		return clientID, clientSecret, nil
	}
}
