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
	if err := v.Unmarshal(cfg); err != nil {
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
	if err := v.Unmarshal(cfg); err != nil {
		t.Fatal(err)
	}
	str, err := util.MarshalJson(cfg)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(str)
}

func TestPhoneConfig(t *testing.T) {
	cfg := &internal.PhoneConfig{}
	if err := v.Unmarshal(cfg); err != nil {
		t.Fatal(err)
	}
	str, err := util.MarshalJson(cfg)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(str)
}
