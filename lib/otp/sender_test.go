package otp

import (
	"github.com/ichaly/go-next/lib/base"
	"github.com/ichaly/go-next/lib/otp/internal"
	"github.com/ichaly/go-next/lib/util"
	"github.com/spf13/viper"
	"testing"
)

var v *viper.Viper

func init() {
	val, err := base.NewViper("../../cfg/dev.yml")
	if err != nil {
		panic(err)
	}
	v = val
}

func TestSenderConfig(t *testing.T) {
	cfg := &internal.CaptchaConfig{}
	err := v.Sub("captcha").Unmarshal(cfg)
	if err != nil {
		t.Fatal(err)
	}
	str, err := util.MarshalJson(cfg)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(str)
}

func TestEmailConfig(t *testing.T) {
	cfg := &internal.EmailConfig{}
	err := v.Unmarshal(cfg)
	if err != nil {
		t.Fatal(err)
	}
	str, err := util.MarshalJson(cfg)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(str)
}
