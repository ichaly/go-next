package auth

import (
	"context"
	"github.com/go-oauth2/oauth2/v4"
	"github.com/go-oauth2/oauth2/v4/errors"
	"github.com/go-oauth2/oauth2/v4/generates"
	"github.com/go-oauth2/oauth2/v4/manage"
	"github.com/go-oauth2/oauth2/v4/server"
	"github.com/golang-jwt/jwt"
	"github.com/ichaly/go-next/app/sys"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"net/http"
	"strconv"
)

const (
	KID = "go.next"
	KEY = "oauth2_secret"
)

func NewOauthServer(db *gorm.DB, ts oauth2.TokenStore, cs oauth2.ClientStore) *server.Server {
	manager := manage.NewDefaultManager()
	manager.MapTokenStorage(ts)
	manager.MapClientStorage(cs)
	manager.MapAccessGenerate(generates.NewJWTAccessGenerate(KID, []byte(KEY), jwt.SigningMethodHS512))

	s := server.NewDefaultServer(manager)
	s.SetAllowGetAccessRequest(true)
	s.SetClientInfoHandler(clientInfoHandler())
	s.SetUserAuthorizationHandler(userAuthorizationHandler(db))
	s.SetPasswordAuthorizationHandler(passwordAuthorizationHandler(db))

	s.SetInternalErrorHandler(func(err error) (re *errors.Response) {
		return errors.NewResponse(err, http.StatusInternalServerError)
	})

	return s
}

func userAuthorizationHandler(db *gorm.DB) func(http.ResponseWriter, *http.Request) (userID string, err error) {
	return func(w http.ResponseWriter, r *http.Request) (userID string, err error) {
		//if userID = utils.GetUserSession(r); userID == "" {
		//	w.Header().Set("Location", "/login?"+r.URL.RawQuery)
		//	w.WriteHeader(302)
		//}
		return userID, nil
	}
}

func passwordAuthorizationHandler(db *gorm.DB) func(context.Context, string, string, string) (string, error) {
	return func(ctx context.Context, clientID, username, password string) (string, error) {
		user := sys.User{}
		err := db.Model(&user).Where("username = ?", username).First(&user).Error
		if err != nil {
			return "", err
		}
		err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
		if err != nil {
			return "", err
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
