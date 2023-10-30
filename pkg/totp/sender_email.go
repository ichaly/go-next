package totp

import (
	"bytes"
	"fmt"
	"github.com/ichaly/go-next/app/sys"
	"github.com/ichaly/go-next/pkg/base"
	"github.com/jordan-wright/email"
	"html/template"
	"net/smtp"
	"time"
)

const EMAIL sys.OauthKind = "email"

type Email struct {
	pool   *email.Pool
	config *base.Config
}

func (my *Email) Support(kind string) bool {
	return string(EMAIL) == kind
}

func (my *Email) Execute(to, code string) error {
	if len(to) <= 0 {
		return nil
	}
	//解析模版
	tpl, err := template.ParseFiles("cfg/email.html")
	if err != nil {
		return err
	}
	buf := new(bytes.Buffer)
	err = tpl.Execute(buf, map[string]any{
		"To":      to,
		"Code":    code,
		"Timeout": my.config.Captcha.Expired.Minutes(),
		"Time":    time.Now().Format("2006-01-02 15:04:05"),
	})
	if err != nil {
		return err
	}
	e := &email.Email{
		To:      []string{to},
		Subject: "登录验证码",
		From:    my.config.Email.From,
		HTML:    buf.Bytes(),
	}
	return my.pool.Send(e, 10*time.Second)
}

func NewEmail(c *base.Config) Sender {
	auth := smtp.PlainAuth("", c.Email.Username, c.Email.Password, c.Email.Host)
	p, err := email.NewPool(fmt.Sprintf("%s:%d", c.Email.Host, c.Email.Port), 5, auth)
	if err != nil {
		return nil
	}
	return &Email{config: c, pool: p}
}
