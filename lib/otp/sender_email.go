package otp

import (
	"bytes"
	"fmt"
	"github.com/ichaly/go-next/lib/otp/internal"
	"github.com/ichaly/go-next/pkg/sys"
	"github.com/jordan-wright/email"
	"github.com/spf13/viper"
	"html/template"
	"net/smtp"
	"time"
)

const EMAIL sys.BindKind = "email"

type Email struct {
	pool   *email.Pool
	config *internal.EmailConfig
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
		"Timeout": my.config.Expired.Minutes(),
		"Time":    time.Now().Format("2006-01-02 15:04:05"),
	})
	if err != nil {
		return err
	}
	e := &email.Email{
		To:      []string{to},
		Subject: "登录验证码",
		From:    my.config.From,
		HTML:    buf.Bytes(),
	}
	return my.pool.Send(e, 10*time.Second)
}

func NewEmail(v *viper.Viper) (Sender, error) {
	c := &internal.EmailConfig{}
	if err := v.Unmarshal(c); err != nil {
		return nil, err
	}

	auth := smtp.PlainAuth("", c.Username, c.Password, c.Host)
	p, err := email.NewPool(fmt.Sprintf("%s:%d", c.Host, c.Port), 5, auth)
	if err != nil {
		return nil, err
	}

	return &Email{config: c, pool: p}, nil
}
