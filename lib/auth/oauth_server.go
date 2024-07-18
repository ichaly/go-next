package auth

import (
	"context"
	"fmt"
	"github.com/go-oauth2/oauth2/v4"
	"github.com/go-oauth2/oauth2/v4/errors"
	"github.com/go-oauth2/oauth2/v4/generates"
	"github.com/go-oauth2/oauth2/v4/manage"
	"github.com/go-oauth2/oauth2/v4/server"
	"github.com/golang-jwt/jwt"
	"github.com/ichaly/go-next/lib/auth/internal"
	"github.com/ichaly/go-next/pkg/sys"
	"github.com/spf13/viper"
	"golang.org/x/crypto/bcrypt"
	"net/http"
	"strconv"
)

func NewOauthServer(v *viper.Viper, se *Session, ts oauth2.TokenStore, cs oauth2.ClientStore, us *sys.UserService) (*server.Server, error) {
	c := &internal.OauthConfig{}
	if err := v.Sub("oauth").Unmarshal(c); err != nil {
		return nil, err
	}

	manager := manage.NewDefaultManager()
	manager.MapTokenStorage(ts)
	manager.MapClientStorage(cs)
	manager.MapAccessGenerate(generates.NewJWTAccessGenerate("", []byte(c.Jwt.Secret), jwt.SigningMethodHS512))

	s := server.NewDefaultServer(manager)
	s.SetAllowGetAccessRequest(true)
	s.SetClientInfoHandler(clientInfoHandler())
	s.SetUserAuthorizationHandler(userAuthorizationHandler(c, se))
	s.SetPasswordAuthorizationHandler(passwordAuthorizationHandler(us))

	s.SetInternalErrorHandler(func(err error) (re *errors.Response) {
		return errors.NewResponse(err, http.StatusInternalServerError)
	})

	return s, nil
}

func userAuthorizationHandler(c *internal.OauthConfig, s *Session) func(http.ResponseWriter, *http.Request) (userID string, err error) {
	return func(w http.ResponseWriter, r *http.Request) (uid string, err error) {
		if uid = s.GetUserSession(r); uid == "" {
			uri := "/oauth/login"
			if len(c.LoginUri) > 0 {
				uri = c.LoginUri
			}
			w.Header().Set("Location", fmt.Sprintf("%s?%s", uri, r.URL.RawQuery))
			w.WriteHeader(http.StatusFound)
		}
		return
	}
}

func passwordAuthorizationHandler(us *sys.UserService) func(context.Context, string, string, string) (string, error) {
	return func(ctx context.Context, clientID, username, password string) (string, error) {
		usr, err := us.FindByUsername(username)
		if err != nil {
			return "", err
		}
		// 空密码不校验
		if password != "" {
			err = bcrypt.CompareHashAndPassword([]byte(usr.Password), []byte(password))
			if err != nil {
				return "", err
			}
		}
		return strconv.FormatUint(uint64(usr.Id), 10), nil
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
