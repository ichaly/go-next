package otp

import (
	openapi "github.com/alibabacloud-go/darabonba-openapi/v2/client"
	dysmsapi "github.com/alibabacloud-go/dysmsapi-20170525/v3/client"
	utilapi "github.com/alibabacloud-go/tea-utils/v2/service"
	"github.com/alibabacloud-go/tea/tea"
	"github.com/ichaly/go-next/lib/otp/internal"
	"github.com/ichaly/go-next/lib/util"
	"github.com/ichaly/go-next/lib/zlog"
	"github.com/ichaly/go-next/pkg/sys"
	"github.com/spf13/viper"
)

const PHONE sys.BindKind = "phone"

type Phone struct {
	client *dysmsapi.Client
	config *internal.PhoneConfig
}

func (my *Phone) Support(kind string) bool {
	return string(PHONE) == kind
}

func (my *Phone) Execute(to, code string) error {
	params, _ := util.MarshalJson(map[string]string{
		"code": code, "purpose": "随心APP",
	})
	request := &dysmsapi.SendSmsRequest{
		PhoneNumbers:  tea.String(to),
		TemplateParam: tea.String(params),
		SignName:      tea.String(my.config.SignName),
		TemplateCode:  tea.String(my.config.TemplateCode),
	}
	res, err := my.client.SendSmsWithOptions(request, &utilapi.RuntimeOptions{})
	if err != nil {
		return err
	}
	zlog.Info().Any("result", res).Msg("短信发送结果")
	return nil
}

func NewPhone(v *viper.Viper) (Sender, error) {
	c := &internal.PhoneConfig{}
	if err := v.Unmarshal(c); err != nil {
		return nil, err
	}

	client, err := dysmsapi.NewClient(&openapi.Config{
		Endpoint:        tea.String("dysmsapi.aliyuncs.com"),
		AccessKeyId:     tea.String(c.AccessKeyId),
		AccessKeySecret: tea.String(c.AccessKeySecret),
	})
	if err != nil {
		return nil, err
	}

	return &Phone{config: c, client: client}, nil
}
