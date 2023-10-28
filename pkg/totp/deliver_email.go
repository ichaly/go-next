package totp

import (
	"bytes"
	"fmt"
	"github.com/ichaly/go-next/app/sys"
	"github.com/ichaly/go-next/pkg/base"
	"github.com/jordan-wright/email"
	"html/template"
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

func (my *Email) Send(code string, to string) error {
	if len(to) <= 0 {
		return nil
	}
	//解析模版
	body, err := template.ParseFiles("cfg/email.html")
	if err != nil {
		return err
	}
	buf := new(bytes.Buffer)
	err = body.Execute(buf, struct {
		To      string
		Code    string
		Time    string
		Timeout int
	}{
		To:      to,
		Code:    code,
		Time:    time.Now().Format("2006-01-02 15:04:05"),
		Timeout: 5,
	})
	if err != nil {
		return err
	}
	e := &email.Email{
		To:      []string{to},
		Subject: "登录验证码",
		From:    my.config.Email.From,
		HTML:    buf.Bytes(),
		Headers: textproto.MIMEHeader{},
	}
	return my.pool.Send(e, 10*time.Second)
}

func NewEmail(c *base.Config) Deliver {
	auth := smtp.PlainAuth("", c.Email.Username, c.Email.Password, c.Email.Host)
	p, err := email.NewPool(fmt.Sprintf("%s:%d", c.Email.Host, c.Email.Port), 5, auth)
	if err != nil {
		return nil
	}
	return &Email{config: c, pool: p}
}
