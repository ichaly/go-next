package totp

import (
	openapi "github.com/alibabacloud-go/darabonba-openapi/v2/client"
	dysmsapi "github.com/alibabacloud-go/dysmsapi-20170525/v3/client"
	utilapi "github.com/alibabacloud-go/tea-utils/v2/service"
	"github.com/alibabacloud-go/tea/tea"
	"github.com/ichaly/go-next/app/sys"
	"github.com/ichaly/go-next/pkg/base"
	"github.com/ichaly/go-next/pkg/util"
	"github.com/ichaly/go-next/pkg/zlog"
)

const MOBILE sys.OauthKind = "mobile"

type Mobile struct {
	config *base.Config
	client *dysmsapi.Client
}

func (my *Mobile) Support(kind string) bool {
	return string(MOBILE) == kind
}

func (my *Mobile) Execute(to, code string) error {
	params, _ := util.MarshalJson(map[string]string{
		"code": code, "purpose": "随心APP",
	})
	request := &dysmsapi.SendSmsRequest{
		PhoneNumbers:  tea.String(to),
		TemplateParam: tea.String(params),
		SignName:      tea.String(my.config.Mobile.SignName),
		TemplateCode:  tea.String(my.config.Mobile.TemplateCode),
	}
	res, err := my.client.SendSmsWithOptions(request, &utilapi.RuntimeOptions{})
	if err != nil {
		return err
	}
	zlog.Info().Any("result", res).Msg("短信发送结果")
	return nil
}

func NewMobile(config *base.Config) (Sender, error) {
	client, err := dysmsapi.NewClient(&openapi.Config{
		Endpoint:        tea.String("dysmsapi.aliyuncs.com"),
		AccessKeyId:     tea.String(config.Mobile.AccessKeyId),
		AccessKeySecret: tea.String(config.Mobile.AccessKeySecret),
	})
	if err != nil {
		return nil, err
	}
	return &Mobile{config: config, client: client}, nil
}
