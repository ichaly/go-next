package totp

import (
	"go.uber.org/fx"
)

var Modules = fx.Options(
	fx.Provide(
		//发送验证码
		fx.Annotate(
			NewCaptcha,
			fx.ResultTags(`group:"plugin"`),
		),
		//验证码验证
		fx.Annotate(
			NewCaptchaVerify,
			fx.ResultTags(`group:"middleware"`),
		),
		fx.Annotate(
			NewEmail,
			fx.ResultTags(`group:"distributor"`),
		),
	),
)
