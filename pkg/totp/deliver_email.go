package totp

import (
	"fmt"
	"github.com/ichaly/go-next/app/sys"
	"github.com/ichaly/go-next/pkg/base"
	"github.com/jordan-wright/email"
	"net/smtp"
	"net/textproto"
	"time"
)

type Email struct {
	pool   *email.Pool
	config *base.Config
}

func (my *Email) Support(kind string) bool {
	return string(sys.Email) == kind
}

func (my *Email) Send(code string, to ...string) error {
	if len(to) <= 0 {
		return nil
	}
	e := &email.Email{
		To:      to,
		Subject: "登录验证码",
		From:    my.config.Email.From,
		Text:    []byte("Text Body is, of course, supported!" + code),
		HTML:    []byte("<h1>Fancy HTML is supported, too!</h1>"),
		Headers: textproto.MIMEHeader{},
	}
	return my.pool.Send(e, 1*time.Minute)
}

func NewEmail(c *base.Config) Deliver {
	auth := smtp.PlainAuth("", c.Email.Username, c.Email.Password, c.Email.Host)
	p, err := email.NewPool(fmt.Sprintf("%s:%d", c.Email.Host, c.Email.Port), 5, auth)
	if err != nil {
		return nil
	}
	return &Email{config: c, pool: p}
}
